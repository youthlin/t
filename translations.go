package t

import (
	"io/fs"

	"github.com/youthlin/t/locale"
)

// Translations holds a map of domain to Translation
// 翻译集结构体包含多个翻译. 这是一个有状态的结构体，因此是协程不安全的。
// locale 字段记录用户希望使用的语言(实际输出的语言不一定就是这个)
// currentDomain 字段记录当前的翻译文本域，通常使用默认的域即可。
type Translations struct {
	locale        string
	currentDomain string
	trMap         map[string]*Translation // key is domain
}

// NewTranslations make a new instace
// 新建默认的翻译集实例
func NewTranslations() *Translations {
	return &Translations{
		locale:        locale.GetDefault(),
		currentDomain: DefaultDomain,
		trMap:         make(map[string]*Translation)}
}

// SetLocale set locale to ll_CC form
// 设置用户偏好的语言
func (ts *Translations) SetLocale(lang string) {
	if lang == "" {
		lang = locale.GetDefault()
	}
	ts.locale = locale.Normalize(lang)
}

// Locale get locale
// 返回用户设置的偏好语言
func (ts *Translations) Locale() string {
	return ts.locale
}

// UseLocale return a new instance of locale
// 用指定的用户偏好语言新建一个新的翻译集实例
func (ts *Translations) UseLocale(lang string) *Translations {
	result := ts.copy()
	result.locale = locale.Normalize(lang)
	return result
}

// DomainUsedLocale get display language
// 返回实际使用的语言，如果没有找到可用的语言，返回空
func (ts *Translations) DomainUsedLocale(domain string) string {
	tr, ok := ts.trMap[domain]
	if !ok {
		return ""
	}
	if _, ok := tr.translators[ts.locale]; ok {
		return ts.locale
	}
	return ""
}

// CurrentDomain get current domain
// 返回当前的文本域
func (ts *Translations) CurrentDomain() string {
	return ts.currentDomain
}

// UseDomain return a new instance of domain
// 用指定的文本域新建一个翻译集
func (ts *Translations) UseDomain(domain string) *Translations {
	result := ts.copy()
	result.TextDomain(domain)
	return result
}

func (ts *Translations) copy() *Translations {
	result := &Translations{
		locale:        ts.locale,
		currentDomain: ts.currentDomain,
		trMap:         make(map[string]*Translation),
	}
	for domain, tr := range ts.trMap {
		result.trMap[domain] = tr
	}
	return result
}

// Bind search .po/.mo file in path, and bind the result with domain
// 绑定翻译文件
func (ts *Translations) Bind(domain, path string) {
	tr := NewTranslation(domain)
	tr.Load(path)
	if len(tr.translators) > 0 {
		ts.trMap[domain] = tr
	}
}

func (ts *Translations) Load(domain, path string) {
	tr, ok := ts.Get(domain)
	if !ok {
		tr = NewTranslation(domain)
	}
	tr.Load(path)
	if len(tr.translators) > 0 {
		ts.trMap[domain] = tr
	}
}

func (ts *Translations) LoadFS(domain string, fsys fs.FS) {
	tr, ok := ts.Get(domain)
	if !ok {
		tr = NewTranslation(domain)
	}
	tr.LoadFS(fsys)
	if len(tr.translators) > 0 {
		ts.trMap[domain] = tr
	}
}

// TextDomain if domain exists, set it as current domain, else set current domain to DefaultDomain
// 绑定文本域，如果翻译集中的翻译没有指定的文本域，则使用默认的。该方法会返回绑定的文本域。
func (ts *Translations) TextDomain(domain string) string {
	if _, ok := ts.trMap[domain]; ok {
		return ts.setDomain(domain)
	}
	return ts.setDomain(DefaultDomain)
}
func (ts *Translations) setDomain(domain string) string {
	ts.currentDomain = domain
	return domain
}

// SupportLangs return all supported languages
func (ts *Translations) SupportLangs(domain string) (langs []string) {
	if tr, ok := ts.trMap[domain]; ok {
		langs = append(langs, tr.langs...)
	}
	return
}

// Add add a domain
// 向翻译集中添加翻译
func (ts *Translations) Add(tr *Translation) {
	if ts.trMap == nil {
		ts.trMap = make(map[string]*Translation)
	}
	ts.trMap[tr.domain] = tr
}

// Get get Translation of domain
// 获取指定文本域的翻译
func (ts *Translations) Get(domain string) (*Translation, bool) {
	tr, ok := ts.trMap[domain]
	return tr, ok
}

// GetOrNoop if no such domain, return a noop Translation
// 获取指定文本域的翻译，如果不存在返回一个 Noop 翻译
func (ts *Translations) GetOrNoop(domain string) *Translation {
	tr, ok := ts.Get(domain)
	if !ok {
		return trNoop
	}
	return tr
}

// #region gettext

func (ts *Translations) T(msgID string, args ...interface{}) string {
	return ts.DLT(ts.currentDomain, ts.locale, msgID, args...)
}
func (ts *Translations) N(msgID, msgIDPlural string, n int, args ...interface{}) string {
	return ts.DLN64(ts.currentDomain, ts.locale, msgID, msgIDPlural, int64(n), args...)
}
func (ts *Translations) N64(msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return ts.DLN64(ts.currentDomain, ts.locale, msgID, msgIDPlural, n, args...)
}
func (ts *Translations) X(msgCtxt, msgID string, args ...interface{}) string {
	return ts.DLX(ts.currentDomain, ts.locale, msgCtxt, msgID, args...)
}
func (ts *Translations) XN(msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return ts.DLXN64(ts.currentDomain, ts.locale, msgCtxt, msgID, msgIDPlural, int64(n), args...)
}
func (ts *Translations) XN64(msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return ts.DLXN64(ts.currentDomain, ts.locale, msgCtxt, msgID, msgIDPlural, n, args...)
}

// #endregion gettext

// #region dgettext

func (ts *Translations) DT(domain, msgID string, args ...interface{}) string {
	return ts.DLT(domain, ts.locale, msgID, args...)
}
func (ts *Translations) DN(domain, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return ts.DLN64(domain, ts.locale, msgID, msgIDPlural, int64(n), args...)
}
func (ts *Translations) DN64(domain, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return ts.DLN64(domain, ts.locale, msgID, msgIDPlural, n, args...)
}
func (ts *Translations) DX(domain, msgCtxt, msgID string, args ...interface{}) string {
	return ts.DLX(domain, ts.locale, msgCtxt, msgID, args...)
}
func (ts *Translations) DXN(domain, msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return ts.DLXN64(domain, ts.locale, msgCtxt, msgID, msgIDPlural, int64(n), args...)
}
func (ts *Translations) DXN64(domain, msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return ts.DLXN64(domain, ts.locale, msgCtxt, msgID, msgIDPlural, n, args...)
}

// #endregion dgettext

// #region locale

func (ts *Translations) LT(lang, msgID string, args ...interface{}) string {
	return ts.DLT(ts.currentDomain, lang, msgID, args...)
}
func (ts *Translations) LN(lang, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return ts.DLN64(ts.currentDomain, lang, msgID, msgIDPlural, int64(n), args...)
}
func (ts *Translations) LN64(lang, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return ts.DLN64(ts.currentDomain, lang, msgID, msgIDPlural, n, args...)
}
func (ts *Translations) LX(lang, msgCtxt, msgID string, args ...interface{}) string {
	return ts.DLX(ts.currentDomain, lang, msgCtxt, msgID, args...)
}
func (ts *Translations) LXN(lang, msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return ts.DLXN64(ts.currentDomain, lang, msgCtxt, msgID, msgIDPlural, int64(n), args...)
}
func (ts *Translations) LXN64(lang, msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return ts.DLXN64(ts.currentDomain, lang, msgCtxt, msgID, msgIDPlural, n, args...)
}

// #endregion locale

// #region domain+locale

func (ts *Translations) DLT(domain, lang, msgID string, args ...interface{}) string {
	tr := ts.GetOrNoop(domain)
	return tr.LT(lang, msgID, args...)
}
func (ts *Translations) DLN(domain, lang, msgID, msgIDPlural string, n int, args ...interface{}) string {
	tr := ts.GetOrNoop(domain)
	return tr.LN(lang, msgID, msgIDPlural, n, args...)
}
func (ts *Translations) DLN64(domain, lang, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	tr := ts.GetOrNoop(domain)
	return tr.LN64(lang, msgID, msgIDPlural, n, args...)
}
func (ts *Translations) DLX(domain, lang, msgCtxt, msgID string, args ...interface{}) string {
	tr := ts.GetOrNoop(domain)
	return tr.LX(lang, msgCtxt, msgID, args...)
}
func (ts *Translations) DLXN(domain, lang, msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	tr := ts.GetOrNoop(domain)
	return tr.LXN(lang, msgCtxt, msgID, msgIDPlural, n, args...)
}
func (ts *Translations) DLXN64(domain, lang, msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	tr := ts.GetOrNoop(domain)
	return tr.LXN64(lang, msgCtxt, msgID, msgIDPlural, n, args...)
}

// #endregion domain+locale
