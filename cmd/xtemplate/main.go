package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"github.com/youthlin/t"
	"github.com/youthlin/t/cmd/xtemplate/internal"
)

const defaultVersion = "dev"

// Version 允许通过 -ldflags 显式覆盖版本号。
// 若未覆盖，则优先从 Go build info 中读取模块版本或 vcs revision，
// 避免每次发布前都手工修改源码里的版本字符串。
var Version string = defaultVersion

//go:embed lang
var embedLangs embed.FS

// initTranslation 初始化 xtemplate 自己的提示文案翻译。
// LANG_PATH 为空时回退到内嵌语言包，避免工作目录变化导致找不到 ./lang。
func initTranslation() {
	path, ok := os.LookupEnv("LANG_PATH")
	if ok && strings.TrimSpace(path) != "" {
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
			os.Exit(2)
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

// buildParam 解析命令行参数并做最基本的必填校验。
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
			versionString())
		os.Exit(0)
	}
	if *help {
		flag.Usage()
		os.Exit(0)
	}
	if strings.TrimSpace(*input) == "" || strings.TrimSpace(*keywords) == "" {
		flag.Usage()
		fmt.Fprintln(os.Stderr, t.T("flags -i and -k are required"))
		os.Exit(2)
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

func versionString() string {
	if Version != "" && Version != defaultVersion {
		return Version
	}
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return defaultVersion
	}
	if info.Main.Version != "" && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	var (
		revision string
		time     string
		dirty    bool
	)
	for _, item := range info.Settings {
		switch item.Key {
		case "vcs.revision":
			revision = item.Value
		case "vcs.time":
			time = item.Value
		case "vcs.modified":
			dirty = item.Value == "true"
		}
	}
	if revision == "" {
		return defaultVersion
	}
	if len(revision) > 12 {
		revision = revision[:12]
	}
	parts := []string{revision}
	if dirty {
		parts = append(parts, "dirty")
	}
	if time != "" {
		parts = append(parts, time)
	}
	return strings.Join(parts, "-")
}
