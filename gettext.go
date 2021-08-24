package t

// T: shor name of gettext
func T(msgID string, args ...interface{}) string {
	return global.X("", msgID, args...)
}

// N is a short name of nettext
func N(msgID, msgIDPlural string, n int, args ...interface{}) string {
	return global.XN64("", msgID, msgIDPlural, int64(n), args...)
}

// N64 is a short name of nettext
func N64(msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return global.XN64("", msgID, msgIDPlural, n, args...)
}

// X is a short name of pgettext
func X(msgCtxt, msgID string, args ...interface{}) string {
	return global.X(msgCtxt, msgID, args...)
}

// XN is a short name of npgettext
func XN(msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return global.XN64(msgCtxt, msgID, msgIDPlural, int64(n), args...)
}

// XN64 is a short name of npgettext
func XN64(msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return global.XN64(msgCtxt, msgID, msgIDPlural, n, args...)
}
