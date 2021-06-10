package files

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewEmptyFile(t *testing.T) {
	tests := []struct {
		name string
		want *File
	}{
		{"not-nil", &File{
			headers:    make(map[string]string),
			messages:   make(map[string]*Message),
			totalForms: -1,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEmptyFile(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEmptyFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_SetLang(t *testing.T) {
	type fields struct {
		headers    map[string]string
		messages   map[string]*Message
		lang       string
		totalForms int
		pluralFunc PluralFunc
	}
	type args struct {
		lang string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *File
	}{
		{"ok", fields{}, args{"zh_CN"}, &File{lang: "zh_CN"}},
		{"normalize", fields{}, args{"zh-CN"}, &File{lang: "zh_CN"}},
		{"invalid", fields{}, args{""}, &File{lang: "en_US"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				headers:    tt.fields.headers,
				messages:   tt.fields.messages,
				lang:       tt.fields.lang,
				totalForms: tt.fields.totalForms,
				pluralFunc: tt.fields.pluralFunc,
			}
			if got := f.SetLang(tt.args.lang); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File.SetLang() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_SetHeaders(t *testing.T) {
	type fields struct {
		headers    map[string]string
		messages   map[string]*Message
		lang       string
		totalForms int
		pluralFunc PluralFunc
	}
	type args struct {
		h map[string]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *File
	}{
		{"nil", fields{}, args{}, &File{}},
		{"non-nil", fields{}, args{map[string]string{"k": "v"}}, &File{headers: map[string]string{"k": "v"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				headers:    tt.fields.headers,
				messages:   tt.fields.messages,
				lang:       tt.fields.lang,
				totalForms: tt.fields.totalForms,
				pluralFunc: tt.fields.pluralFunc,
			}
			if got := f.SetHeaders(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File.SetHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_SetMessages(t *testing.T) {
	type fields struct {
		headers    map[string]string
		messages   map[string]*Message
		lang       string
		totalForms int
		pluralFunc PluralFunc
	}
	type args struct {
		m map[string]*Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *File
	}{
		{"", fields{}, args{}, &File{}},
		{"", fields{}, args{map[string]*Message{"a": nil}}, &File{messages: map[string]*Message{"a": nil}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				headers:    tt.fields.headers,
				messages:   tt.fields.messages,
				lang:       tt.fields.lang,
				totalForms: tt.fields.totalForms,
				pluralFunc: tt.fields.pluralFunc,
			}
			if got := f.SetMessages(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File.SetMessages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_GetHeader(t *testing.T) {
	type fields struct {
		headers    map[string]string
		messages   map[string]*Message
		lang       string
		totalForms int
		pluralFunc PluralFunc
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  bool
	}{
		{"no", fields{}, args{}, "", false},
		{"has", fields{headers: map[string]string{"k": "v"}}, args{"k"}, "v", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				headers:    tt.fields.headers,
				messages:   tt.fields.messages,
				lang:       tt.fields.lang,
				totalForms: tt.fields.totalForms,
				pluralFunc: tt.fields.pluralFunc,
			}
			got, got1 := f.GetHeader(tt.args.key)
			if got != tt.want {
				t.Errorf("File.GetHeader() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("File.GetHeader() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFile_GetTotalForms(t *testing.T) {
	type fields struct {
		headers    map[string]string
		messages   map[string]*Message
		lang       string
		totalForms int
		pluralFunc PluralFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"has-field", fields{totalForms: 1}, 1},
		{"from-header", fields{headers: map[string]string{
			HeaderPluralForms: "nplurals=1;plural=0;",
		}, totalForms: -1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			po := &File{
				headers:    tt.fields.headers,
				messages:   tt.fields.messages,
				lang:       tt.fields.lang,
				totalForms: tt.fields.totalForms,
				pluralFunc: tt.fields.pluralFunc,
			}
			if got := po.GetTotalForms(); got != tt.want {
				t.Errorf("File.GetTotalForms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_GetPluralFunc(t *testing.T) {
	Convey("GetPluralFunc", t, func() {
		Convey("invalid", func() {
			var f = File{}
			var fn = f.GetPluralFunc()
			So(fn(1), ShouldEqual, -1)
		})
		Convey("from-header", func() {
			var f = File{headers: map[string]string{
				HeaderPluralForms: "nplurals=1;plural=0;",
			}}
			var fn = f.GetPluralFunc()
			So(fn(1), ShouldEqual, 0)
		})
	})
}

func TestFile_getPluralArr(t *testing.T) {
	type fields struct {
		headers    map[string]string
		messages   map[string]*Message
		lang       string
		totalForms int
		pluralFunc PluralFunc
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
		want1  bool
	}{
		{"", fields{}, nil, false},
		{"ok", fields{headers: map[string]string{
			HeaderPluralForms: "nplurals=1;plural=n!=1;",
		}}, []string{"1", "n!=1"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			po := &File{
				headers:    tt.fields.headers,
				messages:   tt.fields.messages,
				lang:       tt.fields.lang,
				totalForms: tt.fields.totalForms,
				pluralFunc: tt.fields.pluralFunc,
			}
			got, got1 := po.getPluralArr()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("File.getPluralArr() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("File.getPluralArr() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestFile_AddMessage(t *testing.T) {
	type fields struct {
		headers    map[string]string
		messages   map[string]*Message
		lang       string
		totalForms int
		pluralFunc PluralFunc
	}
	type args struct {
		m *Message
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"nil-msg", fields{}, args{}, true},
		{
			"invalid-header",
			fields{headers: make(map[string]string), messages: make(map[string]*Message)},
			args{&Message{MsgStr: " "}},
			true,
		},
		{
			"header",
			fields{headers: make(map[string]string), messages: make(map[string]*Message)},
			args{&Message{MsgStr: `"k: v"`}},
			false,
		},
		{
			"msg",
			fields{headers: make(map[string]string), messages: make(map[string]*Message)},
			args{&Message{MsgID: "id", MsgStr: "str"}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				headers:    tt.fields.headers,
				messages:   tt.fields.messages,
				lang:       tt.fields.lang,
				totalForms: tt.fields.totalForms,
				pluralFunc: tt.fields.pluralFunc,
			}
			if err := f.AddMessage(tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("File.AddMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_key(t *testing.T) {
	type args struct {
		ctxt string
		id   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"both-empty", args{}, split},
		{"ctxt-empty", args{id: "id"}, split + "id"},
		{"id-empty", args{ctxt: "ctxt"}, "ctxt" + split},
		{"ctxt-id", args{ctxt: "ctxt", id: "id"}, "ctxt" + split + "id"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := key(tt.args.ctxt, tt.args.id); got != tt.want {
				t.Errorf("key() = %v, want %v", got, tt.want)
			}
		})
	}
}
