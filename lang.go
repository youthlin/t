package t

// SupportLangs return all registered language
func SupportLangs(domain string) []string {
	return global.SupportLangs(domain)
}

// SetLocale set user language
func SetLocale(lang string) {
	global.SetLocale(lang)
}

// Locale return the user language
func Locale() string {
	return global.Locale()
}

// UseLocale return a new Translations instance of locale
func UseLocale(lang string) *Translations {
	return global.UseLocale(lang)
}

func LT(lang, msgID string, args ...interface{}) string {
	return global.LT(lang, msgID, args...)
}
func LN(lang, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return global.LN(lang, msgID, msgIDPlural, n, args...)
}
func LN64(lang, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return global.LN64(lang, msgID, msgIDPlural, n, args...)
}
func LX(lang, msgCtxt, msgID string, args ...interface{}) string {
	return global.LX(lang, msgCtxt, msgID, args...)
}
func LXN(lang, msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return global.LXN(lang, msgCtxt, msgID, msgIDPlural, n, args...)
}
func LXN64(lang, msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return global.LXN64(lang, msgCtxt, msgID, msgIDPlural, n, args...)
}
