package t

// Mark is used to mark translation texts
var Mark = noop(0)

// noop return msg directly, used to mark string which should be translated
// so that xgettext tool can extract those strings.
type noop int

func (noop) T(msgID string) string { return msgID }
func (noop) N(msgID string, msgIDPlural string) (string, string) {
	return msgID, msgIDPlural
}
func (noop) X(msgCtxt string, msgID string) (string, string) {
	return msgCtxt, msgID
}
func (noop) XN(msgCtxt string, msgID string, msgIDPlural string) (string, string, string) {
	return msgCtxt, msgID, msgIDPlural
}
