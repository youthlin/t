package internal

import (
	"os"
	"reflect"
	"testing"
)

func TestParseKeywords(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []Keyword
		wantErr    bool
	}{
		{"empty", args{}, nil, true},
		{"T", args{"T"}, []Keyword{{Name: "T", MsgID: 1}}, false},
		{"T:1", args{"T:1"}, []Keyword{{Name: "T", MsgID: 1}}, false},
		{"N:1,2", args{"N:1,2"}, []Keyword{{Name: "N", MsgID: 1, MsgID2: 2}}, false},
		{"X:1c,2", args{"X:1c,2"}, []Keyword{{Name: "X", MsgCtxt: 1, MsgID: 2}}, false},
		{"XN:1c,2,3", args{"XN:1c,2,3"}, []Keyword{{Name: "XN", MsgCtxt: 1, MsgID: 2, MsgID2: 3}}, false},
		{"T;XN:1c,2,3", args{"T;XN:1c,2,3"}, []Keyword{
			{Name: "T", MsgID: 1},
			{Name: "XN", MsgCtxt: 1, MsgID: 2, MsgID2: 3},
		}, false},

		{"invalid-length", args{"T::"}, nil, true},
		{"invalid-nan", args{"T:a"}, nil, true},
		{"invalid-c-nan", args{"T:ac,2"}, nil, true},
		{"invalid-2-nan", args{"T:1c,a"}, nil, true},
		{"invalid-plural-1", args{"T:a,b"}, nil, true},
		{"invalid-plural-2", args{"T:1,b"}, nil, true},
		{"invalid-3-missing-c", args{"XN:1,2,3"}, nil, true},
		{"invalid-3-c-nan", args{"XN:ac,2,3"}, nil, true},
		{"invalid-3-missing-id", args{"XN:1c,,3"}, nil, true},
		{"invalid-3-id", args{"XN:1c,a,3"}, nil, true},
		{"invalid-3-missing-id2", args{"XN:1c,2,"}, nil, true},
		{"invalid-3-id2", args{"XN:1c,2,a"}, nil, true},

		{"no-arg", args{"XN:"}, nil, true},
		{"much-arg", args{"XN:1,2,3,4"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ParseKeywords(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseKeywords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ParseKeywords() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestWriter(t *testing.T) {
	want, err := os.OpenFile("param_test.go", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Log(err)
	}
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		wantWr  *os.File
		wantErr bool
	}{
		{"empty", args{}, os.Stdout, false},
		{"-", args{"-"}, os.Stdout, false},
		{"file", args{"param_test.go"}, want, false},
		{"can-not-create-if-dir-not-exist", args{"no-such-dir/file.pot"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotWr, err := Writer(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Writer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (tt.wantWr == nil) == (gotWr == nil) {
				return
			}
			if (tt.wantWr == nil) != (gotWr == nil) {
				t.Errorf("Writer() = %v, want %v", gotWr, tt.wantWr)
				return
			}
			gotF, e1 := gotWr.Stat()
			wantF, e2 := tt.wantWr.Stat()

			if e1 != nil || e2 != nil || !os.SameFile(gotF, wantF) {
				t.Errorf("Writer() = %v, want %v", gotWr, tt.wantWr)
			}
		})
	}
}
