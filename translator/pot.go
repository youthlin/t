package translator

import (
	"bytes"
	"fmt"
	"io"
)

const (
	msgctxt      = "msgctxt"
	msgid        = "msgid"
	msgid_plural = "msgid_plural"
	msgstr       = "msgstr"
	msgstrN      = "msgstr[%d]"
)

func writeString(buf *bytes.Buffer, key, content string) {
	buf.WriteString(key + " ")
	for _, line := range split(content, lineThreshold) {
		buf.WriteString(fmt.Sprintf("%q\n", line))
	}
}

func (f *File) SaveAsPot(w io.Writer) error {
	var buf = &bytes.Buffer{}
	for _, entry := range f.SortedEntry() {
		for _, comment := range entry.MsgCmts {
			buf.WriteString(comment)
			buf.WriteString("\n")
		}
		if entry.MsgCtxt != "" {
			writeString(buf, msgctxt, entry.MsgCtxt)
		}
		writeString(buf, msgid, entry.MsgID)

		if entry.MsgID2 == "" {
			if entry.MsgID == "" { // header
				writeString(buf, msgstr, entry.MsgStr)
			} else {
				writeString(buf, msgstr, "")
			}
		} else {
			writeString(buf, msgid_plural, entry.MsgID2)
			writeString(buf, fmt.Sprintf(msgstrN, 0), "")
			writeString(buf, fmt.Sprintf(msgstrN, 1), "")
		}
		buf.WriteString("\n")
	}
	_, err := buf.WriteTo(w)
	return err
}
