package translator

// Entry 一个翻译条目
type Entry struct {
	comments []string
	msgCtxt  string
	msgID    string
	msgID2   string
	msgStr   string
	msgStrN  []string
}

// key 一个翻译条目的 key
func (e *Entry) key() string {
	return key(e.msgCtxt, e.msgID)
}

// isHeader 返回该条目是否是一个 header 条目
func (e *Entry) isHeader() bool {
	return e.msgID == ""
}

// isValid 是否合法的条目
func (e *Entry) isValid() bool {
	// header entry: msgid == ""
	return e != nil && (e.msgStr != "" || len(e.msgStrN) > 0)
}
