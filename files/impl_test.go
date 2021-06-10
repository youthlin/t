package files

import "testing"

func TestFile_T(t *testing.T) {
	type fields struct {
		headers    map[string]string
		messages   map[string]*Message
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
		{"translate", fields{messages: map[string]*Message{
			key("", "hello"): {
				MsgID:  "hello",
				MsgStr: "你好",
			},
		}}, args{"hello", nil}, "你好"},
		{"translate-format", fields{messages: map[string]*Message{
			"\u0004hello %s": {
				MsgID:  "hello %s",
				MsgStr: "你好 %s",
			},
		}}, args{"hello %s", nil}, "你好 %s"},
		{"translate-format2", fields{messages: map[string]*Message{
			"\u0004hello %s": {
				MsgID:  "hello %s",
				MsgStr: "你好 %s",
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
		messages   map[string]*Message
		totalForms int
		pluralFunc PluralFunc
	}
	var apple = fields{
		headers: map[string]string{
			HeaderPluralForms: `nplurals=1; plural=0;`,
		},
		messages: map[string]*Message{
			"\u0004one apple": {
				MsgID:   "one apple",
				MsgStrN: []string{"%d 个苹果"},
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
		messages   map[string]*Message
		totalForms int
		pluralFunc PluralFunc
	}
	var exampleFileField = fields{
		headers: map[string]string{
			HeaderPluralForms: `nplurals=1; plural=0;`,
		},
		messages: map[string]*Message{
			"File|\u0004Open File": {
				MsgID:   "Open File",
				MsgStrN: []string{"打开 %d 个文件"},
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
			messages: map[string]*Message{
				"File|\u0004Open File": {
					MsgID:   "Open File",
					MsgStrN: []string{"打开文件"},
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
