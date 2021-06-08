package po

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type message struct {
	msgCTxt string
	msgID   string
	msgID2  string
	msgStr  string
	msgStrN []string
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
				msg.msgCTxt = txt
				continue
			}
			// read msgctxt content below msgctxt
			if strings.HasPrefix(line, quote) {
				txt, err := unquote(line, "")
				if err != nil {
					return nil, err
				}
				msg.msgCTxt += txt
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
			// 当前行不是这些开头 msgctx/msgid/msgid_plural/msgstr 需要报错
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

func (m *message) key() string {
	return key(m.msgCTxt, m.msgID)
}

func (m *message) isEmpty() bool {
	return m == nil || m.msgCTxt == "" && m.msgID == "" &&
		m.msgID2 == "" && m.msgStr == "" && len(m.msgStrN) == 0
}

func (m *message) isValid() bool {
	if m == nil {
		return false
	}
	if m.msgID == "" { // header
		return m.msgCTxt == "" && m.msgID2 == "" && m.msgStr != "" && len(m.msgStrN) == 0
	}
	return m.msgStr != "" || len(m.msgStrN) != 0
}
