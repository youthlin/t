package main

import (
	"fmt"
	"io"
)

func key(msgCtxt, msgID string) string {
	return msgCtxt + "\u0004" + msgID
}

// https://www.gnu.org/software/gettext/manual/html_node/PO-Files.html

type pot struct {
	messages map[string]*message
}

type message struct {
	comments []string // 提取注释
	line     []string // 代码行号
	msgCtxt  string   // 翻译上下文
	msgID    string   // 原文
	msgID2   string   // 原文复数
}

func (m *message) addComment(c string) {
	m.line = append(m.line, c)
}

func (m *message) write(wr io.Writer) {
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
		wr.Write([]byte(fmt.Sprintf("msgstr %q\n", "")))
	}
	if m.msgID2 != "" {
		wr.Write([]byte(fmt.Sprintf("msgid_plural %q\n", m.msgID2)))
		wr.Write([]byte(fmt.Sprintf("msgstr[0] %q\n", "")))
		wr.Write([]byte(fmt.Sprintf("msgstr[1] %q\n", "")))
	}
	wr.Write([]byte("\n"))
}
