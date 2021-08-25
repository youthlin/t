package internal

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"text/template/parse"

	"github.com/cockroachdb/errors"
	"github.com/youthlin/t"
	"github.com/youthlin/t/translator"
)

var noopFun = func() string { return "" }

// run 运行解析任务
func Run(param *Param) error {
	param.debugPrint("run param=%+v", param)
	ctx, err := newCtx(param)
	if err != nil {
		return err
	}
	filenames, err := filepath.Glob(param.Input)
	param.debugPrint("Glob files=%v err=%+v", filenames, err)
	if err != nil {
		return errors.Wrapf(err, t.T("invalid input pattern"))
	}

	for _, filename := range filenames {
		if err := resolveOneFile(filename, ctx); err != nil {
			if param.Debug {
				printErr(t.T("failed to process file %v. error message: %+v"), filename, err)
			} else {
				printErr(t.T("failed to process file %v. error message: %v"), filename, err)
			}
		}
	}

	ctx.debugPrint("extract done, %d entries", len(ctx.entries))

	if err := ctx.Write(); err != nil {
		return err
	}
	return ctx.Output.Close()
}

// printErr print message to stderr
func printErr(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

// resolveOneFile 处理每个文件
func resolveOneFile(filename string, ctx *Context) error {
	ctx.debugPrint("resolve one file: filename=%v", filename)
	tmpl, err := template.New("").
		Delims(ctx.Left, ctx.Right).
		Funcs(ctx.Functions).
		ParseFiles(filename)
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
	root := tmpl.Tree.Root
	for _, node := range root.Nodes {
		// param.debugPrint("  > node=%#v", node)
		// comment 会被忽略，这里拿不到注释信息
		if node.Type() == parse.NodeAction {
			// 只需要关注 action 节点
			actionNode := node.(*parse.ActionNode)
			resolvePipe(filename, actionNode.Line, ctx, actionNode.Pipe)
		}
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
	for _, cmd := range cmds {
		if cmd == nil {
			continue
		}
		ctx.debugPrint("  >  >  Cmd: Line=%v Pos=%v", line, cmd.Pos)
		argC := len(cmd.Args)
		for i := 0; i < argC; i++ {
			arg := cmd.Args[i]
			ctx.debugPrint("  >  >  >  Cmd.Arg: %#v", arg)
			switch arg := arg.(type) {
			case *parse.PipeNode:
				resolvePipe(filename, line, ctx, arg) // 递归
			case *parse.IdentifierNode:
				filter(ctx, fmt.Sprintf("%v:%d", filename, line), arg.Ident, i, cmd.Args)
			case *parse.FieldNode:
				lastID := arg.Ident[len(arg.Ident)-1]
				filter(ctx, fmt.Sprintf("%v:%d", filename, line), lastID, i, cmd.Args)
			}
		}
	}
}

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
				str, ok := arg.(*parse.StringNode)
				if !ok {
					ctx.debugPrint("  >  >  >  ID=%v args[%d] is not string node", name, i)
					argOK = false
					break
				}
				m[i-nameIndex] = str.Text
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
					printErr(t.T("Waringing: %+v"), err)
				} else {
					printErr(t.T("Waringing: %v"), err)
				}
			}
		}
	}
}

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
