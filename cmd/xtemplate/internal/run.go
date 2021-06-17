package internal

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"text/template/parse"

	"github.com/youthlin/t"
)

var noopFun = func() string { return "" }

// run 运行解析任务
func Run(param *Context) {
	filenames, err := filepath.Glob(param.Input)
	if err != nil {
		exit(t.T("invalid input pattern: %+v"), err)
	}
	for _, filename := range filenames {
		resolveOneFile(filename, param)
	}
}

// exit print message and exist
func exit(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(2)
}

// resolveOneFile 处理每个文件
func resolveOneFile(filename string, param *Context) {
	param.debugPrint("filename: %v", filename)
	var funcs = make(template.FuncMap)
	for _, k := range param.Fun {
		funcs[k] = noopFun
	}
	tmpl, err := template.New(filename).
		Delims(param.Left, param.Right).
		Funcs(funcs).
		ParseFiles(filename)
	if err != nil {
		param.debugPrint("error on parse file %v: %+v", filename, err)
		return
	}
	// 一个文件可能有多个模板
	for _, tmpl := range tmpl.Templates() {
		resolveTmpl(filename, param, tmpl)
	}
}

// resolveTmpl 处理每个模板
func resolveTmpl(filename string, param *Context, tmpl *template.Template) {
	param.debugPrint("process template: [filename=%v] [template name=%v]", filename, tmpl.Name())
	if tmpl.Tree == nil || tmpl.Tree.Root == nil {
		param.debugPrint("  > filename=%v, template=%v, tree or Root is nil", filename, tmpl.Name())
		return
	}
	root := tmpl.Tree.Root
	for _, node := range root.Nodes {
		// param.debugPrint("  > node=%#v", node)
		// comment 会被忽略，这里拿不到注释信息
		if node.Type() == parse.NodeAction {
			// 只需要关注 action 节点
			actionNode := node.(*parse.ActionNode)
			resolvePipe(filename, actionNode.Line, param, actionNode.Pipe)
		}
	}
}

// resolvePipe 处理 action 节点中的 pipe
func resolvePipe(filename string, line int, param *Context, pipe *parse.PipeNode) {
	if pipe == nil {
		param.debugPrint("  > line %v: Pipe is nil", line)
		return
	}
	param.debugPrint("  >  Pipe: Line=%v Decl=%#v", line, pipe.Decl)
	if pipe.Decl != nil {
		for _, decl := range pipe.Decl {
			param.debugPrint("  >  > Decl: %#v", decl)
		}
	}
	resolveCmds(filename, line, param, pipe.Cmds)
}

// resolveCmds 处理 Cmd
func resolveCmds(filename string, line int, param *Context, cmds []*parse.CommandNode) {
	for _, cmd := range cmds {
		if cmd == nil {
			continue
		}
		param.debugPrint("  >  > Cmd: Pos %v", cmd.Pos)
		argC := len(cmd.Args)
		for i := 0; i < argC; i++ {
			arg := cmd.Args[i]
			param.debugPrint("  >  >  >  Cmd.Arg: %#v", arg)
			switch arg := arg.(type) {
			case *parse.PipeNode:
				resolvePipe(filename, line, param, arg) // 递归
			case *parse.IdentifierNode:
				filter(param, fmt.Sprintf("%v:%d", filename, line), arg.Ident, i, cmd.Args)
			case *parse.FieldNode:
				lastID := arg.Ident[len(arg.Ident)-1]
				filter(param, fmt.Sprintf("%v:%d", filename, line), lastID, i, cmd.Args)
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
				ctx.debugPrint("  >  >  > ID=%v too few args", name)
				continue
			}
			argOK := true
			m := make(map[int]string)
			for i := nameIndex + 1; i <= lastIndex; i++ {
				arg := args[i]
				str, ok := arg.(*parse.StringNode)
				if !ok {
					ctx.debugPrint("  >  >  > ID=%v  args[%d] is not string node", name, i)
					argOK = false
					break
				}
				m[i-nameIndex] = str.Text
			}
			if !argOK {
				continue
			}
			msg := newMessage(line)
			if kw.MsgCtxt > 0 {
				txt, ok := m[kw.MsgCtxt]
				if !ok {
					ctx.debugPrint("  >  >  > ID=%v  missing ctxt", name)
					continue
				}
				msg.setCtxt(txt)
			}
			txt, ok := m[kw.MsgID]
			if !ok {
				ctx.debugPrint("  >  >  > ID=%v  missing msg id", name)
				continue
			}
			msg.setMsgID(txt)
			if kw.MsgID2 > 0 {
				txt, ok := m[kw.MsgID2]
				if !ok {
					ctx.debugPrint("  >  >  > ID=%v  missing msg plural", name)
					continue
				}
				msg.setMsgPlural(txt)
			}
			ctx.debugPrint("  >  >  >  【ID=%v ok】msg=%+v", name, msg)
			ctx.result.add(msg)
		}
	}
}
