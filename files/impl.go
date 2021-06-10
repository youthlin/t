package files

import "github.com/youthlin/t/f"

// Lang 返回译文的语言
func (f *File) Lang() string {
	if f.lang == "" {
		lang, _ := f.GetHeader(HeaderLanguage)
		f.SetLang(lang)
	}
	return f.lang
}

// T gettext 直接获取翻译内容，如果没有翻译，返回原始内容
// 如果 args 不为空，则将翻译后的字符串作为格式化模版，格式化 args
func (file *File) T(msgID string, args ...interface{}) string {
	return file.X("", msgID, args...)
}

// N ngettext 翻译复数，如果没有翻译，n 大于 1 返回原文复数(msgIDPlural)，否则返回原文单数(msgID)
// 如果 args 不为空，则将翻译后的字符串作为格式化模版，格式化 args
// 注意，n 用于选择第几种复数，如果 需要打印 n，还需要将其传包括在 args 中.
//
// Note: n is used to choose plural forms, is you need print n, you should pass it to args
// 	// no args, so return: `%d apples`
// 	po.N("one apple", "%d apples", 2) -> "%d apples"
// 	// the first numer 2 result in `%d apples`, the second 2 format to `2 apples`
// 	po.N("one apple", "%d apples", 2, 2) -> "2 apples"
// 	po.N("one apple", "%d apples", 2, 200) -> "200 apples"
//
// 	// use `one apple` as template to format number `200`, the extra arg ignored, see f.Format
// 	po.N("one apple", "%d apples", 1, 200) -> "one apple"
func (file *File) N(msgID, msgIDPlural string, n int, args ...interface{}) string {
	return file.XN64("", msgID, msgIDPlural, int64(n), args...)
}

// N64 ngettext 翻译复数，如果没有翻译，n 大于 1 返回原文复数(msgIDPlural)，否则返回原文单数(msgID)
func (file *File) N64(msgID, msgIDPlural string, n int64, args ...interface{}) string {
	return file.XN64("", msgID, msgIDPlural, n, args...)
}

// X pgettext 带上下文翻译，用于区分同一个 msgID 在不同上下文的不同含义
func (file *File) X(msgCtxt, msgID string, args ...interface{}) string {
	msg, ok := file.messages[key(msgCtxt, msgID)]
	if !ok || msg.MsgStr == "" {
		return f.Format(msgID, args...)
	}
	return f.Format(msg.MsgStr, args...)
}

// XN pngettext 带上下文翻译复数
func (file *File) XN(msgCtxt, msgID, msgIDPlural string, n int, args ...interface{}) string {
	return file.XN64(msgCtxt, msgID, msgIDPlural, int64(n), args...)
}

// XN64 pngettext 带上下文翻译复数
func (file *File) XN64(msgCtxt, msgID, msgIDPlural string, n int64, args ...interface{}) string {
	msg, ok := file.messages[key(msgCtxt, msgID)]
	if !ok {
		return f.DefaultPlural(msgID, msgIDPlural, n, args...)
	}
	totalForms := file.GetTotalForms()
	pluralFunc := file.GetPluralFunc()
	if totalForms <= 0 || pluralFunc == nil { // 复数设置不正确
		return f.DefaultPlural(msgID, msgIDPlural, n, args...)
	}
	// 看 n 对应第几种复数
	index := pluralFunc(n)
	if index < 0 || index >= int(totalForms) || index > len(msg.MsgStrN) || msg.MsgStrN[index] == "" {
		// 超出范围
		return f.DefaultPlural(msgID, msgIDPlural, n, args...)
	}
	return f.Format(msg.MsgStrN[index], args...)
}
