package internal

import (
	"strings"

	"github.com/youthlin/t/translator"
)

const eot = "\x04"

func isPlural(e *translator.Entry) bool {
	return e.MsgID2 != ""
}

func key(ctxt, msgid string) string {
	return ctxt + eot + msgid
}

func isGoFormat(e *translator.Entry) bool {
	return strings.Contains(e.MsgID, "%") || strings.Contains(e.MsgID2, "%")
}
