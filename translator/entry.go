package translator

import (
	"fmt"
	"strconv"
	"strings"
)

// Entry 一个翻译条目
type Entry struct {
	MsgCmts []string
	MsgCtxt string
	MsgID   string
	MsgID2  string
	MsgStr  string
	MsgStrN []string
}

// Key 一个翻译条目的 Key
func (e *Entry) Key() string {
	return key(e.MsgCtxt, e.MsgID)
}

// isHeader 返回该条目是否是一个 header 条目
func (e *Entry) isHeader() bool {
	return e.MsgID == ""
}

// isValid 是否合法的条目
func (e *Entry) isValid() bool {
	// header entry: msgid == ""
	return e != nil && (e.MsgStr != "" || len(e.MsgStrN) > 0)
}

// getSortKey 排序
func (e *Entry) getSortKey() string {
	if e.isHeader() {
		return "" // 排在最前
	}
	return e.getLineString() + e.Key()
}

// getLineString 如果有行号注释按行号排序 非行号的注释不使用
func (e *Entry) getLineString() string {
	var ss []string
	for _, cmt := range e.MsgCmts {
		// #: testdata/index.html:16:28 testdata/index.html:18
		if strings.HasPrefix(cmt, "#:") { // 行号注释前缀
			cmt = strings.TrimPrefix(cmt, "#:")
			var s []string
			s = append(s, "#:")
			for _, item := range strings.Split(cmt, " ") { // 多个位置
				pair := strings.Split(item, ":") // 文件名:行号[:列号]
				result := make([]string, 0, len(pair))
				for _, n := range pair {
					i, err := strconv.ParseInt(n, 10, 64)
					if err != nil {
						result = append(result, n)
					} else { // 数字添加前导 0，这样字符串比较时 012<123
						result = append(result, fmt.Sprintf("%10d", i))
					}
				}
				s = append(s, result...)
			}
			ss = append(ss, s...)
		}
	}
	return strings.Join(ss, "")
}
