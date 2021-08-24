package translator

import (
	"bytes"
	"reflect"
	"testing"
)

func TestFile_SaveAsPo(t *testing.T) {
	type fields struct {
		entries []*Entry
	}
	tests := []struct {
		name    string
		fields  fields
		wantW   string
		wantErr bool
	}{
		{"empty", fields{}, "", false},
		{"header-only", fields{[]*Entry{
			{
				MsgID:  "",
				MsgStr: "Project-Id-Version: MyProject\n",
			},
		}}, `msgid ""
msgstr "Project-Id-Version: MyProject\n"

`, false},
		{"header-2", fields{[]*Entry{
			{
				MsgID: "",
				MsgStr: `Project-Id-Version: MyProject
Language: zh_CN
Content-Type: text/plain; charset=UTF-8
Content-Transfer-Encoding: 8bit
Plural-Forms: nplurals=1; plural=0;
`,
			},
		}}, `msgid ""
msgstr "Project-Id-Version: MyProject\nLanguage: zh_CN\nContent-Type: text/plain; charset=UTF-8\nContent-Transfer-Encoding: 8bit\nPlural-Forms: nplurals=1; plural=0;\n"

`, false},
		{"with-cmt", fields{[]*Entry{
			{
				MsgCmts: []string{"# translators comment", "#: path/to/source"},
				MsgID:   "hello",
				MsgStr:  "你好",
			},
		}}, `# translators comment
#: path/to/source
msgid "hello"
msgstr "你好"

`, false},
		{"cmt-ctx-plural", fields{[]*Entry{
			{
				MsgID:  "",
				MsgStr: "Project-Id-Version: MyProject\n",
			},
			{
				MsgCmts: []string{"# translators comment", "#: path/to/source"},
				MsgCtxt: "ctx",
				MsgID:   "one apple",
				MsgID2:  "%d apples",
				MsgStrN: []string{"%d 个苹果"},
			},
		}}, `msgid ""
msgstr "Project-Id-Version: MyProject\n"

# translators comment
#: path/to/source
msgctxt "ctx"
msgid "one apple"
msgid_plural "%d apples"
msgstr[0] "%d 个苹果"

`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				entries: entriesMap(tt.fields.entries),
			}
			w := &bytes.Buffer{}
			if err := f.SaveAsPo(w); (err != nil) != tt.wantErr {
				t.Errorf("File.SaveAsPo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("File.SaveAsPo() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_readEntry(t *testing.T) {
	type args struct {
		r *reader
	}
	tests := []struct {
		name    string
		args    args
		want    *Entry
		wantErr bool
	}{
		{"empty", args{newReader([]string{})}, &Entry{}, true},   // EOF
		{"blank", args{newReader([]string{""})}, &Entry{}, true}, // EOF

		{"simple-id-str", args{newReader([]string{`msgid "hello"`, `msgstr "你好"`})}, &Entry{
			MsgID:  "hello",
			MsgStr: "你好",
		}, true},
		{"simple-id-strN", args{newReader([]string{`msgid "hello"`, `msgstr[0] "你好"`})}, &Entry{
			MsgID:   "hello",
			MsgStrN: []string{"你好"},
		}, true},

		{"simple-id-str-two-entry", args{newReader([]string{`msgid "hello"`, `msgstr "你好"`, `msgid "entry 2`})}, &Entry{
			MsgID:  "hello",
			MsgStr: "你好",
		}, false},

		{"simple-cmt", args{newReader([]string{`# bla bla`, `msgstr "你好"`})}, &Entry{
			MsgCmts: []string{"# bla bla"},
			MsgID:   "",
			MsgStr:  "你好",
		}, true},
		{"cmt-is-entry-start", args{newReader([]string{`# bla bla`, `msgstr "你好"`, "#abc"})}, &Entry{
			MsgCmts: []string{"# bla bla"},
			MsgID:   "",
			MsgStr:  "你好",
		}, false},

		{"simple-ctxt", args{newReader([]string{`msgctxt "ctxt"`, `msgstr "你好"`})}, &Entry{
			MsgCtxt: "ctxt",
			MsgStr:  "你好",
		}, true},
		{"unquote-ctxt", args{newReader([]string{`msgctxt ctxt`, `msgstr "你好"`})}, nil, true}, // quote err
		{"split-ctxt", args{newReader([]string{`msgctxt "ctxt"`, `msgctxt "你好"`})}, &Entry{
			MsgCtxt: "ctxt",
		}, false},

		{"split-msg_plural", args{newReader([]string{`msgid_plural "ctxt"`, `msgid_plural "你好"`})}, &Entry{
			MsgID2: "ctxt",
		}, false},
		{"unquote-msg_plural", args{newReader([]string{`msgid_plural id2`, `msgid_plural "你好"`})}, nil, true},
		{"unquote-msg_id", args{newReader([]string{`msgid id2`, `msgid_plural "你好"`})}, nil, true},

		{"msg_strN", args{newReader([]string{`msgstr[0] "str0"`, `msgstr[1] "str1"`})}, &Entry{
			MsgStrN: []string{"str0", "str1"},
		}, true},
		{"split-msg_strN", args{newReader([]string{`msgstr[0] "str0"`, `msgstr "str"`})}, &Entry{
			MsgStrN: []string{"str0"},
		}, false},

		{"unquote-msg_strN", args{newReader([]string{`msgstr[0] str0`})}, nil, true},

		{"msg_str", args{newReader([]string{`msgstr "str"`})}, &Entry{
			MsgStr: "str",
		}, true},
		{"msg_str", args{newReader([]string{`msgstr str`})}, nil, true},
		{"split-msg_str", args{newReader([]string{`msgstr "str"`, `msgstr[0] "str0"`})}, &Entry{
			MsgStr: "str",
		}, false},

		{"multi-line-ctxt", args{newReader([]string{`msgctxt "line1"`, `"line2"`})}, &Entry{
			MsgCtxt: "line1line2",
		}, true},
		{"multi-line-id", args{newReader([]string{`msgid "line1"`, `"line2"`})}, &Entry{
			MsgID: "line1line2",
		}, true},
		{"multi-line-id2", args{newReader([]string{`msgid_plural "line1"`, `"line2"`})}, &Entry{
			MsgID2: "line1line2",
		}, true},
		{"multi-line-str", args{newReader([]string{`msgstr "line1"`, `"line2"`})}, &Entry{
			MsgStr: "line1line2",
		}, true},
		{"multi-line-strN", args{newReader([]string{`msgstr[0] "line1"`, `"line2"`})}, &Entry{
			MsgStrN: []string{"line1line2"},
		}, true},
		{"unquote-multi-line-ctxt", args{newReader([]string{`msgctxt "line1"`, `"line2`})}, nil, true},
		{"unexpected-multi-line", args{newReader([]string{`msgctxt "line1"`, `line2`})}, nil, true},

		{"case-1", args{newReader([]string{`msgid ""`, `msgstr ""`, `"Project-Id-Version: t\n"`})}, &Entry{
			MsgStr: "Project-Id-Version: t\n",
		}, true},
		{"case-2", args{newReader([]string{
			`#, c-format`,
			`msgctxt "Project|"`,
			`msgid "Open One"`,
			`msgid_plural "Open %d"`,
			`msgstr[0] "打开 %d 个工程"`})}, &Entry{
			MsgCmts: []string{"#, c-format"},
			MsgCtxt: "Project|",
			MsgID:   "Open One",
			MsgID2:  "Open %d",
			MsgStrN: []string{"打开 %d 个工程"},
		}, true},
		{"case-3-invalid-entry", args{newReader([]string{
			`#, c-format`,
			`msgctxt "Project|"`,
			`msgid "Open One"`,
			`msgid_plural "Open %d"`,
			`msgstr "打开"`,
			`msgstr[0] "打开 %d 个工程"`})}, &Entry{
			MsgCmts: []string{"#, c-format"},
			MsgCtxt: "Project|",
			MsgID:   "Open One",
			MsgID2:  "Open %d",
			MsgStr:  "打开",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readEntry(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("readEntry() error = %+v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readEntry() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestReadPo(t *testing.T) {
	type args struct {
		content []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *File
		wantErr bool
	}{
		{"nil", args{}, nil, true},
		{"empty", args{[]byte("")}, nil, true},
		{"one-entry", args{[]byte(`#: lang_test.go:22 lang_test.go:23 main_test.go:37
msgid "Hello, World"
msgstr "你好，世界"`)}, &File{entries: map[string]*Entry{
			key("", "Hello, World"): {
				MsgCmts: []string{"#: lang_test.go:22 lang_test.go:23 main_test.go:37"},
				MsgID:   "Hello, World",
				MsgStr:  "你好，世界",
			},
		}}, false},
		{"two-entry", args{[]byte(`#: lang_test.go:22 lang_test.go:23 main_test.go:37
msgid "Hello, World"
msgstr "你好，世界"

msgid "one apple"
msgid_plural "%d apples"
msgstr[0] "%d 个苹果"
`)}, &File{entries: map[string]*Entry{
			key("", "Hello, World"): {
				MsgCmts: []string{"#: lang_test.go:22 lang_test.go:23 main_test.go:37"},
				MsgID:   "Hello, World",
				MsgStr:  "你好，世界",
			},
			key("", "one apple"): {
				MsgID:   "one apple",
				MsgID2:  "%d apples",
				MsgStrN: []string{"%d 个苹果"},
			},
		}}, false},
		{"3-entry", args{[]byte(`#: lang_test.go:22 lang_test.go:23 main_test.go:37
msgid "Hello, World"
msgstr "你好，世界"

msgid "one apple"
msgid_plural "%d apples"
msgstr[0] "%d 个苹果"

msgctxt "verb"
msgid "Post"
msgstr "发布"
`)}, &File{entries: map[string]*Entry{
			key("", "Hello, World"): {
				MsgCmts: []string{"#: lang_test.go:22 lang_test.go:23 main_test.go:37"},
				MsgID:   "Hello, World",
				MsgStr:  "你好，世界",
			},
			key("", "one apple"): {
				MsgID:   "one apple",
				MsgID2:  "%d apples",
				MsgStrN: []string{"%d 个苹果"},
			},
			key("verb", "Post"): {
				MsgCtxt: "verb",
				MsgID:   "Post",
				MsgStr:  "发布",
			},
		}}, false},
		{"error", args{[]byte(`msgid hallo`)}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadPo(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadPo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadPo() = %v, want %v", got, tt.want)
			}
		})
	}
}
