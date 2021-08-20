package translator

type Translator interface {
	Lang() string
	T(msgID string, args ...interface{}) string
	N(msgID, msgIDPlural string, n int, args ...interface{}) string
	N64(msgID, msgIDPlural string, n int64, args ...interface{}) string
	X(msgCtxt, msgID string, args ...interface{}) string
	XN(msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string
	XN64(msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string
}

type File struct {
	entries []*Entry
	headers map[string]string
}

type Entry struct {
	comments []string
	msgCtxt  string
	msgID    string
	msgID2   string
	msgStr   string
	msgStrN  []string
}
