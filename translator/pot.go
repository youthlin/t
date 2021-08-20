package translator

import (
	"bytes"
	"fmt"
	"io"
)

func (f *File) SaveAsPot(w io.Writer) error {
	var buf bytes.Buffer
	for _, msg := range f.entries {
		for _, comment := range msg.comments {
			buf.WriteString(comment)
			buf.WriteString("\n")
		}
		if msg.msgCtxt != "" {
			buf.WriteString(fmt.Sprintf("msgctxt %q\n", msg.msgCtxt))
		}

		buf.WriteString(fmt.Sprintf("msgid %q\n", msg.msgID))

		if msg.msgID2 == "" {
			if msg.msgID == "" { // header
				buf.WriteString(fmt.Sprintf("msgstr %q\n", msg.msgStr))
			} else {
				buf.WriteString(fmt.Sprintf("msgstr %q\n", ""))
			}
		} else {
			buf.WriteString(fmt.Sprintf("msgid_plural %q\n", msg.msgID2))
			buf.WriteString(fmt.Sprintf("msgstr[%d] %q\n", 0, ""))
			buf.WriteString(fmt.Sprintf("msgstr[%d] %q\n", 1, ""))
		}
		buf.WriteString("\n")
	}
	_, err := buf.WriteTo(w)
	return err
}
