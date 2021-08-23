package translator

import (
	"regexp"
	"strings"

	"github.com/youthlin/t/f"
)

const (
	HeaderPluralForms = "Plural-Forms" // 表明该语言的复数形式
	HeaderLanguage    = "Language"     // 表明该文件是什么语言
)

var _ Translator = (*File)(nil) // 触发编译检查，是否实现接口
var reHeader = regexp.MustCompile(`(.*?): (.*)`)

// File 一个翻译文件
type File struct {
	entries map[string]*Entry
	headers map[string]string
	plural  *plural
}

func (file *File) Lang() string {
	lang, _ := file.GetHeader(HeaderLanguage)
	return lang
}

func (file *File) T(msgID string, args ...interface{}) string {
	return file.X("", msgID, args...)
}

func (file *File) N(msgID, msgIDPlural string, n int, args ...interface{}) string {
	return file.XN64("", msgID, msgIDPlural, int64(n), args...)
}

func (file *File) N64(msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return file.XN64("", msgID, msgIDPlural, n, args...)
}

func (file *File) X(msgCtxt, msgID string, args ...interface{}) string {
	entry, ok := file.entries[key(msgCtxt, msgID)]
	if !ok || entry.msgStr == "" {
		return f.Format(msgID, args...)
	}
	return f.Format(entry.msgStr, args...)
}

func (file *File) XN(msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return file.XN64(msgCtxt, msgID, msgIDPlural, int64(n), args...)
}

func (file *File) XN64(msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	entry, ok := file.entries[key(msgCtxt, msgID)]
	if !ok {
		return f.DefaultPlural(msgID, msgIDPlural, n, args...)
	}
	plural := file.GetPlural()
	if plural.totalForms <= 0 || plural.fn == nil {
		return f.DefaultPlural(msgID, msgIDPlural, n, args...)
	}
	index := plural.fn(n)
	if index < 0 || index >= int(plural.totalForms) || index > len(entry.msgStrN) || entry.msgStrN[index] == "" {
		// 超出范围
		return f.DefaultPlural(msgID, msgIDPlural, n, args...)
	}
	return f.Format(entry.msgStrN[index], args...)
}

func (file *File) AddEntry(e *Entry) {
	if file.entries == nil {
		file.entries = map[string]*Entry{}
	}
	file.entries[e.key()] = e
	if e.isHeader() {
		file.initHeader()
		file.initPlural()
	}
}

func (file *File) GetHeader(key string) (value string, ok bool) {
	file.initHeader()
	value, ok = file.headers[key]
	return
}

func (file *File) initHeader() {
	if file.headers == nil {
		headers := make(map[string]string)
		if headerEntry, ok := file.entries[key("", "")]; ok {
			kvs := strings.Split(headerEntry.msgStr, "\n")
			for _, kv := range kvs {
				if kv == "" {
					continue
				}
				find := reHeader.FindAllStringSubmatch(kv, -1)
				if len(find) != 1 || len(find[0]) != 3 {
					continue
				}
				kv := find[0]
				k := strings.TrimSpace(kv[1])
				v := strings.TrimSpace(kv[2])
				headers[k] = v
			}
		}
		file.headers = headers
	}
}

func (file *File) GetPlural() *plural {
	file.initPlural()
	return file.plural
}

func (file *File) initPlural() {
	if file.plural == nil {
		forms, _ := file.GetHeader(HeaderPluralForms)
		file.plural = parsePlural(forms)
	}
}