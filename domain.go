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
	global.Bind(domain, path)
}

// TextDomain if the domain has bind some translation, return domain, othewise return DefaultDomain(default)
func TextDomain(domain string) string {
	return global.TextDomain(domain)
}

// NewDomain return a new Translations instance which current domain is set to the specify domain
func NewDomain(domain string) *Translations {
	return global.NewDomain(domain)
}

// DT see T. domain was bind at BindTextDomain
// dgettext
func DT(domain, msgID string, args ...interface{}) string {
	return global.DT(domain, msgID, args...)
}

// DN see N.
// dngettext
func DN(domain, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return global.DN(domain, msgID, msgIDPlural, n, args...)
}

// DN64 int64 version of DN
func DN64(domain, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return global.DN64(domain, msgID, msgIDPlural, n, args...)
}

// DX see X.
// dpgettext
func DX(domain, msgCtxt, msgID string, args ...interface{}) string {
	return global.DX(domain, msgCtxt, msgID, args...)
}

// DXN see XN.
// dpngettext
func DXN(domain, msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return global.DXN(domain, msgCtxt, msgID, msgIDPlural, n, args...)
}

// DXN64 int64 version of DXN
func DXN64(domain, msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return global.DXN64(domain, msgCtxt, msgID, msgIDPlural, n, args...)
}

// #region domain+locale

func DLT(domain, lang, msgID string, args ...interface{}) string {
	return global.DLT(domain, lang, msgID, args...)
}
func DLN(domain, lang, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return global.DLN(domain, lang, msgID, msgIDPlural, n, args...)
}
func DLN64(domain, lang, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return global.DLN64(domain, lang, msgID, msgIDPlural, n, args...)
}
func DLX(domain, lang, msgCtxt, msgID string, args ...interface{}) string {
	return global.DLX(domain, lang, msgCtxt, msgID, args...)
}
func DLXN(domain, lang, msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return global.DLXN(domain, lang, msgCtxt, msgID, msgIDPlural, n, args...)
}
func DLXN64(domain, lang, msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return global.DLXN64(domain, lang, msgCtxt, msgID, msgIDPlural, n, args...)
}

// #endregion domain+locale
