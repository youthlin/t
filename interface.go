package t

// File is a interface, it's impl may a .po file or .mo file
type File interface {
	Lang() string
	T(msgID string, args ...interface{}) string
	N(msgID, msgIDPlural string, n int, args ...interface{}) string
	N64(msgID, msgIDPlural string, n int64, args ...interface{}) string
	X(msgCtxt, msgID string, args ...interface{}) string
	XN(msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string
	XN64(msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string
}
