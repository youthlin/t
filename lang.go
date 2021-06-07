package t

// SupportLangs return all registered language
func SupportLangs(domain string) []string {
	return globalTranslatins.SupportLangs(domain)
}

// SetUserLang set user language
func SetUserLang(lang string) {
	globalTranslatins.SetUserLang(lang)
}

// UserLang return the user language
func UserLang() string {
	return globalTranslatins.UserLang()
}

func LT(lang, msgID string, args ...interface{}) string {
	return DLT(DefaultDomain, lang, msgID, args...)
}
func LN(lang, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return DLN(DefaultDomain, lang, msgID, msgIDPlural, n, args...)
}
func LN64(lang, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return DLN64(DefaultDomain, lang, msgID, msgIDPlural, n, args...)
}
func LX(lang, msgCtxt, msgID string, args ...interface{}) string {
	return DLX(DefaultDomain, lang, msgCtxt, msgID, args...)
}
func LXN(lang, msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return DLXN(DefaultDomain, lang, msgCtxt, msgID, msgIDPlural, n, args...)
}
func LXN64(lang, msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return DLXN64(DefaultDomain, lang, msgCtxt, msgID, msgIDPlural, n, args...)
}
