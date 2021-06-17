package internal

import (
	"fmt"
	"io"
	"time"

	"github.com/pkg/errors"
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
			return errors.Errorf(t.T("msgid '%v' is used without plural and with plural."), msg.msgID)
		}
	}
	return nil
}

// write write pot to Writer
func (p *pot) write(wr io.Writer) {
	p.writeHeader(wr)
	p.writeMessage(wr)
}

// writeHeader write pot header
func (p *pot) writeHeader(wr io.Writer) {
	now := time.Now().Format("2006-01-02 15:04:05-0700")
	wr.Write([]byte(fmt.Sprintf(`# SOME DESCRIPTIVE TITLE.
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
"X-Created-By: xtemplate(https://github.com/youthlin/t/tree/main/cmd/xtemplate)\n"`, now)))
	if p.hasPlural {
		wr.Write([]byte(`"Plural-Forms: nplurals=INTEGER; plural=EXPRESSION;\n"`))
	}
}

// writeMessage write pot messages
func (p *pot) writeMessage(wr io.Writer) {
	for _, msg := range p.messages {
		msg.write(wr)
	}
}

type message struct {
	comments []string // 提取注释
	line     []string // 代码行号
	msgCtxt  string   // 翻译上下文
	msgID    string   // 原文
	msgID2   string   // 原文复数
}

func newMessage(line string) *message {
	return &message{line: []string{line}}
}
func (m *message) addComment(cmt string) *message {
	m.comments = append(m.comments, cmt)
	return m
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

func (m *message) write(wr io.Writer) {
	if m == nil || m.msgID == "" {
		return
	}
	for _, comment := range m.comments {
		wr.Write([]byte(fmt.Sprintf("#. %v\n", comment)))
	}
	for _, line := range m.line {
		wr.Write([]byte(fmt.Sprintf("#: %v\n", line)))
	}
	if m.msgCtxt != "" {
		wr.Write([]byte(fmt.Sprintf("msgctxt %q\n", m.msgCtxt)))
	}
	if m.msgID != "" {
		wr.Write([]byte(fmt.Sprintf("msgid %q\n", m.msgID)))
	}
	if m.msgID2 == "" {
		wr.Write([]byte(fmt.Sprintf("msgstr %q\n", "")))
	} else {
		wr.Write([]byte(fmt.Sprintf("msgid_plural %q\n", m.msgID2)))
		wr.Write([]byte(fmt.Sprintf("msgstr[0] %q\n", "")))
		wr.Write([]byte(fmt.Sprintf("msgstr[1] %q\n", "")))
	}
	wr.Write([]byte("\n"))
}
