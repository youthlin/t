package po

import (
	"strconv"
	"strings"

	"github.com/cockroachdb/errors"
)

const (
	comment = "#"
	msgCtxt = "msgctxt"
	msgID   = "msgid"
	msgID2  = "msgid_plural"
	msgStr  = "msgstr"
	msgStrN = "msgstr["
	quote   = "\"" // 写成 `"` xgettest 会认为是没结束的引号
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
