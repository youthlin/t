package po

import (
	"errors"
	"io"
	"reflect"
	"testing"
)

const wp = `# Translation of WordPress - 5.3.x - Development - Administration in Chinese (China)
# This file is distributed under the same license as the WordPress - 5.3.x - Development - Administration package.
msgid ""
msgstr ""
"PO-Revision-Date: 2019-12-08 21:26:25+0000\n"
"MIME-Version: 1.0\n"
"Content-Type: text/plain; charset=UTF-8\n"
"Content-Transfer-Encoding: 8bit\n"
"Plural-Forms: nplurals=1; plural=0;\n"
"X-Generator: GlotPress/2.4.0-alpha\n"
"Language: zh_CN\n"
"Project-Id-Version: WordPress - 5.3.x - Development - Administration\n"

#: wp-admin/includes/media.php:1697 wp-admin/upgrade.php:77
#: wp-admin/upgrade.php:153
msgid "Continue"
msgstr "继续"

#: wp-admin/js/site-health.js:134
msgid "Should be improved"
msgstr "有待改进"

#: wp-admin/js/site-health.js:129
msgid "Good"
msgstr "良好"
`

func TestParse(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		want    *File
		wantErr bool
	}{
		{"empty", args{""}, newEmptyFile(), false},
		{"", args{wp}, newEmptyFile(), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseLines(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name    string
		args    args
		want    *File
		wantErr bool
	}{
		{"nil", args{nil}, newEmptyFile(), false},
		{"empty", args{[]string{}}, newEmptyFile(), false},
		{
			"ok",
			args{[]string{
				`msgctxt "ctx txt"`,
				`msgid "hello, world"`,
				`msgstr "你好，世界"`,
			}},
			&File{
				headers: map[string]string{},
				messages: map[string]*message{
					"hello, world": {
						msgCTxt: "ctx txt",
						msgID:   "hello, world",
						msgStr:  "你好，世界",
					},
				},
			},
			false,
		},

		{
			"ok2",
			args{[]string{
				`msgctxt "ctx txt"`,
				`msgid "hello, world"`,
				`msgstr "你好，世界"`,

				`msgid "hello"`,
				`msgstr "你好"`,
			}},
			&File{
				headers: map[string]string{},
				messages: map[string]*message{
					"hello, world": {
						msgCTxt: "ctx txt",
						msgID:   "hello, world",
						msgStr:  "你好，世界",
					},
					"hello": {
						msgID:  "hello",
						msgStr: "你好",
					},
				},
			},
			false,
		},

		{"invalid-msg-missing-id", args{[]string{
			`msgctxt "ctx txt"`,
			`msgstr "你好，世界"`,
		}}, nil, true},
		{"invalid-msg-missing-str", args{[]string{
			`msgctxt "ctx txt"`,
			`msgid "hello, world"`,
		}}, nil, true},

		{"header", args{[]string{
			`msgid ""`,
			`msgstr ""`,
			`"MIME-Version: 1.0\n"`,
			`"PO-Revision-Date: 2019-12-08 21:26:25+0000\n"`,
			`"Content-Type: text/plain; charset=UTF-8\n"`,
			`"Plural-Forms: nplurals=1; plural=0;\n"`,
		}}, &File{
			headers: map[string]string{
				"MIME-Version":     "1.0",
				"PO-Revision-Date": "2019-12-08 21:26:25+0000",
				"Content-Type":     "text/plain; charset=UTF-8",
				"Plural-Forms":     "nplurals=1; plural=0;",
			},
			messages: map[string]*message{},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseLines(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLines() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_readMessage(t *testing.T) {
	type args struct {
		r *reader
	}
	tests := []struct {
		name    string
		args    args
		want    *message
		wantErr bool
	}{
		{"nil", args{newReader(nil)}, nil, true},
		{"empty", args{newReader([]string{})}, nil, true},
		{"blank", args{newReader([]string{""})}, nil, true},
		{"only-comment", args{newReader([]string{"# hello"})}, nil, true},
		{"only-ctx: empty", args{newReader([]string{`msgctxt ""`})}, &message{}, true}, //EOF
		{"only-ctx: non empty", args{newReader([]string{`msgctxt "txt"`})},
			&message{msgCTxt: "txt"}, true},
		{"only-id", args{newReader([]string{`msgid "id"`})}, &message{msgID: "id"}, true},
		{"only-id2", args{newReader([]string{`msgid_plural "id2"`})}, &message{msgID2: "id2"}, true},
		{"only-str", args{newReader([]string{`msgstr "str"`})}, &message{msgStr: "str"}, true},
		{"only-str2", args{newReader([]string{`msgstr[0] "str0"`})},
			&message{msgStrN: []string{"str0"}}, true},
		{"not-any", args{newReader([]string{`foo "str"`})}, nil, true},
		{"unknown-prefix", args{newReader([]string{`msgid "id"`, `foo "str"`})}, nil, true},

		{"ok", args{newReader([]string{`msgid "id"`, `msgstr "str"`})},
			&message{msgID: "id", msgStr: "str"}, true},
		{"ok2", args{newReader([]string{`msgid "id"`, `msgstr "str"`, `msgid "id"`, `msgstr "str"`})},
			&message{msgID: "id", msgStr: "str"}, false}, // read one, left one, so not reached EOF
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readMessage(tt.args.r)
			if err != nil {
				if errors.Is(err, io.EOF) {
					t.Logf("EOF")
				} else {
					t.Logf("%+v", err)
				}
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("readMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readLine(t *testing.T) {
	type args struct {
		r *reader
	}
	tests := []struct {
		name     string
		args     args
		wantLine string
		wantErr  bool
	}{
		{"nil", args{newReader(nil)}, "", true},
		{"empty", args{newReader([]string{})}, "", true},
		{"blank", args{newReader([]string{""})}, "", true},
		{"comment", args{newReader([]string{"# some text"})}, "", true},
		{"blank comment", args{newReader([]string{"", "# some text"})}, "", true},
		{"content", args{newReader([]string{"some text"})}, "some text", false},
		{"many content", args{newReader([]string{"some text", "hello"})}, "some text", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotLine, err := readLine(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("readLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotLine != tt.wantLine {
				t.Errorf("readLine() = %v, want %v", gotLine, tt.wantLine)
			}
		})
	}
}

func Test_unquote(t *testing.T) {
	type args struct {
		line   string
		prefix string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"empty0", args{``, ""}, "", true},
		{"empty1", args{`""`, ""}, "", false},
		{"err", args{`hello`, ""}, "", true},
		{"ok", args{`"hello"`, ""}, "hello", false},
		{"prefix-line-empty", args{``, "msgid"}, "", true},
		{"prefix-empty2", args{`""`, "msgid"}, "", false},
		{"prefix-not-prefix", args{`"hello"`, "msgid"}, "hello", false},
		{"prefix-not-prefix2", args{`"msgid hello"`, "msgid"}, "msgid hello", false},
		{"prefix-ok", args{`msgid "hello"`, "msgid"}, "hello", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unquote(tt.args.line, tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("unquote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("unquote() = %v, want %v", got, tt.want)
			}
		})
	}
}
