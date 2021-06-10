package files

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Message_Key(t *testing.T) {
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
			m := &Message{
				MsgCtxt: tt.fields.msgCTxt,
				MsgID:   tt.fields.msgID,
				MsgID2:  tt.fields.msgID2,
				MsgStr:  tt.fields.msgStr,
				MsgStrN: tt.fields.msgStrN,
			}
			if got := m.Key(); got != tt.want {
				t.Errorf("message.key() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Message_IsEmpty(t *testing.T) {
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
			m := &Message{
				MsgCtxt: tt.fields.msgCTxt,
				MsgID:   tt.fields.msgID,
				MsgID2:  tt.fields.msgID2,
				MsgStr:  tt.fields.msgStr,
				MsgStrN: tt.fields.msgStrN,
			}
			if got := m.IsEmpty(); got != tt.want {
				t.Errorf("message.isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Message_IsValid(t *testing.T) {
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
			m := &Message{
				MsgCtxt: tt.fields.msgCTxt,
				MsgID:   tt.fields.msgID,
				MsgID2:  tt.fields.msgID2,
				MsgStr:  tt.fields.msgStr,
				MsgStrN: tt.fields.msgStrN,
			}
			if got := m.IsValid(); got != tt.want {
				t.Errorf("message.isValid() = %v, want %v", got, tt.want)
			}
		})
	}
	Convey("nil", t, func() {
		var m *Message
		So(m.IsValid(), ShouldBeFalse)
		So(m.IsEmpty(), ShouldBeTrue)
	})
}
