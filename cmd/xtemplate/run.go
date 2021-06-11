package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/youthlin/t"
)

func run(param *Param) {
	// todo use single file so that we can add line number to pot file.
	fmt.Printf("param=%#v %#v\n", param, os.Args)
	tmpl, err := template.New("").Delims(param.left, param.right).ParseGlob(param.input)
	if err != nil {
		fmt.Fprintf(os.Stderr, t.T("failed to parse template|err=%+v"), err)
		return
	}
	for _, tmpl := range tmpl.Templates() {
		if tmpl.Tree != nil {
			fmt.Printf("%s:\t%#v\n", tmpl.Name(), tmpl.Tree.Root.Pos)
		}
	}
}
