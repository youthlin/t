package po

import (
	"testing"
)

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
