package f_test

import (
	"testing"

	"github.com/youthlin/t/f"
)

func TestFormat(t *testing.T) {
	type args struct {
		format string
		args   []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty-nil", args{"", nil}, ""},
		{"empty-empty", args{"", []interface{}{}}, ""},
		{"nonempty-empty", args{"hello", []interface{}{}}, "hello"},
		{"one-apple", args{"one apple", []interface{}{1}}, "one apple"},
		{"2-apples", args{"%d apples", []interface{}{2}}, "2 apples"},
		{"verb-but-no-args", args{"%d apples", []interface{}{}}, "%d apples"},
		{"args-too-few-1", args{"%s have %[1]d apples", []interface{}{1}}, "%!s(int=1) have 1 apples"},
		{"args-too-few-2", args{"%s have %d apples", []interface{}{"Tom"}}, "Tom have %!d(string=) apples"},
		{"position-index", args{"%[2]s have %[1]d apples", []interface{}{2, "Tom"}}, "Tom have 2 apples"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f.Format(tt.args.format, tt.args.args...); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultPlural(t *testing.T) {
	type args struct {
		msgID       string
		msgIDPlural string
		n           int64
		args        []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"en-1-is-singular", args{"singular", "plural", 1, nil}, "singular"},
		{"en-other-is-plural-0", args{"singular", "plural", 0, nil}, "plural"},
		{"en-other-is-plural-n", args{"singular", "plural", 2, nil}, "plural"},
		{"format", args{"one apple", "%d apples", 1, []interface{}{1}}, "one apple"},
		{"format2", args{"one apple", "%d apples", 2, []interface{}{2}}, "2 apples"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f.DefaultPlural(tt.args.msgID, tt.args.msgIDPlural, tt.args.n, tt.args.args...); got != tt.want {
				t.Errorf("defaultPlural() = %v, want %v", got, tt.want)
			}
		})
	}
}
