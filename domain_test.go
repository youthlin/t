package t_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/youthlin/t"
)

func TestDT(testT *testing.T) {
	Convey("dgettext", testT, func() {
		domain := "test"
		t.BindTextDomain(domain, "testdata")
		t.SetUserLang("zh_CN")
		ctxt, msgID, plural := t.Noop.XN("Project|", "Open One", "Open %d")
		So(t.DT(domain, "Hello, World"), ShouldEqual, "你好，世界")
		So(t.DN64(domain, "One apple", "%d apples", 1, 1), ShouldEqual, "1 个苹果")
		So(t.DXN64(domain, ctxt, msgID, plural, 2, 2), ShouldEqual, "打开 2 个工程")
	})
}
