package parse

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{name: "valid-complex", input: "n==0?0:n==1?1:n==2?2:n%100>=3&&n%100<=10?3:n%100>=11?4:5"},
		{name: "valid-precedence", input: "n==1||n==2&&n==3"},
		{name: "valid-group", input: "(n==1)?0:1"},
		{name: "invalid-prefix-inc-on-int", input: "++1", wantErr: true},
		{name: "invalid-prefix-only-op", input: "++", wantErr: true},
		{name: "invalid-leading-op", input: "*", wantErr: true},
		{name: "invalid-group", input: "(*", wantErr: true},
		{name: "invalid-postfix-on-int", input: "1++", wantErr: true},
		{name: "invalid-binary-missing-rhs", input: "1+", wantErr: true},
		{name: "invalid-ternary-empty", input: "1?:", wantErr: true},
		{name: "invalid-ternary-missing-else", input: "1?0:", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tree, err := ParseInput(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseInput() error = %v, wantErr %v, tree=%v", err, tt.wantErr, tree)
			}
			if err == nil && tree == nil {
				t.Fatalf("ParseInput() returned nil tree")
			}
		})
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		n       int64
		want    int64
		wantErr bool
	}{
		{name: "xor", input: "n^1", n: 1, want: 0},
		{name: "precedence-and-or", input: "n==1||n==2&&n==3", n: 1, want: 1},
		{name: "precedence-and-or-false", input: "n==1||n==2&&n==3", n: 2, want: 0},
		{name: "right-assoc-ternary", input: "n==1?0:n==2?1:2", n: 2, want: 1},
		{name: "invalid-id", input: "nn", n: 1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn, err := String(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("String() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if got := fn(tt.n); got != tt.want {
					t.Fatalf("fn(%d) = %d, want %d", tt.n, got, tt.want)
				}
			}
		})
	}
}
