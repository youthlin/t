package po

import (
	"reflect"
	"testing"

	"github.com/youthlin/t/files"
)

func Test_readMessage(t *testing.T) {
	type args struct {
		r *reader
	}
	tests := []struct {
		name    string
		args    args
		want    *files.Message
		wantErr bool
	}{
		{"nil", args{newReader(nil)}, nil, true},
		{"empty", args{newReader([]string{})}, nil, true},
		{"blank", args{newReader([]string{""})}, nil, true},
		{"only-comment", args{newReader([]string{"# hello"})}, nil, true},
		{"only-ctx: empty", args{newReader([]string{`msgctxt ""`})}, &files.Message{}, true}, //EOF
		{"only-ctx: non empty", args{newReader([]string{`msgctxt "txt"`})},
			&files.Message{MsgCtxt: "txt"}, true},
		{"only-id", args{newReader([]string{`msgid "id"`})}, &files.Message{MsgID: "id"}, true},
		{"only-id2", args{newReader([]string{`msgid_plural "id2"`})}, &files.Message{MsgID2: "id2"}, true},
		{"only-str", args{newReader([]string{`msgstr "str"`})}, &files.Message{MsgStr: "str"}, true},
		{"only-str2", args{newReader([]string{`msgstr[0] "str0"`})},
			&files.Message{MsgStrN: []string{"str0"}}, true},
		{"not-any", args{newReader([]string{`foo "str"`})}, nil, true},
		{"unknown-prefix", args{newReader([]string{`msgid "id"`, `foo "str"`})}, nil, true},

		{"ok", args{newReader([]string{`msgid "id"`, `msgstr "str"`})},
			&files.Message{MsgID: "id", MsgStr: "str"}, true},
		{"ok2", args{newReader([]string{`msgid "id"`, `msgstr "str"`, `msgid "id"`, `msgstr "str"`})},
			&files.Message{MsgID: "id", MsgStr: "str"}, false}, // read one, left one, so not reached EOF

		{"multi-line", args{newReader([]string{
			`msgctxt ""`,
			`"hello, "`,
			`"this is ctxt"`,
		})}, &files.Message{MsgCtxt: "hello, this is ctxt"}, true},

		{"quote-err", args{newReader([]string{
			`msgctxt ""`,
			`"hello, `, // this line missing `"`
			`"this is ctxt"`,
		})}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readMessage(tt.args.r)
			// if err != nil {
			// 	if errors.Is(err, io.EOF) {
			// 		 t.Logf("EOF")
			// 	} else {
			// 		 t.Logf("%+v", err)
			// 	}
			// }
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
