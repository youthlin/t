package files

type Message struct {
	MsgCmt  string
	MsgCtxt string
	MsgID   string
	MsgID2  string
	MsgStr  string
	MsgStrN []string
}

func (m *Message) Key() string {
	return key(m.MsgCtxt, m.MsgID)
}
func (m *Message) IsEmpty() bool {
	return m == nil || m.MsgCtxt == "" && m.MsgID == "" &&
		m.MsgID2 == "" && m.MsgStr == "" && len(m.MsgStrN) == 0
}

func (m *Message) IsValid() bool {
	if m == nil {
		return false
	}
	if m.MsgID == "" { // header
		return m.MsgCtxt == "" && m.MsgID2 == "" && m.MsgStr != "" && len(m.MsgStrN) == 0
	}
	return m.MsgStr != "" || len(m.MsgStrN) != 0
}
