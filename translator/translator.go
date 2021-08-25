package translator

// Translator 翻译家接口
type Translator interface {
	// Lang return the language 返回翻译之后的语言
	Lang() string
	// X short name of pgettext. msgCtxt can be empty. 单数翻译接口
	X(msgCtxt, msgID string, args ...interface{}) string
	// XN64 short name of npgettext. msgCtxt can be empty. 复数翻译接口
	XN64(msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string
}
