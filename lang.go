package t

// SupportLangs return all registered language
func SupportLangs(domain string) []string {
	return global.SupportLangs(domain)
}

// SetUserLang set user language
func SetUserLang(lang string) {
	global.SetUserLang(lang)
}

// UserLang return the user language
func UserLang() string {
	return global.UserLang()
}

func LT(lang, msgID string, args ...interface{}) string {
	return DLT(global.currentDomain, lang, msgID, args...)
}
func LN(lang, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return DLN(global.currentDomain, lang, msgID, msgIDPlural, n, args...)
}
func LN64(lang, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return DLN64(global.currentDomain, lang, msgID, msgIDPlural, n, args...)
}
func LX(lang, msgCtxt, msgID string, args ...interface{}) string {
	return DLX(global.currentDomain, lang, msgCtxt, msgID, args...)
}
func LXN(lang, msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return DLXN(global.currentDomain, lang, msgCtxt, msgID, msgIDPlural, n, args...)
}
func LXN64(lang, msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return DLXN64(global.currentDomain, lang, msgCtxt, msgID, msgIDPlural, n, args...)
}
