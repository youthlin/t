package main

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_run(t *testing.T) {
	Convey("run", t, func() {
		run(&Param{
			input: "testdata/*",
			left:  "{{",
			right: "}}",
			keywords: []Keyword{
				{
					Name:  "T",
					MsgID: 1,
				},
				{
					Name:    "X",
					MsgCtxt: 1,
					MsgID:   2,
				},
				{
					Name:   "N",
					MsgID:  1,
					MsgID2: 2,
				},
				{
					Name:    "XN",
					MsgCtxt: 1,
					MsgID:   2,
					MsgID2:  3,
				},
			},
			output: os.Stdout,
		})
	})
}
