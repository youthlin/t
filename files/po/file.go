package po

// Po 文件结构说明
// https://www.gnu.org/software/gettext/manual/html_node/PO-Files.html

// white-space
// #  translator-comments
// #. extracted-comments
// #: reference…
// #, flag…
// #| msgctxt previous-context
// #| msgid previous-untranslated-string
// msgctxt context
// msgid untranslated-string
// msgid_plural untranslated-string-plural
// msgstr[0] translated-string-case-0
// ...
// msgstr[N] translated-string-case-n

import (
	"io"
	"strings"

	"github.com/pkg/errors"
	"github.com/youthlin/t/files"
)

// Parse 将 po 文件内容解析为结构体
func Parse(src string) (*files.File, error) {
	src = strings.ReplaceAll(src, "\r", "")
	lines := strings.Split(src, "\n")
	return parseLines(lines)
}

func parseLines(lines []string) (*files.File, error) {
	var result = files.NewEmptyFile()
	r := newReader(lines)
	for {
		msg, err := readMessage(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				if msg != nil {
					if err := result.AddMessage(msg); err != nil {
						return nil, err
					}
				}
				break
			}
			return nil, err
		}
		if err := result.AddMessage(msg); err != nil {
			return nil, err
		}
	}
	return result, nil
}
