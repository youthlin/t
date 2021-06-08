package t_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/youthlin/t"
)

func TestLT(testT *testing.T) {
	type args struct {
		lang  string
		msgID string
		args  []interface{}
	}
	t.BindDefaultDomain("testdata/")
	tests := []struct {
		name string
		args args
		want string
	}{
		{"not-support-lang", args{"de_DE", t.Noop.T("Hello, World"), nil}, "Hello, World"},
		{"support-lang", args{"zh_CN", t.Noop.T("Hello, World"), nil}, "你好，世界"},
	}
	for _, tt := range tests {
		testT.Run(tt.name, func(testT *testing.T) {
			if got := t.LT(tt.args.lang, tt.args.msgID, tt.args.args...); got != tt.want {
				testT.Errorf("LT() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestLN(testT *testing.T) {
	domain := "app"
	t.BindTextDomain(domain, "testdata")
	t.TextDomain(domain)
	lang := "zh_CN"
	Convey("LN", testT, func() {
		So(t.LN(lang, "One apple", "%d apples", 1, 1), ShouldEqual, "1 个苹果")
		So(t.LN64(lang, "One apple", "%d apples", 1, 1), ShouldEqual, "1 个苹果")
	})
}

func TestLX(testT *testing.T) {
	t.BindDefaultDomain("testdata/")
	Convey("LX", testT, func() {
		So(t.LX("zh_CN", "File|", "Open"), ShouldEqual, "打开文件")
		So(t.LXN("zh_CN", "File|", "Open One", "Open %d", 1, 1), ShouldEqual, "打开 1 个文件")
		So(t.LXN64("zh_CN", "File|", "Open One", "Open %d", 1, 1), ShouldEqual, "打开 1 个文件")
	})
}
