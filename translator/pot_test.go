package translator

import (
	"bytes"
	"testing"
)

func entriesMap(entries []*Entry) map[string]*Entry {
	m := make(map[string]*Entry)
	for _, e := range entries {
		m[e.key()] = e
	}
	return m
}

func TestFile_SaveAsPot(t *testing.T) {
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
msgstr ""

`, false},
		{"cmt-ctx-plural", fields{[]*Entry{
			{
				comments: []string{"# translators comment", "#: path/to/source"},
				msgCtxt:  "ctx",
				msgID:    "one apple",
				msgID2:   "%d apples",
				msgStrN:  []string{"%d 个苹果"},
			},
		}}, `# translators comment
#: path/to/source
msgctxt "ctx"
msgid "one apple"
msgid_plural "%d apples"
msgstr[0] ""
msgstr[1] ""

`, false},
		{"header-cmt-ctx-plural", fields{[]*Entry{
			{
				msgID: "",
				msgStr: `Project-Id-Version: MyProject
Language: zh_CN
Content-Type: text/plain; charset=UTF-8
Content-Transfer-Encoding: 8bit
Plural-Forms: nplurals=1; plural=0;
`,
			},
			{
				comments: []string{"# translators comment", "#: path/to/source"},
				msgCtxt:  "ctx",
				msgID:    "one apple",
				msgID2:   "%d apples",
				msgStrN:  []string{"%d 个苹果"},
			},
		}}, `msgid ""
msgstr "Project-Id-Version: MyProject\nLanguage: zh_CN\nContent-Type: text/plain; charset=UTF-8\nContent-Transfer-Encoding: 8bit\nPlural-Forms: nplurals=1; plural=0;\n"

# translators comment
#: path/to/source
msgctxt "ctx"
msgid "one apple"
msgid_plural "%d apples"
msgstr[0] ""
msgstr[1] ""

`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				entries: entriesMap(tt.fields.entries),
			}
			w := &bytes.Buffer{}
			if err := f.SaveAsPot(w); (err != nil) != tt.wantErr {
				t.Errorf("File.SaveAsPot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("File.SaveAsPot() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
