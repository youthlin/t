package internal

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"text/template/parse"

	"github.com/cockroachdb/errors"
	"github.com/youthlin/t"
	"github.com/youthlin/t/translator"
)

var noopFun = func() string { return "" }

// templateBuiltins 是 Go template 内置关键字，不应被当作自定义函数注册。
var templateBuiltins = map[string]bool{
	"if": true, "else": true, "end": true, "range": true, "with": true,
	"template": true, "define": true, "block": true, "break": true, "continue": true,
}

// identPattern 匹配模板中所有合法 Go 标识符（字母或下划线开头）。
// 扫描所有标识符并注册为 noop 函数，比精确匹配函数调用位置更简单可靠。
var identPattern = regexp.MustCompile(`[a-zA-Z_]\w*`)

// scanFuncNames 从模板源码中扫描所有标识符，排除内置关键字，作为候选函数名。
func scanFuncNames(src string) map[string]bool {
	names := make(map[string]bool)
	for _, name := range identPattern.FindAllString(src, -1) {
		if !templateBuiltins[name] {
			names[name] = true
		}
	}
	return names
}

// Run 运行解析任务
func Run(param *Param) (err error) {
	param.debugPrint("run param=%+v", param)
	ctx, err := newCtx(param)
	if err != nil {
		return err
	}
	defer func() {
		closeErr := ctx.Close()
		if err == nil {
			err = closeErr
		}
	}()
	filenames, err := filepath.Glob(param.Input)
	param.debugPrint("Glob files=%v err=%+v", filenames, err)
	if err != nil {
		return errors.Wrapf(err, t.T("invalid input pattern"))
	}

	var firstErr error
	for _, filename := range filenames {
		if err := resolveOneFile(filename, ctx); err != nil {
			if param.Debug {
				printErr(t.T("failed to process file %v. error message: %+v"), filename, err)
			} else {
				printErr(t.T("failed to process file %v. error message: %v"), filename, err)
			}
			if firstErr == nil {
				firstErr = err
			}
		}
	}
	if firstErr != nil {
		return firstErr
	}

	ctx.debugPrint("extract done, %d entries", len(ctx.entries))

	return ctx.Write()
}

// printErr print message to stderr
func printErr(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

// resolveOneFile 处理每个文件
func resolveOneFile(filename string, ctx *Context) error {
	ctx.debugPrint("resolve one file: filename=%v", filename)

	// 读取文件内容，自动扫描模板中使用的函数名并注册为 noop，
	// 避免用户必须通过 -f 手动指定所有模板函数。
	src, err := os.ReadFile(filename)
	if err != nil {
		return errors.Wrapf(err, t.T("failed to read file %v"), filename)
	}
	funcs := make(template.FuncMap, len(ctx.Functions)+8)
	for k, v := range ctx.Functions {
		funcs[k] = v
	}
	for name := range scanFuncNames(string(src)) {
		if _, ok := funcs[name]; !ok {
			funcs[name] = noopFun
		}
	}

	tmpl, err := template.New("").
		Delims(ctx.Left, ctx.Right).
		Funcs(funcs).
		Parse(string(src))
	if err != nil {
		return errors.Wrapf(err, t.T("failed to parse file %v"), filename)
	}
	// 一个文件可能有多个模板
	for _, tmpl := range tmpl.Templates() {
		resolveTmpl(filename, ctx, tmpl)
	}
	return nil
}

// resolveTmpl 处理每个模板
func resolveTmpl(filename string, ctx *Context, tmpl *template.Template) {
	ctx.debugPrint("process template: [filename=%v] [template name=%v]", filename, tmpl.Name())
	if tmpl.Tree == nil || tmpl.Tree.Root == nil {
		ctx.debugPrint("  > filename=%v, template=%v, tree or Root is nil", filename, tmpl.Name())
		return
	}
	resolveNode(filename, ctx, tmpl.Tree.Root)
}

// resolveNode 递归遍历模板 AST。
// 这里显式处理控制结构节点，确保 if/range/with/template 中的翻译调用都能继续向下扫描。
// 注意 parse.Node 里可能出现 typed nil，因此每个具体分支里都要再做一次 nil 保护。
func resolveNode(filename string, ctx *Context, node parse.Node) {
	if node == nil {
		return
	}
	switch node := node.(type) {
	case *parse.ListNode:
		if node == nil {
			return
		}
		for _, child := range node.Nodes {
			resolveNode(filename, ctx, child)
		}
	case *parse.ActionNode:
		if node == nil {
			return
		}
		resolvePipe(filename, node.Line, ctx, node.Pipe)
	case *parse.IfNode:
		if node == nil {
			return
		}
		resolvePipe(filename, node.Line, ctx, node.Pipe)
		resolveNode(filename, ctx, node.List)
		resolveNode(filename, ctx, node.ElseList)
	case *parse.RangeNode:
		if node == nil {
			return
		}
		resolvePipe(filename, node.Line, ctx, node.Pipe)
		resolveNode(filename, ctx, node.List)
		resolveNode(filename, ctx, node.ElseList)
	case *parse.WithNode:
		if node == nil {
			return
		}
		resolvePipe(filename, node.Line, ctx, node.Pipe)
		resolveNode(filename, ctx, node.List)
		resolveNode(filename, ctx, node.ElseList)
	case *parse.TemplateNode:
		if node == nil {
			return
		}
		resolvePipe(filename, node.Line, ctx, node.Pipe)
	}
}

// resolvePipe 处理 action 节点中的 pipe
func resolvePipe(filename string, line int, ctx *Context, pipe *parse.PipeNode) {
	if pipe == nil {
		ctx.debugPrint("  > line %v: Pipe is nil", line)
		return
	}
	ctx.debugPrint("  >  Pipe: Line=%v", line)
	resolveCmds(filename, line, ctx, pipe.Cmds)
}

// resolveCmds 处理 Cmd
func resolveCmds(filename string, line int, ctx *Context, cmds []*parse.CommandNode) {
	for cmdIndex, cmd := range cmds {
		if cmd == nil {
			continue
		}
		ctx.debugPrint("  >  >  Cmd: Line=%v Pos=%v", line, cmd.Pos)
		args := cmd.Args
		if pipedArg := pipedStringArg(cmds, cmdIndex); pipedArg != nil {
			args = append(append([]parse.Node{}, cmd.Args...), pipedArg)
		}
		argC := len(args)
		for i := 0; i < argC; i++ {
			arg := args[i]
			ctx.debugPrint("  >  >  >  Cmd.Arg: %#v", arg)
			switch arg := arg.(type) {
			case *parse.PipeNode:
				resolvePipe(filename, line, ctx, arg) // 递归
			case *parse.IdentifierNode:
				filter(ctx, fmt.Sprintf("%v:%d", filename, line), arg.Ident, i, args)
			case *parse.FieldNode:
				if len(arg.Ident) == 0 {
					continue
				}
				lastID := arg.Ident[len(arg.Ident)-1]
				filter(ctx, fmt.Sprintf("%v:%d", filename, line), lastID, i, args)
			case *parse.VariableNode:
				if len(arg.Ident) == 0 {
					continue
				}
				lastID := arg.Ident[len(arg.Ident)-1]
				filter(ctx, fmt.Sprintf("%v:%d", filename, line), lastID, i, args)
			case *parse.ChainNode:
				if len(arg.Field) == 0 {
					continue
				}
				lastID := arg.Field[len(arg.Field)-1]
				filter(ctx, fmt.Sprintf("%v:%d", filename, line), lastID, i, args)
			}
		}
	}
}

// pipedStringArg 处理 "foo" | T 这类管道写法。
// 在模板 AST 中，前一个 command 的输出并不会直接出现在后一个 command 的 Args 里，
// 因此这里把上一段单字符串字面量补回当前命令参数中，统一复用后面的提取逻辑。
func pipedStringArg(cmds []*parse.CommandNode, cmdIndex int) parse.Node {
	if cmdIndex <= 0 {
		return nil
	}
	prev := cmds[cmdIndex-1]
	if prev == nil || len(prev.Args) != 1 {
		return nil
	}
	if _, ok := stringLiteral(prev.Args[0]); !ok {
		return nil
	}
	return prev.Args[0]
}

// filter 在命令参数中识别是否命中了 gettext 风格关键字，若命中则尝试抽取 entry。
func filter(ctx *Context, line, name string, nameIndex int, args []parse.Node) {
	argLength := len(args)
	for _, kw := range ctx.Keywords {
		if kw.Name == name {
			argCount := 1
			if kw.MsgCtxt > 0 {
				argCount++
			}
			if kw.MsgID2 > 0 {
				argCount++
			}
			lastIndex := argCount + nameIndex
			if lastIndex >= argLength {
				ctx.debugPrint("  >  >  >  ID=%v too few args", name)
				continue
			}
			argOK := true
			m := make(map[int]string)
			for i := nameIndex + 1; i <= lastIndex; i++ {
				arg := args[i]
				text, ok := stringLiteral(arg)
				if !ok {
					ctx.debugPrint("  >  >  >  ID=%v args[%d] is not string node", name, i)
					argOK = false
					break
				}
				m[i-nameIndex] = text
			}
			if !argOK {
				continue
			}
			entry, ok := extract(ctx, line, name, kw, m)
			if !ok {
				continue
			}
			if err := ctx.Add(entry); err != nil {
				if ctx.Debug {
					printErr(t.T("Warning: %+v"), err)
				} else {
					printErr(t.T("Warning: %v"), err)
				}
			}
		}
	}
}

// stringLiteral 尝试把一个 AST 节点解析成字符串字面量。
// 除了直接的 StringNode，也支持 T ("wrapped") 这种被 PipeNode 再包一层的情况。
func stringLiteral(node parse.Node) (string, bool) {
	switch node := node.(type) {
	case *parse.StringNode:
		if node == nil {
			return "", false
		}
		return node.Text, true
	case *parse.PipeNode:
		if node == nil || len(node.Cmds) != 1 {
			return "", false
		}
		cmd := node.Cmds[0]
		if cmd == nil || len(cmd.Args) != 1 {
			return "", false
		}
		return stringLiteral(cmd.Args[0])
	default:
		return "", false
	}
}

// extract 根据 keyword 定义的参数位置，把当前命令转换成 translator.Entry。
func extract(ctx *Context, line, name string, kw Keyword, m map[int]string) (*translator.Entry, bool) {
	entry := new(translator.Entry)
	entry.MsgCmts = append(entry.MsgCmts, fmt.Sprintf("#: %v", line))
	if kw.MsgCtxt > 0 {
		txt, ok := m[kw.MsgCtxt]
		if !ok {
			ctx.debugPrint("  >  >  >  ID=%v missing ctxt", name)
			return nil, false
		}
		entry.MsgCtxt = txt
	}
	txt, ok := m[kw.MsgID]
	if !ok {
		ctx.debugPrint("  >  >  >  ID=%v missing msg id", name)
		return nil, false

	}
	entry.MsgID = txt

	if kw.MsgID2 > 0 {
		txt, ok := m[kw.MsgID2]
		if !ok {
			ctx.debugPrint("  >  >  >  ID=%v missing msg plural", name)
			return nil, false

		}
		entry.MsgID2 = txt
	}
	if isGoFormat(entry) {
		entry.MsgCmts = append(entry.MsgCmts, "#, go-format")
	}
	ctx.debugPrint("【ok】  >  >  ID=%v entry=%+v", name, entry)
	return entry, true
}
