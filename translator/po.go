package translator

// The Format of PO Files
// https://www.gnu.org/software/gettext/manual/html_node/PO-Files.html

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/cockroachdb/errors"
)

var errEmptyPo = fmt.Errorf("empty po file")

func ReadPo(content []byte) (*File, error) {
	if len(content) == 0 {
		return nil, errors.Wrapf(errEmptyPo, "read po file failed")
	}
	data := string(content)
	data = strings.ReplaceAll(data, "\r", "")
	r := newReader(strings.Split(data, "\n"))
	file := new(File)
	for {
		entry, err := readEntry(r)
		if entry.isValid() {
			file.AddEntry(entry)
		}
		if errors.Is(err, io.EOF) {
			return file, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func (f *File) SaveAsPo(w io.Writer) error {
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
		if entry.MsgID2 != "" {
			writeString(buf, msgid_plural, entry.MsgID2)
		}
		if entry.MsgStr != "" {
			writeString(buf, msgstr, entry.MsgStr)
		}
		for i, str := range entry.MsgStrN {
			writeString(buf, fmt.Sprintf(msgstrN, i), str)
		}
		buf.WriteString("\n")
	}
	_, err := buf.WriteTo(w)
	return err
}

var errInvalidEntry = fmt.Errorf("invalid entry")

// readEntry 读取翻译条目
func readEntry(r *reader) (*Entry, error) {
	const ( // 记录上一行读取的字段
		stateCtxt = "ctxt"
		stateID2  = "msgid_plural"
		stateID   = "msgid"
		stateStrN = "msgstr[]"
		stateStr  = "msgstr"
	)
	const ( // 各种情形的前缀
		prefixCmt   = "#"
		prefixCtxt  = "msgctxt "
		prefixID2   = "msgid_plural "
		prefixID    = "msgid "
		prefixStrN  = "msgstr[%d] "
		prefixStr   = "msgstr "
		prefixQuote = "\""
	)
	var (
		entry = new(Entry)
		// previousLineState 上一行的状态
		previousLineState = ""
	)
	for {
		line, err := r.nextLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return entry, err
			}
			return nil, errors.WithSecondaryError(errInvalidEntry,
				errors.Wrapf(err, "read entry failed"))
		}
		// 空白行，丢弃
		if strings.TrimSpace(line) == "" {
			continue
		}
		// 注释
		if strings.HasPrefix(line, prefixCmt) {
			if previousLineState != "" {
				// current entry has been parsed
				// this line is next entry's comment
				r.unGetLine()
				return entry, nil
			}
			entry.MsgCmts = append(entry.MsgCmts, line)
			continue
		}
		// ctxt
		if strings.HasPrefix(line, prefixCtxt) {
			if previousLineState != "" {
				r.unGetLine()
				return entry, nil
			}
			previousLineState = stateCtxt
			data, err := removePrefixAndUnquote(line, prefixCtxt)
			if err != nil {
				return nil, errors.WithSecondaryError(errInvalidEntry,
					errors.Wrapf(err, "unquote msgctxt failed|line %d: %s", r.lineNo, line))
			}
			entry.MsgCtxt += data
			continue
		}
		// msgid_plural
		if strings.HasPrefix(line, prefixID2) {
			if equalsAny(previousLineState, stateID2, stateStr, stateStrN) {
				r.unGetLine() // 当前条目已经有这些字段了，说明当前行是下一个条目的
				return entry, nil
			}
			previousLineState = stateID2
			data, err := removePrefixAndUnquote(line, prefixID2)
			if err != nil {
				return nil, errors.WithSecondaryError(errInvalidEntry,
					errors.Wrapf(err, "unquote msgid_plural failed|line %d: %s", r.lineNo, line))
			}
			entry.MsgID2 += data
			continue
		}
		// msgid
		if strings.HasPrefix(line, prefixID) {
			if equalsAny(previousLineState, stateID, stateID2, stateStr, stateStrN) {
				r.unGetLine()
				return entry, nil
			}
			previousLineState = stateID
			data, err := removePrefixAndUnquote(line, prefixID)
			if err != nil {
				return nil, errors.WithSecondaryError(errInvalidEntry,
					errors.Wrapf(err, "unquote msgid failed|line %d: %s", r.lineNo, line))
			}
			entry.MsgID += data
			continue
		}
		// msgstr[0]
		if prefix := fmt.Sprintf(prefixStrN, len(entry.MsgStrN)); strings.HasPrefix(line, prefix) {
			if equalsAny(previousLineState, stateStr) {
				r.unGetLine()
				return entry, nil
			}
			previousLineState = stateStrN
			data, err := removePrefixAndUnquote(line, prefix)
			if err != nil {
				return nil, errors.WithSecondaryError(errInvalidEntry,
					errors.Wrapf(err, "unquote %s failed|line %d: %s", prefix, r.lineNo, line))
			}
			entry.MsgStrN = append(entry.MsgStrN, data)
			continue
		}
		// msgstr
		if strings.HasPrefix(line, prefixStr) {
			if equalsAny(previousLineState, stateStrN) {
				r.unGetLine()
				return entry, nil
			}
			previousLineState = stateStr
			data, err := removePrefixAndUnquote(line, prefixStr)
			if err != nil {
				return nil, errors.WithSecondaryError(errInvalidEntry,
					errors.Wrapf(err, "unquote msgstr failed|line %d: %s", r.lineNo, line))
			}
			entry.MsgStr += data
			continue
		}

		// msgid "previous line"
		// "current line"
		if strings.HasPrefix(line, prefixQuote) {
			data, err := removePrefixAndUnquote(line, "")
			if err != nil {
				return nil, errors.WithSecondaryError(errInvalidEntry,
					errors.Wrapf(err, "unquote %s failed|line %d: %s", previousLineState, r.lineNo, line))
			}
			switch previousLineState {
			case stateCtxt:
				entry.MsgCtxt += data
			case stateID2:
				entry.MsgID2 += data
			case stateID:
				entry.MsgID += data
			case stateStrN:
				entry.MsgStrN[len(entry.MsgStrN)-1] += data
			case stateStr:
				entry.MsgStr += data
			}
		} else {
			return nil, errors.WithSecondaryError(errInvalidEntry,
				errors.Errorf("unexpected line %d: %s", r.lineNo, line))
		}
	}
}

func equalsAny(data string, args ...string) bool {
	for _, arg := range args {
		if data == arg {
			return true
		}
	}
	return false
}
