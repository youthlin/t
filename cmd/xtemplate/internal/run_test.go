package internal

import (
	"html/template"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/youthlin/t/translator"
)

func TestGlob(t *testing.T) {
	Convey("glob", t, func() {
		filenames, err := filepath.Glob("testdata/*.tmpl")
		So(err, ShouldBeNil)
		t.Logf("files: %v", filenames)
	})
}

func TestFile(t *testing.T) {
	Convey("resolveOneFile", t, func() {
		resolveOneFile("testdata/base.tmpl", &Context{
			Param: &Param{
				Left:  "{{",
				Right: "}}",
				Debug: true,
			},
			Keywords: []Keyword{
				{Name: "T", MsgID: 1},
				{Name: "N", MsgID: 1, MsgID2: 2},
				{Name: "X", MsgCtxt: 1, MsgID: 2},
				{Name: "XN", MsgCtxt: 1, MsgID: 2, MsgID2: 3},
			},
			Functions: template.FuncMap{"T": noopFun, "X": noopFun},
			entries:   make(map[string]*translator.Entry),
		})
	})
}
func Test_run(t *testing.T) {
	Convey("run", t, func() {
		Run(&Param{
			Input:      "testdata/*.tmpl",
			Left:       "{{",
			Right:      "}}",
			Keyword:    "T;X:1c,2;N:1,2;XN:1c,2,3",
			Function:   "T",
			OutputFile: "-",
		})
	})
}
