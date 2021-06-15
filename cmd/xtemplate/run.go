package main

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
func run(param *Param) {
	filenames, err := filepath.Glob(param.input)
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

// debugPrint print if is debug mode
func debugPrint(format string, args ...interface{}) {
	if *debug {
		fmt.Printf(format+"\n", args...)
	}
}

// resolveOneFile 处理每个文件
func resolveOneFile(filename string, param *Param) {
	debugPrint("filename: %v", filename)
	var funcs = make(template.FuncMap)
	for _, k := range param.fun {
		funcs[k] = noopFun
	}
	tmpl, err := template.New(filename).
		Delims(param.left, param.right).
		Funcs(funcs).
		ParseFiles(filename)
	if err != nil {
		debugPrint("error on parse file %v: %+v", filename, err)
		return
	}
	// 一个文件可能有多个模板
	for _, tmpl := range tmpl.Templates() {
		resolveTmpl(filename, param, tmpl)
	}
}

// resolveTmpl 处理每个模板
func resolveTmpl(filename string, param *Param, tmpl *template.Template) {
	debugPrint("process template: [filename=%v] [template name=%v]", filename, tmpl.Name())
	if tmpl.Tree == nil || tmpl.Tree.Root == nil {
		debugPrint("  > filename=%v, template=%v, tree or Root is nil", filename, tmpl.Name())
		return
	}
	root := tmpl.Tree.Root
	for _, node := range root.Nodes {
		if node.Type() == parse.NodeAction {
			// 只需要关注 action 节点
			actionNode := node.(*parse.ActionNode)
			resolvePipe(filename, actionNode.Line, param, actionNode.Pipe)
		}
	}
}

// resolvePipe 处理 action 节点中的 pipe
func resolvePipe(filename string, line int, param *Param, pipe *parse.PipeNode) {
	if pipe == nil {
		debugPrint("  > line %v: Pipe is nil", line)
		return
	}
	debugPrint("  >  Pipe: Line=%v Decl=%#v", line, pipe.Decl)
	if pipe.Decl != nil {
		for _, decl := range pipe.Decl {
			debugPrint("  >  > Decl: %#v", decl)
		}
	}
	resolveCmds(filename, line, param, pipe.Cmds)
}

// resolveCmds 处理 Cmd
func resolveCmds(filename string, line int, param *Param, cmds []*parse.CommandNode) {
	for _, cmd := range cmds {
		if cmd == nil {
			continue
		}
		debugPrint("  >  > Cmd: Pos %v", cmd.Pos)
		for _, arg := range cmd.Args {
			debugPrint("  >  >  >  Cmd.Arg: %#v", arg)
			switch arg := arg.(type) {
			case *parse.PipeNode:
				resolvePipe(filename, line, param, arg) // 递归
			case *parse.IdentifierNode:
			case *parse.FieldNode:
			}
		}
	}
}
