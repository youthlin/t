package f

import "fmt"

// Format format a string with args, like fmt.Sprintf,
// but if args tow many, not prints %!(EXTRA type=value)
// 格式化字符串，功能同 fmt.Sprintf, 但是当参数多于占位符时，
// 不会输出额外的 %!(EXTRA type=value)
func Format(format string, args ...interface{}) string {
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
