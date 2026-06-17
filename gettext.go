package t

// D return a new Translations with domain
func D(domain string) *Translations { return global.D(domain) }

// L return a new Translations with locale
func L(locale string) *Translations { return global.L(locale) }

// T is a short name of gettext
func T(msgID string, args ...any) string {
	return global.X("", msgID, args...)
}

// N is a short name of ngettext
func N(msgID, msgIDPlural string, n int, args ...any) string {
	return global.XN64("", msgID, msgIDPlural, int64(n), args...)
}

// N1 用于单复数同形的简写,
// 比如 N("%d 个苹果", "%d 个苹果", n, n) 可以简写为 N1("%d 个苹果", n, n)
// 同时抽取脚本使用 -kN1:1,1 指定, 可正常抽取出 msgid_plural
func N1(msgId string, n int, args ...any) string {
	return global.XN64("", msgId, "", int64(n), args...)
}

// N64 is a short name of ngettext
func N64(msgID, msgIDPlural string, n int64, args ...any) string {
	return global.XN64("", msgID, msgIDPlural, n, args...)
}

// X is a short name of pgettext
func X(msgCtxt, msgID string, args ...any) string {
	return global.X(msgCtxt, msgID, args...)
}

// XN is a short name of npgettext
func XN(msgCtxt, msgID, msgIDPlural string, n int, args ...any) string {
	return global.XN64(msgCtxt, msgID, msgIDPlural, int64(n), args...)
}

// XN64 is a short name of npgettext
func XN64(msgCtxt, msgID, msgIDPlural string, n int64, args ...any) string {
	return global.XN64(msgCtxt, msgID, msgIDPlural, n, args...)
}
