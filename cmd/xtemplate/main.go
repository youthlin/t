package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

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
	output   = flag.String("o", "message.pot", t.T("output file"))
	debug    = flag.Bool("d", false, t.T("debug mode"))
	version  = flag.Bool("v", false, t.T("show version"))
	help     = flag.Bool("h", false, t.T("show this help message"))
)

func main() {
	flag.Parse()
	if *help || len(os.Args) < 5 {
		flag.Usage()
		return
	}
	if *version {
		fmt.Fprintf(os.Stdout, t.T("version: %v\n"), "v0.0.0")
		return
	}
	defer func() {
		if e := recover(); e != nil {
			fmt.Fprintf(os.Stderr, t.T("unexpected error: %v"), e)
			flag.Usage()
		}
	}()
	keywords, err := internal.ParseKeywords(*keywords)
	if err != nil {
		flag.Usage()
		os.Exit(2)
	}
	out, err := internal.Writer(*output)
	if err != nil {
		flag.Usage()
		os.Exit(2)
	}
	internal.Run(&internal.Context{
		Input:    *input,
		Left:     *left,
		Right:    *right,
		Keywords: keywords,
		Fun:      strings.Split(*fun, ","),
		Output:   out,
		Debug:    *debug,
	})
}
