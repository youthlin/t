package translator

import (
	"regexp"
	"sort"
	"strings"

	"github.com/youthlin/t/f"
)

const (
	// HeaderPluralForms 表明该语言的复数形式
	HeaderPluralForms = "Plural-Forms"
	// HeaderLanguage 表明该文件是什么语言
	HeaderLanguage = "Language"
)

var _ Translator = (*File)(nil) // 触发编译检查，是否实现接口
var reHeader = regexp.MustCompile(`(.*?): (.*)`)

// File 一个翻译文件
type File struct {
	entries map[string]*Entry
	headers map[string]string
	plural  *plural
}

// Lang get this translations' language
func (file *File) Lang() string {
	lang, _ := file.GetHeader(HeaderLanguage)
	return lang
}

// X is ashort name for pgettext
func (file *File) X(msgCtxt, msgID string, args ...interface{}) string {
	entry, ok := file.entries[key(msgCtxt, msgID)]
	if !ok || entry.MsgStr == "" {
		return f.Format(msgID, args...)
	}
	return f.Format(entry.MsgStr, args...)
}

// XN64 is ashort name for npgettext
func (file *File) XN64(msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	entry, ok := file.entries[key(msgCtxt, msgID)]
	if !ok {
		return f.DefaultPlural(msgID, msgIDPlural, n, args...)
	}
	plural := file.getPlural()
	if plural.totalForms <= 0 || plural.fn == nil {
		return f.DefaultPlural(msgID, msgIDPlural, n, args...)
	}
	index := plural.fn(n)
	if index < 0 || index >= int(plural.totalForms) || index > len(entry.MsgStrN) || entry.MsgStrN[index] == "" {
		// 超出范围
		return f.DefaultPlural(msgID, msgIDPlural, n, args...)
	}
	return f.Format(entry.MsgStrN[index], args...)
}

// SortedEntry sort entry by key
func (file *File) SortedEntry() (entries []*Entry) {
	for _, e := range file.entries {
		entries = append(entries, e)
	}
	sort.Slice(entries, func(i, j int) bool {
		left := entries[i]
		right := entries[j]
		return left.getSortKey() < right.getSortKey()
	})
	return
}

// AddEntry adds a Entry
func (file *File) AddEntry(e *Entry) {
	if file.entries == nil {
		file.entries = map[string]*Entry{}
	}
	file.entries[e.Key()] = e
	if e.isHeader() {
		file.initHeader()
		file.initPlural()
	}
}

// GetHeader get header value by key
func (file *File) GetHeader(key string) (value string, ok bool) {
	file.initHeader()
	value, ok = file.headers[key]
	return
}

func (file *File) initHeader() {
	if file.headers == nil {
		headers := make(map[string]string)
		if headerEntry, ok := file.entries[key("", "")]; ok {
			kvs := strings.Split(headerEntry.MsgStr, "\n")
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
				if _, ok := headers[k]; !ok {
					headers[k] = v
				}
			}
		}
		file.headers = headers
	}
}

func (file *File) getPlural() *plural {
	file.initPlural()
	return file.plural
}

func (file *File) initPlural() {
	if file.plural == nil {
		forms, _ := file.GetHeader(HeaderPluralForms)
		file.plural = parsePlural(forms)
	}
}
