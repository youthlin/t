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
	"github.com/youthlin/t/f"
	"github.com/youthlin/t/plurals"
)

// Parse 将 po 文件内容解析为结构体
func Parse(src string) (*File, error) {
	src = strings.ReplaceAll(src, "\r", "")
	lines := strings.Split(src, "\n")
	return parseLines(lines)
}

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
	lang       string
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

// Lang 返回译文的语言
func (po *File) Lang() string {
	if po.lang == "" {
		po.lang, _ = po.GetHeader(HeaderLanguage)
	}
	return po.lang
}
func (po *File) SetLang(lang string) {
	po.lang = lang
}

// T gettext 直接获取翻译内容，如果没有翻译，返回原始内容
// 如果 args 不为空，则将翻译后的字符串作为格式化模版，格式化 args
func (po *File) T(msgID string, args ...interface{}) string {
	return po.X("", msgID, args...)
}

// N ngettext 翻译复数，如果没有翻译，n 大于 1 返回原文复数(msgIDPlural)，否则返回原文单数(msgID)
// 如果 args 不为空，则将翻译后的字符串作为格式化模版，格式化 args
// 注意，n 用于选择第几种复数，如果 需要打印 n，还需要将其传包括在 args 中.
//
// Note: n is used to choose plural forms, is you need print n, you should pass it to args
//  // no args, so return: `%d apples`
//  po.N("one apple", "%d apples", 2) -> "%d apples"
//  // the first numer 2 result in `%d apples`, the second 2 format to `2 apples`
//  po.N("one apple", "%d apples", 2, 2) -> "2 apples"
//  po.N("one apple", "%d apples", 2, 200) -> "200 apples"
//
//  // use `one apple` as template to format number `200`, the extra arg ignored, see f.Format
//  po.N("one apple", "%d apples", 1, 200) -> "one apple"
func (po *File) N(msgID, msgIDPlural string, n int, args ...interface{}) string {
	return po.XN64("", msgID, msgIDPlural, int64(n), args...)
}

// N64 ngettext 翻译复数，如果没有翻译，n 大于 1 返回原文复数(msgIDPlural)，否则返回原文单数(msgID)
func (po *File) N64(msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return po.XN64("", msgID, msgIDPlural, n, args...)
}

// X pgettext 带上下文翻译，用于区分同一个 msgID 在不同上下文的不同含义
func (po *File) X(msgCtxt, msgID string, args ...interface{}) string {
	msg, ok := po.messages[key(msgCtxt, msgID)]
	if !ok {
		return f.Format(msgID, args...)
	}
	return f.Format(msg.msgStr, args...)
}

// XN pngettext 带上下文翻译复数
func (po *File) XN(msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return po.XN64(msgCtxt, msgID, msgIDPlural, int64(n), args...)
}

// XN64 pngettext 带上下文翻译复数
func (po *File) XN64(msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	msg, ok := po.messages[key(msgCtxt, msgID)]
	if !ok {
		return f.DefaultPlural(msgID, msgIDPlural, n, args...)
	}
	totalForms := po.GetTotalForms()
	pluralFunc := po.GetPluralFunc()
	if totalForms <= 0 || pluralFunc == nil { // 复数设置不正确
		return f.DefaultPlural(msgID, msgIDPlural, n, args...)
	}
	// 看 n 对应第几种复数
	index := pluralFunc(n)
	if index < 0 || index >= int(totalForms) || index > len(msg.msgStrN) {
		// 超出范围
		return f.DefaultPlural(msgID, msgIDPlural, n, args...)
	}
	return f.Format(msg.msgStrN[index], args...)
}

// SetPluralForms 设置复数形式一共有 totalForms 种；函数 nForm 设置每个 n 对应第几种复数
func (po *File) SetPluralForms(totalForms int, pluralFunc PluralFunc) {
	po.totalForms = totalForms
	po.pluralFunc = pluralFunc
}

func (po *File) GetTotalForms() int {
	if po.totalForms >= 0 {
		return po.totalForms
	}
	find, ok := po.getPluralArr()
	if !ok {
		po.totalForms = 0
		return 0
	}
	num, err := strconv.ParseInt(find[0], 10, 64)
	if err != nil {
		num = 0
	}
	po.totalForms = int(num)
	return int(num)
}

func (po *File) GetPluralFunc() PluralFunc {
	if po.pluralFunc != nil {
		return po.pluralFunc
	}
	find, ok := po.getPluralArr()
	if !ok {
		po.pluralFunc = invalidPluralFunc
		return invalidPluralFunc
	}
	exp := find[1]
	po.pluralFunc = func(n int64) int {
		index, err := plurals.Eval(context.Background(), exp, n)
		if err != nil {
			return -1
		}
		return int(index)
	}
	return po.pluralFunc
}

func (po *File) GetHeader(key string) (string, bool) {
	v, ok := po.headers[key]
	return v, ok
}

func (po *File) getPluralArr() ([]string, bool) {
	forms, ok := po.GetHeader(HeaderPluralForms)
	if !ok {
		return nil, false
	}
	find := rePlurals.FindAllStringSubmatch(forms, -1)
	if len(find) != 1 || len(find[0]) != 3 {
		return nil, false
	}
	return find[0][1:], true
}

func (po *File) addMessage(m *message) error {
	if m.isValid() {
		if m.msgID == "" {
			headerLines := strings.Split(m.msgStr, "\n")
			for _, headerLine := range headerLines {
				if headerLine == "" {
					continue
				}
				find := reHeader.FindAllStringSubmatch(headerLine, -1)
				if len(find) != 1 || len(find[0]) != 3 {
					return errors.Errorf("invalid header|line=%v|entry=%+v", headerLine, m)
				}
				kv := find[0]
				k := strings.TrimSpace(kv[1])
				v := strings.TrimSpace(kv[2])
				v = strings.TrimSuffix(v, `\n`)
				po.headers[k] = v
			}
			return nil
		}
		po.messages[m.key()] = m
		return nil
	}
	return errors.Errorf("invalid message|%+v", m)
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
