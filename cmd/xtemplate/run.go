package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"text/template/parse"

	"github.com/youthlin/t"
)

func run(param *Param) {
	filenames, err := filepath.Glob(param.input)
	if err != nil {
		exit(t.T("invalid input pattern: %+v"), err)
	}
	for _, filename := range filenames {
		resolveOneFile(filename, param)
	}
}
func exit(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(2)
}
func debugPrint(format string, args ...interface{}) {
	if *debug {
		fmt.Printf(format+"\n", args...)
	}
}
func resolveOneFile(filename string, param *Param) {
	debugPrint("filename: %v", filename)
	var funcs = make(template.FuncMap)
	for _, k := range param.keywords {
		funcs[k.Name] = func() string {
			return ""
		}
	}
	tmpl, err := template.New(filename).
		Delims(param.left, param.right).
		Funcs(funcs).
		ParseFiles(filename)
	if err != nil {
		debugPrint("error on parse file %v: %+v", filename, err)
		return
	}
	for _, tmpl := range tmpl.Templates() {
		resolveTmpl(filename, param, tmpl)
	}
}

func resolveTmpl(filename string, param *Param, tmpl *template.Template) {
	debugPrint("process template: [filename=%v] [template name=%v]", filename, tmpl.Name())
	if tmpl.Tree == nil || tmpl.Tree.Root == nil {
		debugPrint("  > filename=%v, template=%v, tree or Root is nil", filename, tmpl.Name())
		return
	}
	root := tmpl.Tree.Root
	for _, node := range root.Nodes {
		if node.Type() == parse.NodeAction {
			actionNode := node.(*parse.ActionNode)
			pipe := actionNode.Pipe
			if pipe == nil {
				debugPrint("  > line %v: Pipe is nil", actionNode.Line)
				continue
			}
			debugPrint("  >  Pipe: Line=%v Decl=%#v", actionNode.Line, pipe.Decl)
			if pipe.Decl != nil {
				for _, decl := range pipe.Decl {
					debugPrint("  >  > Decl: %#v", decl)
				}
			}
			resolveCmds(filename, param, pipe.Cmds)
		}
	}
}
func resolveCmds(filename string, param *Param, cmds []*parse.CommandNode) {
	for _, cmd := range cmds {
		if cmd == nil {
			continue
		}
		debugPrint("  >  > Cmd: Pos %v", cmd.Pos)
		for _, arg := range cmd.Args {
			debugPrint("  >  >  >  Cmd.Arg: %#v", arg)
		}
	}
}
