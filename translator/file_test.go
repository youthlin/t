package translator

import (
	"testing"
)

func TestFile_Lang(t *testing.T) {
	type fields struct {
		entries map[string]*Entry
		headers map[string]string
		plural  *plural
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty", fields{}, ""},
		{"header", fields{headers: map[string]string{HeaderLanguage: "zh_CN"}}, "zh_CN"},
		{"entry", fields{entries: map[string]*Entry{
			key("", ""): {MsgStr: "Language: zh_CN"},
		}}, "zh_CN"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := &File{
				entries: tt.fields.entries,
				headers: tt.fields.headers,
				plural:  tt.fields.plural,
			}
			if got := file.Lang(); got != tt.want {
				t.Errorf("File.Lang() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_T(t *testing.T) {
	type fields struct {
		entries map[string]*Entry
		headers map[string]string
		plural  *plural
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
		{"empty", fields{}, args{"hello", []interface{}{}}, "hello"},
		{"empty-args", fields{}, args{"hello %s", []interface{}{"world"}}, "hello world"},
		{"t", fields{
			entries: map[string]*Entry{
				key("", "hello"): {MsgStr: "你好"},
			},
		}, args{"hello", []interface{}{}}, "你好"},
		{"t-args", fields{
			entries: map[string]*Entry{
				key("", "hello %s"): {MsgStr: "你好 %s"},
			},
		}, args{"hello %s", []interface{}{"world"}}, "你好 world"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := &File{
				entries: tt.fields.entries,
				headers: tt.fields.headers,
				plural:  tt.fields.plural,
			}
			if got := file.X("", tt.args.msgID, tt.args.args...); got != tt.want {
				t.Errorf("File.T() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_N(t *testing.T) {
	type fields struct {
		entries map[string]*Entry
		headers map[string]string
		plural  *plural
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
		{"empty-no-arg-single", fields{}, args{
			"one apple",
			"%d apples",
			1,
			[]interface{}{},
		}, "one apple"},
		{"empty-no-arg-plural", fields{}, args{
			"one apple",
			"%d apples",
			2,
			[]interface{}{},
		}, "%d apples"},
		{"empty-no-arg-single-args", fields{}, args{
			"one apple",
			"%d apples",
			1,
			[]interface{}{1},
		}, "one apple"},
		{"empty-no-arg-plural-args", fields{}, args{
			"one apple",
			"%d apples",
			2,
			[]interface{}{2},
		}, "2 apples"},

		{
			"no-arg-no-plural-header",
			fields{
				entries: map[string]*Entry{
					key("", "one apple"): {MsgStrN: []string{"%d 个苹果"}},
				},
			},
			args{
				"one apple",
				"%d apples",
				1,
				[]interface{}{},
			},
			"one apple",
		},
		{
			"with-arg-no-plural-header",
			fields{
				entries: map[string]*Entry{
					key("", "one apple"): {MsgStrN: []string{"%d 个苹果"}},
				},
			},
			args{
				"one apple",
				"%d apples",
				2,
				[]interface{}{2},
			},
			"2 apples",
		},

		{
			"no-arg-single",
			fields{
				entries: map[string]*Entry{
					key("", "one apple"): {MsgStrN: []string{"%d 个苹果"}},
				},
				headers: map[string]string{HeaderPluralForms: "nplurals=1;plural=0;"},
			},
			args{
				"one apple",
				"%d apples",
				1,
				[]interface{}{},
			},
			"%d 个苹果",
		},
		{
			"no-arg-plural",
			fields{
				entries: map[string]*Entry{
					key("", "one apple"): {MsgStrN: []string{"%d 个苹果"}},
				},
				headers: map[string]string{HeaderPluralForms: "nplurals=1;plural=0;"},
			},
			args{
				"one apple",
				"%d apples",
				2,
				[]interface{}{},
			},
			"%d 个苹果",
		},
		{
			"with-arg",
			fields{
				entries: map[string]*Entry{
					key("", "one apple"): {MsgStrN: []string{"%d 个苹果"}},
				},
				headers: map[string]string{HeaderPluralForms: "nplurals=1;plural=0;"},
			},
			args{
				"one apple",
				"%d apples",
				1,
				[]interface{}{1},
			},
			"1 个苹果",
		},

		{
			"invalid-plural",
			fields{
				entries: map[string]*Entry{
					key("", "one apple"): {MsgStrN: []string{"%d 个苹果"}},
				},
				headers: map[string]string{HeaderPluralForms: "nplurals=1;plural=1;"},
			},
			args{
				"one apple",
				"%d apples",
				1,
				[]interface{}{1},
			},
			"one apple",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := &File{
				entries: tt.fields.entries,
				headers: tt.fields.headers,
				plural:  tt.fields.plural,
			}
			if got := file.XN64("", tt.args.msgID, tt.args.msgIDPlural, int64(tt.args.n), tt.args.args...); got != tt.want {
				t.Errorf("File.N() = %v, want %v", got, tt.want)
			}
		})
	}
}
