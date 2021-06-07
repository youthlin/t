package t

// T mark msgID should be translated, and return tranlated msgstr.
// If no translation of msgID, return msgID itself
// if args are not empty, will format with args
// gettext
func T(msgID string, args ...interface{}) string {
	return DT(DefaultDomain, msgID, args...)
}

// N like T, and support plural forms. the integer n was used to be choose plural forms
// if you want format n, you should contains it in args, too.
// ngettext
func N(msgID, msgIDPlural string, n int, args ...interface{}) string {
	return DN(DefaultDomain, msgID, msgIDPlural, n, args...)
}

// N64 like N, but the integer n is a int64
func N64(msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return DN64(DefaultDomain, msgID, msgIDPlural, n, args...)
}

// X like T, and this function support pass a context text to disambiguation
// pgettext
func X(msgCtxt, msgID string, args ...interface{}) string {
	return DX(DefaultDomain, msgCtxt, msgID, args...)
}

// XN see X, N
// npgettext
func XN(msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return DXN(DefaultDomain, msgCtxt, msgID, msgIDPlural, n, args...)
}

// XN64 int64 version of XN
func XN64(msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return DXN64(DefaultDomain, msgCtxt, msgID, msgIDPlural, n, args...)
}
