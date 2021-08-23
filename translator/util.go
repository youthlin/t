package translator

import (
	"strconv"
	"strings"
)

// key helper function, return message key
func key(ctxt, msgid string) string {
	return ctxt + eot + msgid
}

// removePrefixAndUnquote 去掉前缀并返回引号中的内容
func removePrefixAndUnquote(line, prefix string) (string, error) {
	line = strings.TrimPrefix(line, prefix)
	line = strings.TrimSpace(line)
	return strconv.Unquote(line)
}
