package t

// noop return msg directly, used to mark string which should be translated
// so that xgettext tool can extract those strings.
type noop int

var Noop = noop(0)

func (p noop) T(msgID string) string {
	return msgID
}
func (p noop) N(msgID string, msgIDPlural string) (string, string) {
	return msgID, msgIDPlural
}
func (p noop) X(msgCtxt string, msgID string) (string, string) {
	return msgCtxt, msgID
}
func (p noop) XN(msgCtxt string, msgID string, msgIDPlural string) (string, string, string) {
	return msgCtxt, msgID, msgIDPlural
}
