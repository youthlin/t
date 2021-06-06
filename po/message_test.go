package po

import (
	"io"
	"reflect"
	"testing"

	"github.com/pkg/errors"
)

func Test_readMessage(t *testing.T) {
	type args struct {
		r *reader
	}
	tests := []struct {
		name    string
		args    args
		want    *message
		wantErr bool
	}{
		{"nil", args{newReader(nil)}, nil, true},
		{"empty", args{newReader([]string{})}, nil, true},
		{"blank", args{newReader([]string{""})}, nil, true},
		{"only-comment", args{newReader([]string{"# hello"})}, nil, true},
		{"only-ctx: empty", args{newReader([]string{`msgctxt ""`})}, &message{}, true}, //EOF
		{"only-ctx: non empty", args{newReader([]string{`msgctxt "txt"`})},
			&message{msgCTxt: "txt"}, true},
		{"only-id", args{newReader([]string{`msgid "id"`})}, &message{msgID: "id"}, true},
		{"only-id2", args{newReader([]string{`msgid_plural "id2"`})}, &message{msgID2: "id2"}, true},
		{"only-str", args{newReader([]string{`msgstr "str"`})}, &message{msgStr: "str"}, true},
		{"only-str2", args{newReader([]string{`msgstr[0] "str0"`})},
			&message{msgStrN: []string{"str0"}}, true},
		{"not-any", args{newReader([]string{`foo "str"`})}, nil, true},
		{"unknown-prefix", args{newReader([]string{`msgid "id"`, `foo "str"`})}, nil, true},

		{"ok", args{newReader([]string{`msgid "id"`, `msgstr "str"`})},
			&message{msgID: "id", msgStr: "str"}, true},
		{"ok2", args{newReader([]string{`msgid "id"`, `msgstr "str"`, `msgid "id"`, `msgstr "str"`})},
			&message{msgID: "id", msgStr: "str"}, false}, // read one, left one, so not reached EOF
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readMessage(tt.args.r)
			if err != nil {
				if errors.Is(err, io.EOF) {
					// t.Logf("EOF")
				} else {
					// t.Logf("%+v", err)
				}
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("readMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
