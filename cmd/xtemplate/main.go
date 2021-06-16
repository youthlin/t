package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/youthlin/t"
)

// xtemplate -k.t.T:1 -o file.pot -i input-pattern

var input = flag.String("i", "", t.T("input file pattern"))
var left = flag.String("left", "{{", t.T("left delim"))
var right = flag.String("right", "}}", t.T("right delim"))
var keywords = flag.String("k", "", t.T("keywords e.g.: T:1;N1,2;X:1c,2;XN:1c,2,3"))
var fun = flag.String("f", "", t.T("function names of template"))
var output = flag.String("o", "message.pot", t.T("output file"))
var comment = flag.String("c", "\x00", t.T("extract comment"))

var help = flag.Bool("h", false, t.T("show this help message"))
var debug = flag.Bool("d", false, t.T("debug mode"))
var version = flag.Bool("v", false, t.T("show version"))

func main() {
	flag.Parse()
	if *help || len(os.Args) < 5 {
		flag.Usage()
		return
	}
	if *version {
		fmt.Fprintf(os.Stdout, t.T("version: %v\n"), "v0.0.0")
		return
	}
	defer func() {
		if e := recover(); e != nil {
			fmt.Fprintf(os.Stderr, t.T("unexpected error: %v"), e)
			flag.Usage()
		}
	}()

	run(&Param{
		input:         *input,
		left:          *left,
		right:         *right,
		keywords:      parseKeywords(),
		fun:           strings.Split(*fun, ","),
		needComment:   *comment != "\x00",
		commentPrefix: *comment,
		output:        writer(),
	})
}

type Param struct {
	input         string
	left          string
	right         string
	keywords      []Keyword
	fun           []string
	needComment   bool
	commentPrefix string
	output        io.Writer
}
type Keyword struct {
	Name    string
	MsgCtxt int
	MsgID   int
	MsgID2  int
}

func parseKeywords() (result []Keyword) {
	kw := strings.Split(*keywords, ";")
	msg := t.T("invalid keywords: %s\n", *keywords)
	for _, key := range kw {
		// .t.T:1c,2
		nameIndex := strings.Split(key, ":")
		if len(nameIndex) != 2 {
			panic(msg)
		}
		k := Keyword{
			Name: nameIndex[0],
		}
		index := strings.Split(nameIndex[1], ",")
		switch len(index) {
		case 1:
			i, err := strconv.ParseInt(index[0], 10, 64)
			if err != nil {
				panic(msg + t.T("msg id index is not a number: %v", err))
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
					panic(msg + t.T("context index is not a number: %v|%v", c, err))
				}
				k.MsgCtxt = int(cIndex)

				index, err := strconv.ParseInt(i2, 10, 64)
				if err != nil {
					panic(msg + t.T("msg id index is not a number: %v|%v", i2, err))
				}
				k.MsgID = int(index)
			} else {
				index, err := strconv.ParseInt(i1, 10, 64)
				if err != nil {
					panic(msg + t.T("msg id index is not a number: %v|%v", i1, err))
				}
				k.MsgID = int(index)

				index, err = strconv.ParseInt(i2, 10, 64)
				if err != nil {
					panic(msg + t.T("msg plural index is not a number: %v|%v", i2, err))
				}
				k.MsgID2 = int(index)
			}
		case 3:
			i1 := index[0]
			i2 := index[1]
			i3 := index[1]
			if !strings.HasSuffix(i1, "c") {
				panic(msg + t.T("context index must suffix by 'c': %v", i1))
			}
			c := i1[:len(i1)-1]
			index, err := strconv.ParseInt(c, 10, 64)
			if err != nil {
				panic(msg + t.T("msg context index is not a number: %v|%v", c, err))
			}
			k.MsgCtxt = int(index)

			index, err = strconv.ParseInt(i2, 10, 64)
			if err != nil {
				panic(msg + t.T("msg id index is not a number: %v|%v", i2, err))
			}
			k.MsgID = int(index)

			index, err = strconv.ParseInt(i3, 10, 64)
			if err != nil {
				panic(msg + t.T("msg id index is not a number: %v|%v", i3, err))
			}
			k.MsgID2 = int(index)
		default:
			panic(msg + t.T("tow much keyword index"))
		}
		result = append(result, k)
	}
	return
}

func writer() io.Writer {
	var wr = os.Stdout
	var err error
	fileName := *output
	if fileName != "" && fileName != "-" {
		wr, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(t.T("can not open output file: %v", err))
		}
	}
	return wr
}
