package po

import (
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

func TestFile_T(t *testing.T) {
	type fields struct {
		headers    map[string]string
		messages   map[string]*message
		totalForms int
		pluralFunc PluralFunc
	}
	type args struct {
		msgID string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"no-translation", fields{}, args{"hello", nil}, "hello"},
		{"translate", fields{messages: map[string]*message{
			"\u0004hello": {
				msgID:  "hello",
				msgStr: "你好",
			},
		}}, args{"hello", nil}, "你好"},
		{"translate-format", fields{messages: map[string]*message{
			"\u0004hello %s": {
				msgID:  "hello %s",
				msgStr: "你好 %s",
			},
		}}, args{"hello %s", nil}, "你好 %s"},
		{"translate-format2", fields{messages: map[string]*message{
			"\u0004hello %s": {
				msgID:  "hello %s",
				msgStr: "你好 %s",
			},
		}}, args{"hello %s", []interface{}{"Tom"}}, "你好 Tom"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				headers:    tt.fields.headers,
				messages:   tt.fields.messages,
				totalForms: tt.fields.totalForms,
				pluralFunc: tt.fields.pluralFunc,
			}
			if got := f.T(tt.args.msgID, tt.args.args...); got != tt.want {
				t.Errorf("File.T() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_N(t *testing.T) {
	type fields struct {
		headers    map[string]string
		messages   map[string]*message
		totalForms int
		pluralFunc PluralFunc
	}
	var apple = fields{
		headers: map[string]string{
			HeaderPluralForms: `nplurals=1; plural=0;`,
		},
		messages: map[string]*message{
			"\u0004one apple": {
				msgID:   "one apple",
				msgStrN: []string{"%d 个苹果"},
			},
		},
		totalForms: 1,
	}
	type args struct {
		msgID       string
		msgIDPlural string
		n           int
		args        []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"no-translation", fields{}, args{"one apple", "%d apples", 1, nil}, "one apple"},
		{"no-translation2", fields{}, args{"one apple", "%d apples", 2, nil}, "%d apples"},
		{"no-translation-format1", fields{}, args{"one apple", "%d apples", 1, []interface{}{1}}, "one apple"},
		{"no-translation-format2", fields{}, args{"one apple", "%d apples", 2, []interface{}{2}}, "2 apples"},

		{"tr", apple, args{"one apple", "%d apples", 1, nil}, "%d 个苹果"},
		{"tr-format", apple, args{"one apple", "%d apples", 1, []interface{}{1}}, "1 个苹果"},
		{"tr2", apple, args{"one apple", "%d apples", 2, nil}, "%d 个苹果"},
		{"tr2-format", apple, args{"one apple", "%d apples", 1, []interface{}{2}}, "2 个苹果"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				headers:    tt.fields.headers,
				messages:   tt.fields.messages,
				totalForms: tt.fields.totalForms,
				pluralFunc: tt.fields.pluralFunc,
			}
			if got := f.N(tt.args.msgID, tt.args.msgIDPlural, tt.args.n, tt.args.args...); got != tt.want {
				t.Errorf("File.N() = %v, want %v", got, tt.want)
			}
			if got := f.N64(tt.args.msgID, tt.args.msgIDPlural, int64(tt.args.n), tt.args.args...); got != tt.want {
				t.Errorf("File.N() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_XN(t *testing.T) {
	type fields struct {
		headers    map[string]string
		messages   map[string]*message
		totalForms int
		pluralFunc PluralFunc
	}
	var exampleFileField = fields{
		headers: map[string]string{
			HeaderPluralForms: `nplurals=1; plural=0;`,
		},
		messages: map[string]*message{
			"File|\u0004Open File": {
				msgID:   "Open File",
				msgStrN: []string{"打开 %d 个文件"},
			},
		},
		totalForms: 1,
	}
	type args struct {
		msgCtxt     string
		msgID       string
		msgIDPlural string
		n           int
		args        []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"no-tr", fields{}, args{"File|", "Open File", "Open Files", 1, nil}, "Open File"},
		{"no-tr2", fields{}, args{"File|", "Open File", "Open Files", 2, nil}, "Open Files"},
		{"tr1", fields{
			headers: map[string]string{
				HeaderPluralForms: `nplurals=1; plural=0;`,
			},
			messages: map[string]*message{
				"File|\u0004Open File": {
					msgID:   "Open File",
					msgStrN: []string{"打开文件"},
				},
			},
			totalForms: 1,
		}, args{"File|", "Open File", "Open Files", 2, nil}, "打开文件"},
		{"tr1", exampleFileField, args{"File|", "Open File", "Open %d Files", 1, nil}, "打开 %d 个文件"},
		{"tr2", exampleFileField, args{"File|", "Open File", "Open %d Files", 1, []interface{}{1}}, "打开 1 个文件"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			po := &File{
				headers:    tt.fields.headers,
				messages:   tt.fields.messages,
				totalForms: tt.fields.totalForms,
				pluralFunc: tt.fields.pluralFunc,
			}
			if got := po.XN(tt.args.msgCtxt, tt.args.msgID, tt.args.msgIDPlural, tt.args.n, tt.args.args...); got != tt.want {
				t.Errorf("File.XN() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{"wp", args{wp}, &File{
			headers: map[string]string{
				"PO-Revision-Date":          "2019-12-08 21:26:25+0000",
				"MIME-Version":              "1.0",
				"Content-Type":              "text/plain; charset=UTF-8",
				"Content-Transfer-Encoding": "8bit",
				"Plural-Forms":              "nplurals=1; plural=0;",
				"X-Generator":               "GlotPress/2.4.0-alpha",
				"Language":                  "zh_CN",
				"Project-Id-Version":        "WordPress - 5.3.x - Development - Administration",
			},
			messages: map[string]*message{
				"\u0004Continue":           {msgID: "Continue", msgStr: "继续"},
				"\u0004Should be improved": {msgID: "Should be improved", msgStr: "有待改进"},
				"\u0004Good":               {msgID: "Good", msgStr: "良好"},
			},
			totalForms: -1,
			pluralFunc: nil,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Logf("headers: \n%v\n%v\n", got.headers, tt.want.headers)
				t.Logf("message: %v", reflect.DeepEqual(got.messages, tt.want.messages))
				t.Errorf("fail")
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
					"ctx txt\u0004hello, world": {
						msgCTxt: "ctx txt",
						msgID:   "hello, world",
						msgStr:  "你好，世界",
					},
				},
				totalForms: -1,
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
					"ctx txt\u0004hello, world": {
						msgCTxt: "ctx txt",
						msgID:   "hello, world",
						msgStr:  "你好，世界",
					},
					"\u0004hello": {
						msgID:  "hello",
						msgStr: "你好",
					},
				},
				totalForms: -1,
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
			messages:   map[string]*message{},
			totalForms: -1,
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
