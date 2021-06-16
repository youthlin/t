package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/youthlin/t"
)

// -C C++
// -o output file
// -k keywords
// --from-code 因为文件中有中文所以需要指定
// -c 提取注释
// xgettext -C -o testdata/pot_test.pot -kT -kX:1c,2 -kN:1,2 -kXN:1c,2,3 --from-code=UTF-8 -cTRANSLATORS: pot_test.go

// # SOME DESCRIPTIVE TITLE.
// # Copyright (C) YEAR THE PACKAGE'S COPYRIGHT HOLDER
// # This file is distributed under the same license as the PACKAGE package.
// # FIRST AUTHOR <EMAIL@ADDRESS>, YEAR.
// #
// #, fuzzy
// msgid ""
// msgstr ""
// "Project-Id-Version: PACKAGE VERSION\n"
// "Report-Msgid-Bugs-To: \n"
// "POT-Creation-Date: 2021-06-16 17:21+0800\n"
// "PO-Revision-Date: YEAR-MO-DA HO:MI+ZONE\n"
// "Last-Translator: FULL NAME <EMAIL@ADDRESS>\n"
// "Language-Team: LANGUAGE <LL@li.org>\n"
// "Language: \n"
// "MIME-Version: 1.0\n"
// "Content-Type: text/plain; charset=CHARSET\n"
// "Content-Transfer-Encoding: 8bit\n"
// "Plural-Forms: nplurals=INTEGER; plural=EXPRESSION;\n"
//
// #. TRANSLATORS: msg_id is message id
// #: pot_test.go:57 pot_test.go:58
// msgid "msg_id"
// msgstr ""
//
// #: pot_test.go:59
// msgctxt "ctxt"
// msgid "msg_id_x"
// msgstr ""
//
// #: pot_test.go:60
// msgid ""
// "msg\n"
// "id"
// msgstr ""
//
// #: pot_test.go:61
// msgid "msg\tid"
// msgstr ""
//
// #: pot_test.go:62
// msgid "one apple"
// msgid_plural "%v apples"
// msgstr[0] ""
// msgstr[1] ""

func TestT(testT *testing.T) {
	// 如果没有出现复数，就不会有复数表达式这个 Header: Plural-Forms
	// 警告 同一个key同时用于单复数
	t.T("msg_id")
	// TRANSLATORS: msg_id is message id
	t.T("msg_id")
	t.X("ctxt", "msg_id_x")
	t.T("msg\nid") // 会换行
	t.T("msg\tid")
	t.N("one apple", "%v apples", 1) // 复数是 msgstr[0] 和 msgstr[1]
}

func Test_pot_add(t *testing.T) {

}

func Test_message_write(t *testing.T) {
	type fields struct {
		comments []string
		line     []string
		msgCtxt  string
		msgID    string
		msgID2   string
	}
	tests := []struct {
		name   string
		fields fields
		wantWr string
	}{
		{"empty", fields{}, ""},
		{"simple", fields{msgID: "msg_id"},
			strings.Join([]string{`msgid "msg_id"`, `msgstr ""`, ``, ``}, "\n")},
		{"plural", fields{msgID: "msg_id", msgID2: "msg_plural"},
			strings.Join([]string{`msgid "msg_id"`, `msgid_plural "msg_plural"`, `msgstr[0] ""`, `msgstr[1] ""`, ``, ``}, "\n")},
		{"context", fields{msgCtxt: "ctxt", msgID: "msg_id"},
			strings.Join([]string{`msgctxt "ctxt"`, `msgid "msg_id"`, `msgstr ""`, ``, ``}, "\n")},
		{"context-plural", fields{msgCtxt: "ctxt", msgID: "msg_id", msgID2: "msg_plural"},
			strings.Join([]string{`msgctxt "ctxt"`, `msgid "msg_id"`, `msgid_plural "msg_plural"`, `msgstr[0] ""`, `msgstr[1] ""`, ``, ``}, "\n")},
		{"line", fields{line: []string{"testdata/base.tmpl:10"}, msgID: "hello, world"},
			strings.Join([]string{`#: testdata/base.tmpl:10`, `msgid "hello, world"`, `msgstr ""`, ``, ``}, "\n")},
		{"comment", fields{comments: []string{"TRANSLATORS: xxx"}, msgID: "hello, world"},
			strings.Join([]string{`#. TRANSLATORS: xxx`, `msgid "hello, world"`, `msgstr ""`, ``, ``}, "\n")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &message{
				comments: tt.fields.comments,
				line:     tt.fields.line,
				msgCtxt:  tt.fields.msgCtxt,
				msgID:    tt.fields.msgID,
				msgID2:   tt.fields.msgID2,
			}
			wr := &bytes.Buffer{}
			m.write(wr)
			if gotWr := wr.String(); gotWr != tt.wantWr {
				t.Errorf("message.write() = %v, want %v", gotWr, tt.wantWr)
			}
		})
	}
}
