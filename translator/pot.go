package translator

import (
	"bytes"
	"fmt"
	"io"
)

func (f *File) SaveAsPot(w io.Writer) error {
	var buf bytes.Buffer
	for _, entry := range f.SortedEntry() {
		for _, comment := range entry.MsgCmts {
			buf.WriteString(comment)
			buf.WriteString("\n")
		}
		if entry.MsgCtxt != "" {
			buf.WriteString(fmt.Sprintf("msgctxt %q\n", entry.MsgCtxt))
		}

		buf.WriteString(fmt.Sprintf("msgid %q\n", entry.MsgID))

		if entry.MsgID2 == "" {
			if entry.MsgID == "" { // header
				buf.WriteString(fmt.Sprintf("msgstr %q\n", entry.MsgStr))
			} else {
				buf.WriteString(fmt.Sprintf("msgstr %q\n", ""))
			}
		} else {
			buf.WriteString(fmt.Sprintf("msgid_plural %q\n", entry.MsgID2))
			buf.WriteString(fmt.Sprintf("msgstr[%d] %q\n", 0, ""))
			buf.WriteString(fmt.Sprintf("msgstr[%d] %q\n", 1, ""))
		}
		buf.WriteString("\n")
	}
	_, err := buf.WriteTo(w)
	return err
}
