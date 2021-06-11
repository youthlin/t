package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/youthlin/t"
)

func run(param *Param) {
	fmt.Printf("param=%#v %#v\n", param, os.Args)
	tmpl, err := template.New("").Delims(param.left, param.right).ParseGlob(param.input)
	if err != nil {
		fmt.Fprintf(os.Stderr, t.T("failed to parse template|err=%+v"), err)
		return
	}
	fmt.Printf("%#v\n", tmpl.Tree)
}
