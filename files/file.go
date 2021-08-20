package files

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/youthlin/t/locale"
	"github.com/youthlin/t/plurals"
)

const (
	split             = "\u0004"       // 分割 msgCtxt, msgID
	HeaderPluralForms = "Plural-Forms" // 表明该语言的复数形式
	HeaderLanguage    = "Language"     // 表明该文件是什么语言
)

var (
	// parse header key-value 解析键值对
	reHeader = regexp.MustCompile(`(.*?): (.*)`)
	// parse Plural-Forms header 解析复数字段
	rePlurals = regexp.MustCompile(`^\s*nplurals\s*=\s*(\d)\s*;\s*plural\s*=\s*(.*)\s*;$`)
	// invalidPluralFunc 返回 -1 会使用原文的单复数
	invalidPluralFunc PluralFunc = func(i int64) int { return -1 }
)

// PluralFunc 定义一个整数 n 对应翻译语言中的第几种复数
type PluralFunc func(int64) int

// File 表示一个 po/mo 文件
type File struct {
	headers    map[string]string
	messages   map[string]*Message
	lang       string
	totalForms int
	pluralFunc PluralFunc
}

// NewEmptyFile 构造函数
func NewEmptyFile() *File {
	return &File{
		headers:    make(map[string]string),
		messages:   make(map[string]*Message),
		totalForms: -1,
	}
}

// SetLang 设置语言
func (f *File) SetLang(lang string) *File {
	f.lang = locale.Normalize(lang)
	return f
}
func (f *File) SetHeaders(h map[string]string) *File {
	f.headers = h
	return f
}
func (f *File) SetMessages(m map[string]*Message) *File {
	f.messages = m
	return f
}

// GetHeader 获取 header 中指定的字段
func (f *File) GetHeader(key string) (string, bool) {
	v, ok := f.headers[key]
	return v, ok
}

// SetPluralForms 设置复数形式一共有 totalForms 种；函数 nForm 设置每个 n 对应第几种复数
func (f *File) SetPluralForms(totalForms int, pluralFunc PluralFunc) *File {
	f.totalForms = totalForms
	f.pluralFunc = pluralFunc
	return f
}

// GetTotalForms 返回有几种复数形式
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

// GetPluralFunc 返回计算复数的函数
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

// getPluralArr 返回复数总数、复数表达式
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

// AddMessage add Message to File.
func (f *File) AddMessage(m *Message) error {
	if m.IsValid() {
		if m.MsgID == "" {
			headerLines := strings.Split(m.MsgStr, "\n")
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
				f.headers[k] = v
			}
			return nil
		}
		f.messages[m.Key()] = m
		return nil
	}
	return errors.Errorf("invalid message|%+v", m)
}

// key 生成查找 message 的 key
func key(ctxt, id string) string {
	return ctxt + split + id
}
