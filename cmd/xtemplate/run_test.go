package main

import (
	"os"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGlob(t *testing.T) {
	Convey("glob", t, func() {
		filenames, err := filepath.Glob("testdata/*")
		So(err, ShouldBeNil)
		t.Logf("files: %v", filenames)
	})
}

func TestFile(t *testing.T) {
	Convey("resolveOneFile", t, func() {
		*debug = true
		resolveOneFile("testdata/base.tmpl", &Param{
			left: "{{", right: "}}",
			keywords: []Keyword{
				{Name: "T"},
				{Name: "X"},
			},
			fun: []string{"T", "X"},
		})
	})
}
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
			fun:    []string{"T"},
			output: os.Stdout,
		})
	})
}
