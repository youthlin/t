package po

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

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

type File struct {
	headers  map[string]string
	messages map[string]*message
}

func newEmptyFile() *File {
	return &File{
		headers:  make(map[string]string),
		messages: make(map[string]*message),
	}
}

type message struct {
	context string
	msgID   string
	msgID2  string
	msgStr  string
	msgStrN []string
}

func (m *message) isEmpty() bool {
	return m == nil || m.context == "" && m.msgID == "" &&
		m.msgID2 == "" && m.msgStr == "" && len(m.msgStrN) == 0
}

const (
	comment = "#"
	msgCtxt = "msgctxt"
	msgID   = "msgid"
	msgID2  = "msgid_plural"
	msgStr  = "msgstr"
	msgStrN = "msgstr["
	quote   = `"`
)
const (
	HeaderPluralForms = "Plural-Forms"
)

// Parse 将 po 文件内容解析为结构体
func Parse(src string) (*File, error) {
	src = strings.ReplaceAll(src, "\r", "")
	lines := strings.Split(src, "\n")
	return parseLines(lines)
}

func parseLines(lines []string) (*File, error) {
	var result = newEmptyFile()
	r := newReader(lines)
	for {
		msg, err := readMessage(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				if msg != nil {
					result.messages[msg.msgID] = msg
				}
				break
			}
			return nil, err
		}
		result.messages[msg.msgID] = msg
	}
	return result, nil
}

// readMessage read message from reader
func readMessage(r *reader) (*message, error) {
	// 状态机的状态
	const (
		stateNew = iota
		stateCtx
		stateID
		stateID2
		stateStr
		stateStrN
		stateDone
	)
	var (
		state = stateNew   // 初始状态
		msg   = &message{} // 结果
		strN  = 0          // 复数索引
	)
	for {
		switch state {
		case stateNew:
			_, err := readLine(r)
			if err != nil {
				return nil, err
			}
			// 有内容，回退一行，流转状态至读取 msgctxt
			r.unGetLine()
			state = stateCtx
		case stateCtx:
			line, err := readLine(r)
			if err != nil {
				return msg, err
			}
			// read msgctxt line
			if strings.HasPrefix(line, msgCtxt) {
				txt, err := unquote(line, msgCtxt)
				if err != nil {
					return nil, err
				}
				msg.context = txt
				continue
			}
			// read msgctxt content below msgctxt
			if strings.HasPrefix(line, quote) {
				txt, err := unquote(line, "")
				if err != nil {
					return nil, err
				}
				msg.context += txt
				continue
			}
			// 不是 msgctxt 流转状态至 msgid
			r.unGetLine()
			state = stateID
		case stateID:
			line, err := readLine(r)
			if err != nil {
				return msg, err
			}
			// read msgid line
			if strings.HasPrefix(line, msgID) &&
				!strings.HasPrefix(line, msgID2) {
				txt, err := unquote(line, msgID)
				if err != nil {
					return nil, err
				}
				msg.msgID = txt
				continue
			}
			// read content below
			if strings.HasPrefix(line, quote) {
				txt, err := unquote(line, "")
				if err != nil {
					return nil, err
				}
				msg.msgID += txt
				continue
			}
			// 不是 msgid 流转状态至 msgid_plural
			r.unGetLine()
			state = stateID2
		case stateID2:
			line, err := readLine(r)
			if err != nil {
				return msg, err
			}
			// read msgid2 line
			if strings.HasPrefix(line, msgID2) {
				txt, err := unquote(line, msgID2)
				if err != nil {
					return nil, err
				}
				msg.msgID2 = txt
				continue
			}
			// read content below
			if strings.HasPrefix(line, quote) {
				txt, err := unquote(line, "")
				if err != nil {
					return nil, err
				}
				msg.msgID2 += txt
				continue
			}
			// 不是 msgid_plural 流转状态至 msgstr[x]
			// 因为 msgstr[x] 的前缀也是 msgstr, 而且不会共存，所以先处理复数
			r.unGetLine()
			state = stateStrN
		case stateStrN:
			line, err := readLine(r)
			if err != nil {
				return msg, err
			}
			// read msgid2 line
			var prefix = fmt.Sprintf("%s%d]", msgStrN, strN)
			if strings.HasPrefix(line, prefix) {
				txt, err := unquote(line, prefix)
				if err != nil {
					return nil, err
				}
				msg.msgStrN = append(msg.msgStrN, txt)
				strN++
				continue
			}
			// read content below
			if strings.HasPrefix(line, quote) {
				txt, err := unquote(line, "")
				if err != nil {
					return nil, err
				}
				msg.msgStrN[strN-1] += txt
				continue
			}
			// 不是复数，状态转移为 msgstr
			r.unGetLine()
			state = stateStr
		case stateStr:
			line, err := readLine(r)
			if err != nil {
				return msg, err
			}
			// read msgid2 line
			if strings.HasPrefix(line, msgStr) {
				txt, err := unquote(line, msgStr)
				if err != nil {
					return nil, err
				}
				msg.msgStr = txt
				continue
			}
			// read content below
			if strings.HasPrefix(line, quote) {
				txt, err := unquote(line, "")
				if err != nil {
					return nil, err
				}
				msg.msgStr += txt
				continue
			}
			r.unGetLine()
			state = stateDone
		case stateDone:
			line, _ := readLine(r)
			r.unGetLine()
			if msg.isEmpty() {
				return nil, errors.Errorf("not valid content: %v", line)
			} else {
				switch {
				case strings.HasPrefix(line, msgCtxt):
					break
				case strings.HasPrefix(line, msgID):
					break
				case strings.HasPrefix(line, msgStr):
					break
				default:
					return nil, errors.Errorf("unexpected text: %v", line)
				}
			}
			return msg, nil
		}
	}
}

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
