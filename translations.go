package t

import (
	"io/fs"
	"sort"

	"github.com/youthlin/t/locale"
)

// DefaultDomain 默认的文本域
const DefaultDomain = "default"

// DefaultSourceCodeLocale 默认的源代码语言
const DefaultSourceCodeLocale = "en_US"

// Translations holds several translation domains
// ts. [翻译集]包含多个翻译，每个翻译分别属于一个文本域
type Translations struct {
	locale  string
	domain  string
	domains map[string]*Translation // key is domain
	// sourceCodeLocale 源代码中的语言, 通常应该使用英文
	sourceCodeLocale string
}

// NewTranslations create a new Translations 新建翻译集
func NewTranslations() *Translations {
	return &Translations{
		locale:           locale.GetDefault(),
		domain:           DefaultDomain,
		domains:          make(map[string]*Translation),
		sourceCodeLocale: DefaultSourceCodeLocale,
	}
}

// clone clones a Translations
func (ts *Translations) clone() *Translations {
	return &Translations{
		locale: ts.locale,
		domain: ts.domain,
		domains: func() map[string]*Translation {
			m := make(map[string]*Translation)
			for d, tr := range ts.domains {
				m[d] = tr
			}
			return m
		}(),
		sourceCodeLocale: ts.sourceCodeLocale,
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

// HasDomain return true if ts has the specified domain
func (ts *Translations) HasDomain(domain string) bool {
	for d := range ts.domains {
		if d == domain {
			return true
		}
	}
	return false
}

// Domains return all domains
func (ts *Translations) Domains() (domains []string) {
	for domain := range ts.domains {
		domains = append(domains, domain)
	}
	return
}

// Locale return current locale 返回当前使用的语言
func (ts *Translations) Locale() string {
	return ts.locale
}

// Locales return all supported locales of domain 返回文本域中支持的所有语言
func (ts *Translations) Locales() (locales []string) {
	tr := ts.GetOrNoop(ts.domain)
	m := make(map[string]struct{}, len(tr.langs)+1)
	m[ts.sourceCodeLocale] = struct{}{}
	locales = append(locales, ts.sourceCodeLocale)
	for lang := range tr.langs {
		lang = locale.Normalize(lang)
		if _, ok := m[lang]; !ok {
			m[lang] = struct{}{}
			locales = append(locales, lang)
		}
	}
	sort.Strings(locales)
	return
}

// UsedLocale return the locale that actually used
func (ts *Translations) UsedLocale() string {
	tr, ok := ts.Get(ts.domain)
	if !ok {
		return ts.sourceCodeLocale
	}
	_, ok = tr.Get(ts.locale)
	if !ok {
		return ts.sourceCodeLocale
	}
	return ts.locale
}

// SetLocale set current locale 设置要使用的语言
func (ts *Translations) SetLocale(lang string) {
	if lang == "" {
		lang = locale.GetDefault()
	} else {
		lang = locale.Normalize(lang)
	}
	ts.locale = lang
}

// SourceCodeLocale 设置源代码语言
func (ts *Translations) SourceCodeLocale() string { return ts.sourceCodeLocale }

// SetSourceCodeLocale 设置源代码语言
func (ts *Translations) SetSourceCodeLocale(lang string) {
	lang = locale.Normalize(lang)
	ts.sourceCodeLocale = lang
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

// D return a new Translations with domain
func (ts *Translations) D(domain string) *Translations {
	result := ts.clone()
	result.SetDomain(domain)
	return result
}

// L return a new Translations with locale
func (ts *Translations) L(locale string) *Translations {
	result := ts.clone()
	result.SetLocale(locale)
	return result
}

// T is a short name of gettext
func (ts *Translations) T(msgID string, args ...interface{}) string {
	return ts.X("", msgID, args...)
}

// N is a short name of nettext
func (ts *Translations) N(msgID, msgIDPlural string, n int, args ...interface{}) string {
	return ts.XN64("", msgID, msgIDPlural, int64(n), args...)
}

// N64 is a short name of nettext
func (ts *Translations) N64(msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return ts.XN64("", msgID, msgIDPlural, n, args...)
}

// X is a short name of pgettext
func (ts *Translations) X(msgCtxt, msgID string, args ...interface{}) string {
	tr := ts.GetOrNoop(ts.domain)
	tor := tr.GetOrNoop(ts.locale)
	return tor.X(msgCtxt, msgID, args...)
}

// XN is a short name of npgettext
func (ts *Translations) XN(msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return ts.XN64(msgCtxt, msgID, msgIDPlural, int64(n), args...)
}

// XN64 is a short name of npgettext
func (ts *Translations) XN64(msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	tr := ts.GetOrNoop(ts.domain)
	tor := tr.GetOrNoop(ts.locale)
	return tor.XN64(msgCtxt, msgID, msgIDPlural, n, args...)
}
