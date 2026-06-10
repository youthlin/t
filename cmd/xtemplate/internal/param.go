package internal

import (
	"fmt"
	"html/template"
	"os"
	"strings"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/youthlin/t"
	"github.com/youthlin/t/translator"
)

// Param 输入参数
type Param struct {
	Input      string
	Left       string
	Right      string
	Keyword    string
	Function   string
	OutputFile string
	Debug      bool
}

// debugPrint print if is debug mode
func (p *Param) debugPrint(format string, args ...interface{}) {
	if p.Debug {
		fmt.Printf(format+"\n", args...)
	}
}

// Context parameters and result
type Context struct {
	*Param
	Keywords  []Keyword
	Functions template.FuncMap
	Output    *os.File
	hasPlural bool
	entries   map[string]*translator.Entry
}

func newCtx(param *Param) (*Context, error) {
	ctx := &Context{
		Param:     param,
		Functions: make(template.FuncMap),
		entries:   make(map[string]*translator.Entry),
	}
	kw, err := ParseKeywords(param.Keyword)
	if err != nil {
		return nil, err
	}
	ctx.Keywords = kw
	for _, k := range kw {
		ctx.Functions[k.Name] = noopFun
	}
	wr, err := Writer(param.OutputFile)
	if err != nil {
		return nil, err
	}
	ctx.Output = wr

	if ctx.Function != "" {
		fun := strings.Split(ctx.Function, ",")
		for _, name := range fun {
			ctx.Functions[name] = noopFun
		}
	}
	return ctx, nil
}

// Add add a message entry
func (ctx *Context) Add(entry *translator.Entry) error {
	plural := isPlural(entry)
	key := entry.Key()
	pre, ok := ctx.entries[key]
	if ok {
		if isPlural(pre) != plural {
			return errors.Errorf(t.T("msgid '%v' is used without plural and with plural.\nLine    =%v\nPrevious=%v"),
				entry.MsgID, entry.MsgCmts, pre.MsgCmts)
		}
		pre.MsgCmts = append(pre.MsgCmts, entry.MsgCmts...)
	} else {
		ctx.entries[key] = entry
		if plural {
			ctx.hasPlural = true
		}
	}
	return nil
}

// Write write pot file to output
func (ctx *Context) Write() error {
	pot := ctx.pot()
	return pot.SaveAsPot(ctx.Output)
}

func (ctx *Context) pot() *translator.File {
	pot := new(translator.File)
	pot.AddEntry(ctx.header())
	for _, e := range ctx.entries {
		pot.AddEntry(e)
	}
	return pot
}

func (ctx *Context) header() *translator.Entry {
	e := new(translator.Entry)
	e.MsgCmts = []string{
		"# SOME DESCRIPTIVE TITLE.",
		"# Copyright (C) YEAR THE PACKAGE'S COPYRIGHT HOLDER",
		"# This file is distributed under the same license as the PACKAGE package.",
		"# FIRST AUTHOR <EMAIL@ADDRESS>, YEAR.",
		"#",
		"#, fuzzy",
	}
	headers := []string{
		"Project-Id-Version: PACKAGE VERSION",
		"Report-Msgid-Bugs-To: ",
		fmt.Sprintf("POT-Creation-Date: %v", time.Now().Format("2006-01-02 15:04:05-0700")),
		"PO-Revision-Date: YEAR-MO-DA HO:MI+ZONE",
		"Last-Translator: FULL NAME <EMAIL@ADDRESS>",
		"Language-Team: LANGUAGE <LL@li.org>",
		"Language: ",
		"MIME-Version: 1.0",
		"Content-Type: text/plain; charset=CHARSET",
		"Content-Transfer-Encoding: 8bit",
		"X-Created-By: xtemplate(https://github.com/youthlin/t/tree/main/cmd/xtemplate)",
	}
	if ctx.hasPlural {
		headers = append(headers, "Plural-Forms: nplurals=INTEGER; plural=EXPRESSION;")
	}
	headers = append(headers, fmt.Sprintf("X-Xtemplate-Input: %v", ctx.Input))
	headers = append(headers, fmt.Sprintf("X-Xtemplate-Left: %v", ctx.Left))
	headers = append(headers, fmt.Sprintf("X-Xtemplate-Right: %v", ctx.Right))
	headers = append(headers, fmt.Sprintf("X-Xtemplate-Keywords: %v", ctx.Keyword))
	headers = append(headers, fmt.Sprintf("X-Xtemplate-Functions: %v", ctx.Function))
	headers = append(headers, fmt.Sprintf("X-Xtemplate-Output: %v", ctx.OutputFile))
	headers = append(headers, "")
	e.MsgStr = strings.Join(headers, "\n")
	return e
}
