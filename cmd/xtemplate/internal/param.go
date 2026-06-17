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

// debugPrint 只在 debug 模式下输出调试日志。
func (p *Param) debugPrint(format string, args ...any) {
	if p.Debug {
		fmt.Printf(format+"\n", args...)
	}
}

// Context 保存一次提取任务运行期间的上下文。
// 除了解析参数外，也负责缓存 keyword/function 映射、收集 entry 以及输出目标。
type Context struct {
	*Param
	Keywords  []Keyword
	Functions template.FuncMap
	Output    *os.File
	hasPlural bool
	entries   map[string]*translator.Entry
}

// newCtx 根据命令行参数构造运行上下文，并完成 keyword/function/output 的初始化。
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
			name = strings.TrimSpace(name)
			if name == "" {
				continue
			}
			ctx.Functions[name] = noopFun
		}
	}
	return ctx, nil
}

// Add 合并一条抽取结果；相同 msgid 会合并来源注释。
func (ctx *Context) Add(entry *translator.Entry) error {
	plural := entry.IsPlural
	key := entry.Key()
	pre, ok := ctx.entries[key]
	if ok {
		if pre.IsPlural != plural {
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

// Write 将当前收集到的条目写入 POT。
func (ctx *Context) Write() error {
	pot := ctx.pot()
	return pot.SaveAsPot(ctx.Output)
}

// Close 在输出目标是普通文件时关闭它；stdout/stderr 无需关闭。
func (ctx *Context) Close() error {
	if ctx.Output == nil || ctx.Output == os.Stdout || ctx.Output == os.Stderr {
		return nil
	}
	return ctx.Output.Close()
}

// pot 组装最终输出的 pot 文件对象。
func (ctx *Context) pot() *translator.File {
	pot := new(translator.File)
	pot.AddEntry(ctx.header())
	for _, e := range ctx.entries {
		pot.AddEntry(e)
	}
	return pot
}

// header 生成 pot 头信息，同时把本次提取的关键参数记录进去，便于后续排查问题。
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
