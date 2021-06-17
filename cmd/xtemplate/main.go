package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/youthlin/t"
	"github.com/youthlin/t/cmd/xtemplate/internal"
)

// xtemplate -kT:1 -o file.pot -i input-pattern

var (
	input    = flag.String("i", "", t.T("input file pattern"))
	left     = flag.String("left", "{{", t.T("left delim"))
	right    = flag.String("right", "}}", t.T("right delim"))
	keywords = flag.String("k", "", t.T("keywords e.g.: gettext;T:1;N1,2;X:1c,2;XN:1c,2,3"))
	fun      = flag.String("f", "", t.T("function names of template"))
	output   = flag.String("o", "", t.T("output file, - is stdout"))
	debug    = flag.Bool("d", false, t.T("debug mode"))
	version  = flag.Bool("v", false, t.T("show version"))
	help     = flag.Bool("h", false, t.T("show help message"))
)

// inject when build
var Version string

func main() {
	flag.Parse()
	if *version {
		fmt.Fprintf(os.Stdout, t.T(`xtemplate
https://github.com/youthlin/t/tree/main/cmd/xtemplate
by Youth．霖(https://youthlin.com)
version: %v
`),
			Version)
		return
	}
	if *help || len(os.Args) < 5 {
		flag.Usage()
		return
	}
	defer func() {
		if e := recover(); e != nil {
			if err, ok := e.(error); ok {
				if *debug {
					fmt.Fprintf(os.Stderr, t.T("unexpected error: %+v\n"), err)
				} else {
					fmt.Fprintf(os.Stderr, t.T("unexpected error: %v\n"), err)
				}
			} else {
				fmt.Fprintf(os.Stderr, t.T("unexpected error: %v\n"), e)
			}
		}
	}()

	err := internal.Run(&internal.Context{
		Input:      *input,
		Left:       *left,
		Right:      *right,
		Keyword:    *keywords,
		Function:   *fun,
		OutputFile: *output,
		Debug:      *debug,
	})
	if err != nil {
		if *debug {
			fmt.Fprintf(os.Stderr, t.T("run error: %+v"), *output, err)
		} else {
			fmt.Fprintf(os.Stderr, t.T("run error: %v"), *output, err)
		}
		os.Exit(2)
	}
}
