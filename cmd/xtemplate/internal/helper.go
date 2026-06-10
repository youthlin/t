package internal

import (
	"strings"

	"github.com/youthlin/t/translator"
)

// eot 使用 gettext 常见的 EOT 分隔符把 msgctxt 和 msgid 拼成 map key。
// 这样既能复用一个 map，又不会把不同上下文但相同文案的条目混在一起。
const eot = "\x04"

// isPlural 判断当前条目是否包含复数形式。
func isPlural(e *translator.Entry) bool {
	return e.MsgID2 != ""
}

// key 生成内部去重使用的唯一键：msgctxt + EOT + msgid。
func key(ctxt, msgid string) string {
	return ctxt + eot + msgid
}

// isGoFormat 粗略判断是否可能是 fmt 风格占位符。
// 这里故意保持简单：宁可多打一个 go-format 标记，也不要漏掉明显的格式化字符串。
func isGoFormat(e *translator.Entry) bool {
	return strings.Contains(e.MsgID, "%") || strings.Contains(e.MsgID2, "%")
}
