package po

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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

		{"multi-line", args{newReader([]string{
			`msgctxt ""`,
			`"hello, "`,
			`"this is ctxt"`,
		})}, &message{msgCTxt: "hello, this is ctxt"}, true},

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

func Test_message_key(t *testing.T) {
	type fields struct {
		msgCTxt string
		msgID   string
		msgID2  string
		msgStr  string
		msgStrN []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"empty", fields{msgCTxt: "", msgID: ""}, "\u0004"},
		{"empty-ctxt", fields{msgCTxt: "", msgID: "id"}, "\u0004id"},
		{"empty-id", fields{msgCTxt: "ctxt", msgID: ""}, "ctxt\u0004"},
		{"kv", fields{msgCTxt: "ctxt", msgID: "id"}, "ctxt\u0004id"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &message{
				msgCTxt: tt.fields.msgCTxt,
				msgID:   tt.fields.msgID,
				msgID2:  tt.fields.msgID2,
				msgStr:  tt.fields.msgStr,
				msgStrN: tt.fields.msgStrN,
			}
			if got := m.key(); got != tt.want {
				t.Errorf("message.key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_isEmpty(t *testing.T) {
	type fields struct {
		msgCTxt string
		msgID   string
		msgID2  string
		msgStr  string
		msgStrN []string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"empty", fields{}, true},
		{"not-empty-ctxt", fields{msgCTxt: "ctxt"}, false},
		{"not-empty-id", fields{msgID: "id"}, false},
		{"not-empty-id2", fields{msgID2: "id2"}, false},
		{"not-empty-str", fields{msgStr: "str"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &message{
				msgCTxt: tt.fields.msgCTxt,
				msgID:   tt.fields.msgID,
				msgID2:  tt.fields.msgID2,
				msgStr:  tt.fields.msgStr,
				msgStrN: tt.fields.msgStrN,
			}
			if got := m.isEmpty(); got != tt.want {
				t.Errorf("message.isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_message_isValid(t *testing.T) {
	type fields struct {
		msgCTxt string
		msgID   string
		msgID2  string
		msgStr  string
		msgStrN []string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"empty", fields{}, false},
		{"empty-header", fields{msgID: ""}, false},
		{"header", fields{msgID: "", msgStr: "k: v"}, true},
		{"header-invalid", fields{msgID: "", msgID2: "id2", msgStr: "k: v"}, false},
		{"entry", fields{msgID: "id", msgStr: "str"}, true},
		{"only-id", fields{msgID: "id"}, false}, // pot file can not used
		{"entry-ctxt", fields{msgCTxt: "txt", msgID: "id", msgStr: "str"}, true},
		{"entry-plural", fields{msgID: "id", msgStrN: []string{"str"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &message{
				msgCTxt: tt.fields.msgCTxt,
				msgID:   tt.fields.msgID,
				msgID2:  tt.fields.msgID2,
				msgStr:  tt.fields.msgStr,
				msgStrN: tt.fields.msgStrN,
			}
			if got := m.isValid(); got != tt.want {
				t.Errorf("message.isValid() = %v, want %v", got, tt.want)
			}
		})
	}
	Convey("nil", t, func() {
		var m *message
		So(m.isValid(), ShouldBeFalse)
		So(m.isEmpty(), ShouldBeTrue)
	})
}
