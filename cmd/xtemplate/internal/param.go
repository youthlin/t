package internal

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/youthlin/t"
)

// Context parameters
type Context struct {
	Input    string
	Left     string
	Right    string
	Keywords []Keyword
	Fun      []string
	Output   io.Writer
	Debug    bool
	result   pot
}

// debugPrint print if is debug mode
func (p *Context) debugPrint(format string, args ...interface{}) {
	if p.Debug {
		fmt.Printf(format+"\n", args...)
	}
}

// Keyword gettext keyword
type Keyword struct {
	Name    string
	MsgCtxt int
	MsgID   int
	MsgID2  int
}

// parseKeywords gettext;T:1;N:1,2;X:1c,2;XN:1c,2,3
func ParseKeywords(str string) (result []Keyword, err error) {
	kw := strings.Split(str, ";")
	msg := t.T("invalid keywords: %s", str)
	for _, key := range kw {
		// T
		// T:1
		// N:1,2
		// X:1c,2
		// XN:1c,2,3
		nameIndex := strings.Split(key, ":")
		if len(nameIndex) == 1 {
			name := nameIndex[0]
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
			Name: nameIndex[0],
		}
		index := strings.Split(nameIndex[1], ",")
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
			return nil, errors.Errorf(msg + t.T("tow much keyword index"))
		}
		result = append(result, k)
	}
	return
}

// writer is fileName is empty or - use stdout, otherwise use file
func Writer(fileName string) (wr *os.File, err error) {
	wr = os.Stdout
	if fileName != "" && fileName != "-" {
		wr, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			err = errors.Wrapf(err, t.T("can not open output file"))
		}
	}
	return
}
