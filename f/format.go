package f

import "fmt"

// Format format a string with args, like fmt.Sprintf,
// but if args tow many, not prints %!(EXTRA type=value);
// and if no args, will return original string,
// even it contains some verb(like %v/%d), it would not prints MISSING error.
// 格式化字符串，功能同 fmt.Sprintf, 但是当参数多于占位符时，
// 不会输出额外的 %!(EXTRA type=value)；
// 当 args 为空时直接返回原字符串（若包含格式化动词也原样返回而不会打印 MISSING 错误）
func Format(format string, args ...any) string {
	var length = len(args)
	if length == 0 {
		return format
	}
	// 原理：使用索引指定参数位置，在 args 后拼接一个空白字符串，
	// 然后在格式化字符串上使用 %[n]s 输出拼接的空白字符串，这样就没有多余的参数了
	// see fmt 包注释，或中文文档 https://studygolang.com/static/pkgdoc/pkg/fmt.htm
	args = append(args, "")
	format = fmt.Sprintf("%s%%[%d]s", format, length+1)
	return fmt.Sprintf(format, args...)
}

// DefaultPlural 根据 n 选择单数或复数形式。
// 当 msgIDPlural 为空时（例如源语言本身没有复数区分），直接使用 msgID。
func DefaultPlural(msgID, msgIDPlural string, n int64, args ...any) string {
	if msgIDPlural == "" {
		return Format(msgID, args...)
	}
	if n != 1 {
		return Format(msgIDPlural, args...)
	}
	return Format(msgID, args...)
}
