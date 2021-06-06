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
