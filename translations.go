package t

import (
	"io/fs"

	"github.com/youthlin/t/locale"
)

const DefaultDomain = "default"

// Translations holds several translation domains
// 翻译集包含多个翻译，每个翻译分别属于一个文本域
type Translations struct {
	locale  string
	domain  string
	domains map[string]*Translation // key is domain
}

// NewTranslations create a new Translations 新建翻译集
func NewTranslations() *Translations {
	return &Translations{
		locale:  locale.GetDefault(),
		domain:  DefaultDomain,
		domains: make(map[string]*Translation),
	}
}

// BindFS load a Translation form file system and bind to a domain
// 从文件系统绑定翻译域
func (ts *Translations) BindFS(domain string, fsys fs.FS) {
	tr := NewTranslation(domain)
	if tr.LoadFS(fsys) {
		ts.domains[domain] = tr
	}
}

// Domain return current domain 返回当前使用的文本域
func (ts *Translations) Domain() string {
	return ts.domain
}

// SetDomain set current domain 设置要使用的文本域
func (ts *Translations) SetDomain(domain string) {
	ts.domain = domain
}

// Locale return current locale 返回当前使用的语言
func (ts *Translations) Locale() string {
	return ts.locale
}

// SetLocale set current locale 设置要使用的语言
func (ts *Translations) SetLocale(lang string) {
	ts.locale = lang
}

// Get return the Translation of the specified domain
// 获取指定的的翻译域
func (ts *Translations) Get(domain string) (*Translation, bool) {
	tr, ok := ts.domains[domain]
	return tr, ok
}

// GetOrNoop return the Translation of the specified domain
// 获取指定的的翻译域
func (ts *Translations) GetOrNoop(domain string) *Translation {
	if tr, ok := ts.domains[domain]; ok {
		return tr
	}
	return trNoop
}

// T is a short name of gettext, which will translate and format the msgID
func (ts *Translations) T(msgID string, args ...interface{}) string {
	tr := ts.GetOrNoop(ts.domain)
	tor := tr.GetOrNoop(ts.locale)
	return tor.T(msgID, args...)
}
