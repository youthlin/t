package internal

import (
	"html/template"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/youthlin/t/translator"
)

func newTestContext() *Context {
	return &Context{
		Param: &Param{
			Left:  "{{",
			Right: "}}",
		},
		Keywords: []Keyword{
			{Name: "T", MsgID: 1},
			{Name: "N", MsgID: 1, MsgID2: 2},
			{Name: "X", MsgCtxt: 1, MsgID: 2},
			{Name: "XN", MsgCtxt: 1, MsgID: 2, MsgID2: 3},
		},
		Functions: template.FuncMap{"T": noopFun, "X": noopFun},
		entries:   make(map[string]*translator.Entry),
	}
}

// TestGlob 只验证测试数据模式能正常匹配，避免后续因为路径调整导致测试数据没被读到。
func TestGlob(t *testing.T) {
	Convey("glob", t, func() {
		filenames, err := filepath.Glob("testdata/*.tmpl")
		So(err, ShouldBeNil)
		t.Logf("files: %v", filenames)
	})
}

// TestFile 验证 resolveOneFile 能从控制结构、管道和括号包裹参数中抽取翻译文本。
func TestFile(t *testing.T) {
	Convey("resolveOneFile", t, func() {
		ctx := newTestContext()
		So(resolveOneFile("testdata/base.tmpl", ctx), ShouldBeNil)
		So(ctx.entries[key("", "inside if")], ShouldNotBeNil)
		So(ctx.entries[key("", "inside else")], ShouldNotBeNil)
		So(ctx.entries[key("", "inside range")], ShouldNotBeNil)
		So(ctx.entries[key("", "inside range else")], ShouldNotBeNil)
		So(ctx.entries[key("", "inside with")], ShouldNotBeNil)
		So(ctx.entries[key("", "inside with else")], ShouldNotBeNil)
		So(ctx.entries[key("", "inside if without else")], ShouldNotBeNil)
		So(ctx.entries[key("", "inside range without else")], ShouldNotBeNil)
		So(ctx.entries[key("", "inside with without else")], ShouldNotBeNil)
		So(ctx.entries[key("", "piped through T")], ShouldNotBeNil)
		So(ctx.entries[key("", "piped through field T")], ShouldNotBeNil)
		So(ctx.entries[key("ctxt-pipe", "id-pipe")], ShouldNotBeNil)
		So(ctx.entries[key("", "wrapped string arg")], ShouldNotBeNil)
	})
}

// Test_run 验证整体入口 Run 能按测试模板完成一次提取流程。
func Test_run(t *testing.T) {
	Convey("run", t, func() {
		err := Run(&Param{
			Input:      "testdata/*.tmpl",
			Left:       "{{",
			Right:      "}}",
			Keyword:    "T;X:1c,2;N:1,2;XN:1c,2,3",
			Function:   "T",
			OutputFile: "-",
		})
		So(err, ShouldBeNil)
	})
}
