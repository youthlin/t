package internal

import (
	"fmt"
	"io"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/youthlin/t"
)

func key(msgCtxt, msgID string) string {
	return msgCtxt + "\u0004" + msgID
}

// https://www.gnu.org/software/gettext/manual/html_node/PO-Files.html

// pot the pot file
type pot struct {
	hasPlural bool
	messages  map[string]*message
}

func newPot() *pot {
	return &pot{
		messages: make(map[string]*message),
	}
}

// add add message to pot
func (p *pot) add(msg *message) error {
	if msg.plural() {
		p.hasPlural = true
	}
	previous, ok := p.messages[msg.key()]
	if ok {
		if previous.plural() != msg.plural() {
			return errors.Errorf(t.T("msgid '%v' is used without plural and with plural.\nLine    =%v\nPrevious=%v"),
				msg.msgID, msg.line, previous.line)
		}
		previous.addLine(msg.line[0])
	} else {
		p.messages[msg.key()] = msg
	}
	return nil
}

// write write pot to Writer
func (p *pot) write(ctx *Context, wr io.Writer) error {
	if err := p.writeHeader(ctx, wr); err != nil {
		return err
	}
	if err := p.writeMessage(wr); err != nil {
		return err
	}
	return nil
}

// writeHeader write pot header
func (p *pot) writeHeader(ctx *Context, wr io.Writer) error {
	now := time.Now().Format("2006-01-02 15:04:05-0700")
	plural := ""
	if p.hasPlural {
		plural = `"Plural-Forms: nplurals=INTEGER; plural=EXPRESSION;\n"` + "\n"
	}
	_, err := wr.Write([]byte(fmt.Sprintf(`# SOME DESCRIPTIVE TITLE.
# Copyright (C) YEAR THE PACKAGE'S COPYRIGHT HOLDER
# This file is distributed under the same license as the PACKAGE package.
# FIRST AUTHOR <EMAIL@ADDRESS>, YEAR.
#
#, fuzzy
msgid ""
msgstr ""
"Project-Id-Version: PACKAGE VERSION\n"
"Report-Msgid-Bugs-To: \n"
"POT-Creation-Date: %v\n"
"PO-Revision-Date: YEAR-MO-DA HO:MI+ZONE\n"
"Last-Translator: FULL NAME <EMAIL@ADDRESS>\n"
"Language-Team: LANGUAGE <LL@li.org>\n"
"Language: \n"
"MIME-Version: 1.0\n"
"Content-Type: text/plain; charset=CHARSET\n"
"Content-Transfer-Encoding: 8bit\n"
"X-Created-By: xtemplate(https://github.com/youthlin/t/tree/main/cmd/xtemplate)\n"
%v"X-Xtemplate-Input: %v\n"
"X-Xtemplate-Left: %v\n"
"X-Xtemplate-Right: %v\n"
"X-Xtemplate-Keywords: %v\n"
"X-Xtemplate-Functions: %v\n"
"X-Xtemplate-Output: %v\n"

`, now, plural, replaceQuote(ctx.Input), replaceQuote(ctx.Left),
		replaceQuote(ctx.Right), replaceQuote(ctx.Keyword),
		replaceQuote(ctx.Function), replaceQuote(ctx.OutputFile))))
	return err
}

// replaceQuote 将双引号转义
func replaceQuote(str string) string {
	str = fmt.Sprintf("%q", str)
	str = str[1 : len(str)-1]
	return str
}

// writeMessage write pot messages
func (p *pot) writeMessage(wr io.Writer) error {
	for _, msg := range p.messages {
		if err := msg.write(wr); err != nil {
			return err
		}
	}
	return nil
}

type message struct {
	line    []string // 代码行号
	msgCtxt string   // 翻译上下文
	msgID   string   // 原文
	msgID2  string   // 原文复数
}

func newMessage(line string) *message {
	return &message{line: []string{line}}
}

func (m *message) addLine(txt string) *message {
	m.line = append(m.line, txt)
	return m
}
func (m *message) setCtxt(txt string) *message {
	m.msgCtxt = txt
	return m
}
func (m *message) setMsgID(txt string) *message {
	m.msgID = txt
	return m
}
func (m *message) setMsgPlural(txt string) *message {
	m.msgID2 = txt
	return m
}

func (m *message) key() string {
	return key(m.msgCtxt, m.msgID)
}
func (m *message) plural() bool {
	return m.msgID2 != ""
}

func (m *message) write(wr io.Writer) error {
	if m == nil || m.msgID == "" {
		return nil
	}
	for _, line := range m.line {
		if _, err := wr.Write([]byte(fmt.Sprintf("#: %v\n", line))); err != nil {
			return err
		}
	}
	if m.msgCtxt != "" {
		if _, err := wr.Write([]byte(fmt.Sprintf("msgctxt %q\n", m.msgCtxt))); err != nil {
			return err
		}
	}
	if m.msgID != "" {
		if _, err := wr.Write([]byte(fmt.Sprintf("msgid %q\n", m.msgID))); err != nil {
			return err
		}
	}
	if m.msgID2 == "" {
		if _, err := wr.Write([]byte(fmt.Sprintf("msgstr %q\n", ""))); err != nil {
			return err
		}
	} else {
		if _, err := wr.Write([]byte(fmt.Sprintf("msgid_plural %q\n", m.msgID2))); err != nil {
			return err
		}
		if _, err := wr.Write([]byte(fmt.Sprintf("msgstr[0] %q\n", ""))); err != nil {
			return err
		}
		if _, err := wr.Write([]byte(fmt.Sprintf("msgstr[1] %q\n", ""))); err != nil {
			return err
		}
	}
	if _, err := wr.Write([]byte("\n")); err != nil {
		return err
	}
	return nil
}
