package translator

import (
	"reflect"
	"testing"
)

func Test_split(t *testing.T) {
	type args struct {
		long      string
		threshold int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"empty", args{"", 10}, []string{""}},
		{"no-need-split", args{"0123456789", 10}, []string{"0123456789"}},
		{"can-not-split", args{"0123456789", 5}, []string{"0123456789"}},
		{"split-by-blank", args{" 0123456789", 5}, []string{" ", "0123456789"}},
		{"split-by-blank", args{"0 123456789", 5}, []string{"0 ", "123456789"}},
		{"split-by-blank", args{"01 23456789", 5}, []string{"01 ", "23456789"}},
		{"split-by-blank", args{"012 3456789", 5}, []string{"012 ", "3456789"}},
		{"split-by-blank", args{"0123 456789", 5}, []string{"0123 ", "456789"}},
		{"split-by-blank", args{"01234 56789", 5}, []string{"01234 ", "56789"}},
		{"split-by-blank", args{"012345 6789", 5}, []string{"012345 ", "6789"}},
		{"split-by-blank", args{"0123456 789", 5}, []string{"0123456 ", "789"}},
		{"split-by-blank", args{"01234567 89", 5}, []string{"01234567 ", "89"}},
		{"split-by-blank", args{"012345678 9", 5}, []string{"012345678 ", "9"}},
		{"split-by-blank", args{"0123456789 ", 5}, []string{"0123456789 "}},
		{"split-by-blank", args{"0123456789  ", 5}, []string{"0123456789  "}},

		{"split-by-newline", args{"0123456789\n", 5}, []string{"0123456789\n"}},

		{"split-by-newline", args{"0123456789\n\n", 5}, []string{"", "0123456789\n", "\n"}},
		{"split-by-newline", args{"012345678\n9", 5}, []string{"", "012345678\n", "9"}},
		{"split-by-newline", args{"01234567\n89", 5}, []string{"", "01234567\n", "89"}},
		{"split-by-newline", args{"0123456\n789", 5}, []string{"", "0123456\n", "789"}},
		{"split-by-newline", args{"012345\n6789", 5}, []string{"", "012345\n", "6789"}},
		{"split-by-newline", args{"01234\n56789", 5}, []string{"", "01234\n", "56789"}},
		{"split-by-newline", args{"0123\n456789", 5}, []string{"", "0123\n", "456789"}},
		{"split-by-newline", args{"012\n3456789", 5}, []string{"", "012\n", "3456789"}},
		{"split-by-newline", args{"01\n23456789", 5}, []string{"", "01\n", "23456789"}},
		{"split-by-newline", args{"0\n123456789", 5}, []string{"", "0\n", "123456789"}},
		{"split-by-newline", args{"\n0123456789", 5}, []string{"", "\n", "0123456789"}},

		{"split-by-newline", args{"01234\n56789", 15}, []string{"", "01234\n", "56789"}},

		{"split-by-newline", args{"01234\n56789\n01234", 5}, []string{"", "01234\n", "56789\n", "01234"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := split(tt.args.long, tt.args.threshold); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("split() = %q, want %q", got, tt.want)
			}
		})
	}
}
