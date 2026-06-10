package internal

import (
	"os"
	"strconv"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/youthlin/t"
)

// Keyword gettext keyword
type Keyword struct {
	Name    string
	MsgCtxt int
	MsgID   int
	MsgID2  int
}

// ParseKeywords 解析 xgettext 风格的关键字声明，例如：
// gettext;T:1;N:1,2;X:1c,2;XN:1c,2,3
func ParseKeywords(str string) (result []Keyword, err error) {
	kw := strings.Split(str, ";")
	msg := t.T("invalid keywords: %s", str)
	for _, key := range kw {
		key = strings.TrimSpace(key)
		// T
		// T:1
		// N:1,2
		// X:1c,2
		// XN:1c,2,3
		nameIndex := strings.Split(key, ":")
		if len(nameIndex) == 1 {
			name := strings.TrimSpace(nameIndex[0])
			if name == "" {
				return nil, errors.Errorf(msg)
			}
			result = append(result, Keyword{Name: name, MsgID: 1})
			continue
		}
		if len(nameIndex) != 2 {
			return nil, errors.Errorf(msg)
		}
		k := Keyword{
			Name: strings.TrimSpace(nameIndex[0]),
		}
		if k.Name == "" {
			return nil, errors.Errorf(msg)
		}
		index := strings.Split(nameIndex[1], ",")
		for i := range index {
			index[i] = strings.TrimSpace(index[i])
		}
		switch len(index) {
		case 1:
			i, err := strconv.ParseInt(index[0], 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, msg+t.T("msg id index is not a number"))
			}
			k.MsgID = int(i)
		case 2:
			i1 := index[0]
			i2 := index[1]
			i1c := strings.HasSuffix(i1, "c")
			if i1c {
				c := i1[:len(i1)-1]
				cIndex, err := strconv.ParseInt(c, 10, 64)
				if err != nil {
					return nil, errors.Wrapf(err, msg+t.T("context index is not a number: %v", c))
				}
				k.MsgCtxt = int(cIndex)

				index, err := strconv.ParseInt(i2, 10, 64)
				if err != nil {
					return nil, errors.Wrapf(err, msg+t.T("msg id index is not a number: %v", i2))
				}
				k.MsgID = int(index)
			} else {
				index, err := strconv.ParseInt(i1, 10, 64)
				if err != nil {
					return nil, errors.Wrapf(err, msg+t.T("msg id index is not a number: %v", i1))
				}
				k.MsgID = int(index)

				index, err = strconv.ParseInt(i2, 10, 64)
				if err != nil {
					return nil, errors.Wrapf(err, msg+t.T("msg plural index is not a number: %v", i2))
				}
				k.MsgID2 = int(index)
			}
		case 3:
			i1 := index[0]
			i2 := index[1]
			i3 := index[2]
			if !strings.HasSuffix(i1, "c") {
				return nil, errors.Errorf(msg + t.T("context index must end with 'c': %v", i1))
			}
			c := i1[:len(i1)-1]
			index, err := strconv.ParseInt(c, 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, msg+t.T("msg context index is not a number: %v", c))
			}
			k.MsgCtxt = int(index)

			index, err = strconv.ParseInt(i2, 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, msg+t.T("msg id index is not a number: %v", i2))
			}
			k.MsgID = int(index)

			index, err = strconv.ParseInt(i3, 10, 64)
			if err != nil {
				return nil, errors.Wrapf(err, msg+t.T("msg id index is not a number: %v", i3))
			}
			k.MsgID2 = int(index)
		default:
			return nil, errors.Errorf(msg + t.T("too much keyword index"))
		}
		result = append(result, k)
	}
	return
}

// Writer 根据输出参数返回写入目标。
// fileName 为空或为 - 时输出到 stdout，否则覆盖写入指定文件。
func Writer(fileName string) (wr *os.File, err error) {
	wr = os.Stdout
	if fileName != "" && fileName != "-" {
		wr, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			err = errors.Wrapf(err, t.T("can not open output file: %s", fileName))
		}
	}
	return
}
