package internal

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
		resolveOneFile("testdata/base.tmpl", &Context{
			Left:  "{{",
			Right: "}}",
			Keywords: []Keyword{
				{Name: "T", MsgID: 1},
				{Name: "N", MsgID: 1, MsgID2: 2},
				{Name: "X", MsgCtxt: 1, MsgID: 2},
				{Name: "XN", MsgCtxt: 1, MsgID: 2, MsgID2: 3},
			},
			Fun:   []string{"T", "X"},
			Debug: true,
		})
	})
}
func Test_run(t *testing.T) {
	Convey("run", t, func() {
		Run(&Context{
			Input: "testdata/*",
			Left:  "{{",
			Right: "}}",
			Keywords: []Keyword{
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
			Fun:    []string{"T"},
			Output: os.Stdout,
		})
	})
}
