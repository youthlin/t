package po

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	comment = "#"
	msgCtxt = "msgctxt"
	msgID   = "msgid"
	msgID2  = "msgid_plural"
	msgStr  = "msgstr"
	msgStrN = "msgstr["
	quote   = `"`
)

func readLine(r *reader) (string, error) {
	for {
		line, err := r.nextLine()
		if err != nil {
			return "", errors.Wrapf(err, "failed to read next line")
		}
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, comment) {
			continue
		}
		return line, nil
	}
}

func unquote(line, prefix string) (string, error) {
	line = strings.TrimPrefix(line, prefix)
	line = strings.TrimSpace(line)
	return strconv.Unquote(line)
}

func defaultPlural(msgID, msgIDPlural string, n int64) string {
	if n != 1 {
		return msgIDPlural
	}
	return msgID
}

// key 生成查找 message 的 key
func key(ctxt, id string) string {
	return ctxt + "\u0004" + id
}
