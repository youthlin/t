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
	"context"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/youthlin/t/plurals"
	"golang.org/x/text/language"
)

// PluralFunc 定义一个整数 n 对应翻译语言中的第几种复数
type PluralFunc func(int64) int

// invalidPluralFunc 返回 -1 会使用原文的单复数
var invalidPluralFunc PluralFunc = func(i int64) int { return -1 }

var (
	// parse header key-value
	reHeader = regexp.MustCompile(`(.*?): (.*)`)
	// parse Plural-Forms header
	rePlurals = regexp.MustCompile(`^\s*nplurals\s*=\s*(\d)\s*;\s*plural\s*=\s*(.*)\s*;$`)
)

const (
	HeaderPluralForms = "Plural-Forms" // 表明该语言的复数形式
	HeaderLanguage    = "Language"     // 表明该文件是什么语言
)

// File po file struct
type File struct {
	headers    map[string]string
	messages   map[string]*message
	totalForms int
	pluralFunc PluralFunc
}

// newEmptyFile 构造函数
func newEmptyFile() *File {
	return &File{
		headers:    make(map[string]string),
		messages:   make(map[string]*message),
		totalForms: -1,
	}
}

// T gettext 直接获取翻译内容，如果没有翻译，返回原始内容
func (f *File) T(msgID string) string {
	return f.X("", msgID)
}

// N ngettext 翻译复数，如果没有翻译，n 大于 1 返回原文复数(msgIDPlural)，否则返回原文单数(msgID)
func (f *File) N(msgID, msgIDPlural string, n int) string {
	return f.XN64("", msgID, msgIDPlural, int64(n))
}

// N64 ngettext 翻译复数，如果没有翻译，n 大于 1 返回原文复数(msgIDPlural)，否则返回原文单数(msgID)
func (f *File) N64(msgID, msgIDPlural string, n int64) string {
	return f.XN64("", msgID, msgIDPlural, n)
}

// X pgettext 带上下文翻译，用于区分同一个 msgID 在不同上下文的不同含义
func (f *File) X(msgCtxt, msgID string) string {
	msg, ok := f.messages[key(msgCtxt, msgID)]
	if !ok {
		return msgID
	}
	return msg.msgStr
}

// XN pngettext 带上下文翻译复数
func (f *File) XN(msgCtxt, msgID, msgIDPlural string, n int) string {
	return f.XN64(msgCtxt, msgID, msgIDPlural, int64(n))
}

// XN64 pngettext 带上下文翻译复数
func (f *File) XN64(msgCtxt, msgID, msgIDPlural string, n int64) string {
	msg, ok := f.messages[key(msgCtxt, msgID)]
	if !ok {
		return defaultPlural(msgID, msgIDPlural, n)
	}
	totalForms := f.GetTotalForms()
	pluralFunc := f.GetPluralFunc()
	if totalForms <= 0 || pluralFunc == nil { // 复数设置不正确
		return defaultPlural(msgID, msgIDPlural, n)
	}
	// 看 n 对应第几种复数
	index := pluralFunc(n)
	if index < 0 || index >= int(totalForms) || index > len(msg.msgStrN) {
		// 超出范围
		return defaultPlural(msgID, msgIDPlural, n)
	}
	return msg.msgStrN[index]
}

// SetPluralForms 设置复数形式一共有 totalForms 种；函数 nForm 设置每个 n 对应第几种复数
func (f *File) SetPluralForms(totalForms int, pluralFunc PluralFunc) {
	f.totalForms = totalForms
	f.pluralFunc = pluralFunc
}

func (f *File) GetTotalForms() int {
	if f.totalForms >= 0 {
		return f.totalForms
	}
	find, ok := f.getPluralArr()
	if !ok {
		f.totalForms = 0
		return 0
	}
	num, err := strconv.ParseInt(find[0], 10, 64)
	if err != nil {
		num = 0
	}
	f.totalForms = int(num)
	return int(num)
}

func (f *File) GetPluralFunc() PluralFunc {
	if f.pluralFunc != nil {
		return f.pluralFunc
	}
	find, ok := f.getPluralArr()
	if !ok {
		f.pluralFunc = invalidPluralFunc
		return invalidPluralFunc
	}
	exp := find[1]
	f.pluralFunc = func(n int64) int {
		index, err := plurals.Eval(context.Background(), exp, n)
		if err != nil {
			return -1
		}
		return int(index)
	}
	return f.pluralFunc
}

func (f *File) GetHeader(key string) (string, bool) {
	v, ok := f.headers[key]
	return v, ok
}

func (f *File) GetLanguage() (language.Tag, bool) {
	lang, ok := f.GetHeader(HeaderLanguage)
	if !ok {
		return language.Und, false
	}
	tag, err := language.Default.Parse(lang)
	if err != nil {
		return language.Und, false
	}
	return tag, true
}

func (f *File) getPluralArr() ([]string, bool) {
	forms, ok := f.GetHeader(HeaderPluralForms)
	if !ok {
		return nil, false
	}
	find := rePlurals.FindAllStringSubmatch(forms, -1)
	if len(find) != 1 || len(find[0]) != 3 {
		return nil, false
	}
	return find[0][1:], true
}

func (f *File) addMessage(m *message) error {
	if m.isValid() {
		if m.msgID == "" {
			headerLines := strings.Split(m.msgStr, "\n")
			for _, headerLine := range headerLines {
				if headerLine == "" {
					continue
				}
				find := reHeader.FindAllStringSubmatch(headerLine, -1)
				if len(find) != 1 || len(find[0]) != 3 {
					return errors.Errorf("invalid header|line=%v|header=%+v", headerLine, m)
				}
				kv := find[0]
				k := strings.TrimSpace(kv[1])
				v := strings.TrimSpace(kv[2])
				v = strings.TrimSuffix(v, `\n`)
				f.headers[k] = v
			}
			return nil
		}
		f.messages[m.key()] = m
		return nil
	}
	return errors.Errorf("invalid message|%+v", m)
}

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
					if err := result.addMessage(msg); err != nil {
						return nil, err
					}
				}
				break
			}
			return nil, err
		}
		if err := result.addMessage(msg); err != nil {
			return nil, err
		}
	}
	return result, nil
}
