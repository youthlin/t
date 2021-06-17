package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/youthlin/t"
	"github.com/youthlin/t/cmd/xtemplate/internal"
)

// inject when build
var Version string

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Fprintf(os.Stderr, t.T("unexpected error: %+v\n"), e)
		}
	}()
	initTranslation()
	ctx := buildCtx()
	err := internal.Run(ctx)
	if err != nil {
		if ctx.Debug {
			fmt.Fprintf(os.Stderr, t.T("run error: %+v"), err)
		} else {
			fmt.Fprintf(os.Stderr, t.T("run error: %v"), err)
		}
		os.Exit(2)
	}
}

// initTranslation init i18n
func initTranslation() {
	dir, ok := os.LookupEnv("LOCALEDIR")
	if !ok {
		var err error
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			dir = "lang"
		} else {
			dir = dir + "/lang"
		}
	}
	// fmt.Printf("dir=%v\n", dir)
	t.BindDefaultDomain(dir)
	t.SetLocale("")
}

// buildCtx parse os.Args
func buildCtx() *internal.Context {
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
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), t.T("Usage of %s:\nxtemplate -i input-pattern -k keywords [-f functions] [-o output]\n"), os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	if *version {
		fmt.Fprintf(os.Stdout, `xtemplate
https://github.com/youthlin/t/tree/main/cmd/xtemplate
by Youth．霖(https://youthlin.com)
version: %v
`,
			Version)
		os.Exit(0)
	}
	if *help || len(os.Args) < 5 {
		flag.Usage()
		os.Exit(0)
	}
	return &internal.Context{
		Input:      *input,
		Left:       *left,
		Right:      *right,
		Keyword:    *keywords,
		Function:   *fun,
		OutputFile: *output,
		Debug:      *debug,
	}
}
