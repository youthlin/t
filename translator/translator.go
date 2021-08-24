package translator

// Translator 翻译家接口
type Translator interface {
	Lang() string
	X(msgCtxt, msgID string, args ...interface{}) string
	XN64(msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string
}
