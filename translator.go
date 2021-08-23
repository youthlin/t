package t

import (
	"github.com/youthlin/t/f"
	"github.com/youthlin/t/translator"
)

type Translator = translator.Translator

func NoopTranslator() Translator { return noopTranslator }

var noopTranslator Translator = nooptor{}

type nooptor struct{}

func (tor nooptor) Lang() string { return "" }
func (tor nooptor) T(msgID string, args ...interface{}) string {
	return f.Format(msgID, args...)
}
func (tor nooptor) N(msgID, msgIDPlural string, n int, args ...interface{}) string {
	return f.DefaultPlural(msgID, msgIDPlural, int64(n), args...)
}
func (tor nooptor) N64(msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return f.DefaultPlural(msgID, msgIDPlural, n, args...)
}
func (tor nooptor) X(msgCtxt, msgID string, args ...interface{}) string {
	return f.Format(msgID, args...)
}
func (tor nooptor) XN(msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return f.DefaultPlural(msgID, msgIDPlural, int64(n), args...)
}
func (tor nooptor) XN64(msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return f.DefaultPlural(msgID, msgIDPlural, n, args...)
}
