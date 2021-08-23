package translator

// Translator 翻译家接口
type Translator interface {
	Lang() string
	T(msgID string, args ...interface{}) string
	N(msgID, msgIDPlural string, n int, args ...interface{}) string
	N64(msgID, msgIDPlural string, n int64, args ...interface{}) string
	X(msgCtxt, msgID string, args ...interface{}) string
	XN(msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string
	XN64(msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string
}
