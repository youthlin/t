package translator

// Entry 一个翻译条目
type Entry struct {
	MsgCmts []string
	MsgCtxt string
	MsgID   string
	MsgID2  string
	MsgStr  string
	MsgStrN []string
}

// key 一个翻译条目的 key
func (e *Entry) key() string {
	return key(e.MsgCtxt, e.MsgID)
}

// isHeader 返回该条目是否是一个 header 条目
func (e *Entry) isHeader() bool {
	return e.MsgID == ""
}

// isValid 是否合法的条目
func (e *Entry) isValid() bool {
	// header entry: msgid == ""
	return e != nil && (e.MsgStr != "" || len(e.MsgStrN) > 0)
}
