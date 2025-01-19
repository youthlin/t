package main

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"github.com/youthlin/t"
	"github.com/youthlin/t/cmd/xtemplate/internal"
)

// Version the version
var Version string = "v0.0.9"

//go:embed lang
var embedLangs embed.FS

// initTranslation init i18n
func initTranslation() {
	path, ok := os.LookupEnv("LANG_PATH")
	if ok {
		t.Load(path)
	} else {
		t.LoadFS(embedLangs)
	}
	t.SetLocale("")
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Fprintf(os.Stderr, t.T("unexpected error: %+v\n"), e)
		}
	}()
	initTranslation()
	param := buildParam()
	err := internal.Run(param)
	if err != nil {
		if param.Debug {
			fmt.Fprintf(os.Stderr, t.T("run error: %+v"), err)
		} else {
			fmt.Fprintf(os.Stderr, t.T("run error: %v"), err)
		}
		os.Exit(2)
	}
}

// buildCtx parse os.Args
func buildParam() *internal.Param {
	var (
		input    = flag.String("i", "", t.T("input file pattern"))
		left     = flag.String("left", "{{", t.T("left delim"))
		right    = flag.String("right", "}}", t.T("right delim"))
		keywords = flag.String("k", "", t.T("keywords e.g.: gettext;T:1;N:1,2;X:1c,2;XN:1c,2,3"))
		fun      = flag.String("f", "", t.T("function names of template"))
		output   = flag.String("o", "", t.T("output file, - is stdout"))
		debug    = flag.Bool("d", false, t.T("debug mode"))
		version  = flag.Bool("v", false, t.T("show version"))
		help     = flag.Bool("h", false, t.T("show help message"))
	)
	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			t.T("Usage of %s:\nxtemplate -i input-pattern -k keywords [-f functions] [-o output]\n"),
			os.Args[0],
		)
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
	if *help || len(os.Args) < 5 { // 必填参数: [xtemplate -i xxx -k xxx]
		flag.Usage()
		os.Exit(0)
	}
	return &internal.Param{
		Input:      *input,
		Left:       *left,
		Right:      *right,
		Keyword:    *keywords,
		Function:   *fun,
		OutputFile: *output,
		Debug:      *debug,
	}
}
