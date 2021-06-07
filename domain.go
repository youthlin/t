package t

// DefaultDomain the default domain: "default"
const DefaultDomain = "default"

// BindDefaultDomain bind 'default' domain to the dir of po/mo file.
// p.s. the path can be a po/mo file as well.
//  - path is dir: search all po/mo files in dir and register them;
//  - path is file: if it's a po or mo file, register it
func BindDefaultDomain(path string) {
	BindTextDomain(DefaultDomain, path)
}

// BindTextDomain bind domain in path
func BindTextDomain(domain string, path string) {
	globalTranslatins.Bind(domain, path)
}

// TextDomain if the domain has bind some translation, return domain, othewise return DefaultDomain(default)
func TextDomain(domain string) string {
	return globalTranslatins.TextDomain(domain)
}

// getTranslation get a Traslation instance which bind the domain
func getTranslation(domain string) (*Translation, string) {
	tr := globalTranslatins.GetOrNoop(domain)
	lang := globalTranslatins.UserLang()
	return tr, lang
}

// DT see T. domain was bind at BindTextDomain
// dgettext
func DT(domain, msgID string, args ...interface{}) string {
	tr, lang := getTranslation(domain)
	return tr.LT(lang, msgID, args...)
}

// DN see N.
// dngettext
func DN(domain, msgID, msgIDPlural string, n int, args ...interface{}) string {
	tr, lang := getTranslation(domain)
	return tr.LN(lang, msgID, msgIDPlural, n, args...)
}

// DN64 int64 version of DN
func DN64(domain, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	tr, lang := getTranslation(domain)
	return tr.LN64(lang, msgID, msgIDPlural, n, args...)
}

// DX see X.
// dpgettext
func DX(domain, msgCtxt, msgID string, args ...interface{}) string {
	tr, lang := getTranslation(domain)
	return tr.LX(lang, msgCtxt, msgID, args...)
}

// DXN see XN.
// dpngettext
func DXN(domain, msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	tr, lang := getTranslation(domain)
	return tr.LXN(lang, msgCtxt, msgID, msgIDPlural, n, args...)
}

// DXN64 int64 version of DXN
func DXN64(domain, msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	tr, lang := getTranslation(domain)
	return tr.LXN64(lang, msgCtxt, msgID, msgIDPlural, n, args...)
}

func DLT(domain, lang, msgID string, args ...interface{}) string {
	tr := globalTranslatins.GetOrNoop(domain)
	return tr.LT(lang, msgID, args...)
}
func DLN(domain, lang, msgID, msgIDPlural string, n int, args ...interface{}) string {
	tr := globalTranslatins.GetOrNoop(domain)
	return tr.LN(lang, msgID, msgIDPlural, n, args...)
}
func DLN64(domain, lang, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	tr := globalTranslatins.GetOrNoop(domain)
	return tr.LN64(lang, msgID, msgIDPlural, n, args...)
}
func DLX(domain, lang, msgCtxt, msgID string, args ...interface{}) string {
	tr := globalTranslatins.GetOrNoop(domain)
	return tr.LX(lang, msgCtxt, msgID, args...)
}
func DLXN(domain, lang, msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	tr := globalTranslatins.GetOrNoop(domain)
	return tr.LXN(lang, msgCtxt, msgID, msgIDPlural, n, args...)
}
func DLXN64(domain, lang, msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	tr := globalTranslatins.GetOrNoop(domain)
	return tr.LXN64(lang, msgCtxt, msgID, msgIDPlural, n, args...)
}
