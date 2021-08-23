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
				msgID:  "",
				msgStr: "Project-Id-Version: MyProject\n",
			},
		}}, `msgid ""
msgstr "Project-Id-Version: MyProject\n"

`, false},
		{"header-2", fields{[]*Entry{
			{
				msgID: "",
				msgStr: `Project-Id-Version: MyProject
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
				comments: []string{"# translators comment", "#: path/to/source"},
				msgID:    "hello",
				msgStr:   "你好",
			},
		}}, `# translators comment
#: path/to/source
msgid "hello"
msgstr "你好"

`, false},
		{"cmt-ctx-plural", fields{[]*Entry{
			{
				msgID:  "",
				msgStr: "Project-Id-Version: MyProject\n",
			},
			{
				comments: []string{"# translators comment", "#: path/to/source"},
				msgCtxt:  "ctx",
				msgID:    "one apple",
				msgID2:   "%d apples",
				msgStrN:  []string{"%d 个苹果"},
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
			msgID:  "hello",
			msgStr: "你好",
		}, true},
		{"simple-id-strN", args{newReader([]string{`msgid "hello"`, `msgstr[0] "你好"`})}, &Entry{
			msgID:   "hello",
			msgStrN: []string{"你好"},
		}, true},

		{"simple-id-str-two-entry", args{newReader([]string{`msgid "hello"`, `msgstr "你好"`, `msgid "entry 2`})}, &Entry{
			msgID:  "hello",
			msgStr: "你好",
		}, false},

		{"simple-cmt", args{newReader([]string{`# bla bla`, `msgstr "你好"`})}, &Entry{
			comments: []string{"# bla bla"},
			msgID:    "",
			msgStr:   "你好",
		}, true},
		{"cmt-is-entry-start", args{newReader([]string{`# bla bla`, `msgstr "你好"`, "#abc"})}, &Entry{
			comments: []string{"# bla bla"},
			msgID:    "",
			msgStr:   "你好",
		}, false},

		{"simple-ctxt", args{newReader([]string{`msgctxt "ctxt"`, `msgstr "你好"`})}, &Entry{
			msgCtxt: "ctxt",
			msgStr:  "你好",
		}, true},
		{"unquote-ctxt", args{newReader([]string{`msgctxt ctxt`, `msgstr "你好"`})}, nil, true}, // quote err
		{"split-ctxt", args{newReader([]string{`msgctxt "ctxt"`, `msgctxt "你好"`})}, &Entry{
			msgCtxt: "ctxt",
		}, false},

		{"split-msg_plural", args{newReader([]string{`msgid_plural "ctxt"`, `msgid_plural "你好"`})}, &Entry{
			msgID2: "ctxt",
		}, false},
		{"unquote-msg_plural", args{newReader([]string{`msgid_plural id2`, `msgid_plural "你好"`})}, nil, true},
		{"unquote-msg_id", args{newReader([]string{`msgid id2`, `msgid_plural "你好"`})}, nil, true},

		{"msg_strN", args{newReader([]string{`msgstr[0] "str0"`, `msgstr[1] "str1"`})}, &Entry{
			msgStrN: []string{"str0", "str1"},
		}, true},
		{"split-msg_strN", args{newReader([]string{`msgstr[0] "str0"`, `msgstr "str"`})}, &Entry{
			msgStrN: []string{"str0"},
		}, false},

		{"unquote-msg_strN", args{newReader([]string{`msgstr[0] str0`})}, nil, true},

		{"msg_str", args{newReader([]string{`msgstr "str"`})}, &Entry{
			msgStr: "str",
		}, true},
		{"msg_str", args{newReader([]string{`msgstr str`})}, nil, true},
		{"split-msg_str", args{newReader([]string{`msgstr "str"`, `msgstr[0] "str0"`})}, &Entry{
			msgStr: "str",
		}, false},

		{"multi-line-ctxt", args{newReader([]string{`msgctxt "line1"`, `"line2"`})}, &Entry{
			msgCtxt: "line1line2",
		}, true},
		{"multi-line-id", args{newReader([]string{`msgid "line1"`, `"line2"`})}, &Entry{
			msgID: "line1line2",
		}, true},
		{"multi-line-id2", args{newReader([]string{`msgid_plural "line1"`, `"line2"`})}, &Entry{
			msgID2: "line1line2",
		}, true},
		{"multi-line-str", args{newReader([]string{`msgstr "line1"`, `"line2"`})}, &Entry{
			msgStr: "line1line2",
		}, true},
		{"multi-line-strN", args{newReader([]string{`msgstr[0] "line1"`, `"line2"`})}, &Entry{
			msgStrN: []string{"line1line2"},
		}, true},
		{"unquote-multi-line-ctxt", args{newReader([]string{`msgctxt "line1"`, `"line2`})}, nil, true},
		{"unexpected-multi-line", args{newReader([]string{`msgctxt "line1"`, `line2`})}, nil, true},

		{"case-1", args{newReader([]string{`msgid ""`, `msgstr ""`, `"Project-Id-Version: t\n"`})}, &Entry{
			msgStr: "Project-Id-Version: t\n",
		}, true},
		{"case-2", args{newReader([]string{
			`#, c-format`,
			`msgctxt "Project|"`,
			`msgid "Open One"`,
			`msgid_plural "Open %d"`,
			`msgstr[0] "打开 %d 个工程"`})}, &Entry{
			comments: []string{"#, c-format"},
			msgCtxt:  "Project|",
			msgID:    "Open One",
			msgID2:   "Open %d",
			msgStrN:  []string{"打开 %d 个工程"},
		}, true},
		{"case-3-invalid-entry", args{newReader([]string{
			`#, c-format`,
			`msgctxt "Project|"`,
			`msgid "Open One"`,
			`msgid_plural "Open %d"`,
			`msgstr "打开"`,
			`msgstr[0] "打开 %d 个工程"`})}, &Entry{
			comments: []string{"#, c-format"},
			msgCtxt:  "Project|",
			msgID:    "Open One",
			msgID2:   "Open %d",
			msgStr:   "打开",
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
				comments: []string{"#: lang_test.go:22 lang_test.go:23 main_test.go:37"},
				msgID:    "Hello, World",
				msgStr:   "你好，世界",
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
				comments: []string{"#: lang_test.go:22 lang_test.go:23 main_test.go:37"},
				msgID:    "Hello, World",
				msgStr:   "你好，世界",
			},
			key("", "one apple"): {
				msgID:   "one apple",
				msgID2:  "%d apples",
				msgStrN: []string{"%d 个苹果"},
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
				comments: []string{"#: lang_test.go:22 lang_test.go:23 main_test.go:37"},
				msgID:    "Hello, World",
				msgStr:   "你好，世界",
			},
			key("", "one apple"): {
				msgID:   "one apple",
				msgID2:  "%d apples",
				msgStrN: []string{"%d 个苹果"},
			},
			key("verb", "Post"): {
				msgCtxt: "verb",
				msgID:   "Post",
				msgStr:  "发布",
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
