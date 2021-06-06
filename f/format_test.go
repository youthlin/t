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
