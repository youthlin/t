package t_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/youthlin/t"
)

func TestBindDefaultDomain(tt *testing.T) {
	Convey("Bind", tt, func() {
		Convey("Default", func() {
			t.BindDefaultDomain("testdata/")
			So(t.SupportLangs(t.DefaultDomain), ShouldResemble, []string{"zh_CN"})
		})
		Convey("domain", func() {
			t.BindTextDomain("my-domain", "testdata/")
			So(t.TextDomain("my-domain"), ShouldEqual, "my-domain")
			So(t.TextDomain("aaa"), ShouldEqual, t.DefaultDomain)
			So(t.SupportLangs("my-domain"), ShouldResemble, []string{"zh_CN"})
		})
	})
}
func TestUserLang(tt *testing.T) {
	Convey("UserLang", tt, func() {
		// So(t.UserLang(), ShouldEqual, "")
		t.SetUserLang("zh_CN")
		So(t.UserLang(), ShouldEqual, "zh_CN")
	})
}
func TestGettext(tt *testing.T) {
	Convey("gettext", tt, func() {
		testCase := []struct {
			msg  string
			want string
		}{
			{t.Noop.T("Hello, World"), "你好，世界"},
			{t.Noop.T("Hello, %s"), "你好，%s"},
		}
		t.BindDefaultDomain("testdata")
		t.SetUserLang("zh_CN")
		for _, tCase := range testCase {
			So(t.T(tCase.msg), ShouldEqual, tCase.want)
		}
		So(t.T("Hello, %s", "Tom"), ShouldEqual, "你好，Tom")
	})
}

func TestN(testT *testing.T) {
	type args struct {
		msgID       string
		msgIDPlural string
		n           int
		args        []interface{}
	}
	s, p := t.Noop.N("One apple", "%d apples")
	tests := []struct {
		name  string
		args  args
		want  string
		want2 string
	}{
		// TODO: Add test cases.
		{"1", args{s, p, 1, nil}, "One apple", "1 个苹果"},
		{"2", args{s, p, 2, nil}, "%d apples", "2 个苹果"},
	}
	// for _, tt := range tests {
	// 	testT.Run(tt.name, func(testT *testing.T) {
	// 		if got := t.N(tt.args.msgID, tt.args.msgIDPlural, tt.args.n, tt.args.args...); got != tt.want {
	// 			testT.Errorf("N() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
	t.BindDefaultDomain("testdata")
	t.SetUserLang("zh_CN")
	for _, tt := range tests {
		testT.Run(tt.name, func(testT *testing.T) {
			if got := t.N(tt.args.msgID, tt.args.msgIDPlural, tt.args.n, tt.args.n); got != tt.want2 {
				testT.Errorf("N() = %v, want %v", got, tt.want2)
			}
		})
	}
	Convey("N64", testT, func() {
		So(t.N64("One apple", "%d apples", 2, 200), ShouldEqual, "200 个苹果")
	})
}

func TestX(testT *testing.T) {
	Convey("pgettext", testT, func() {
		// So(t.X("File|", "Open"), ShouldEqual, "Open")
		// So(t.X("Project|", "Open"), ShouldEqual, "Open")
		// So(t.XN("File|", "Open One", "Open %d", 1), ShouldEqual, "Open One")
		// So(t.XN("Project|", "Open One", "Open %d", 2), ShouldEqual, "Open %d")
		// So(t.XN("Project|", "Open One", "Open %d", 2, 2), ShouldEqual, "Open 2")
		t.BindDefaultDomain("testdata/zh_CN.po")
		t.SetUserLang("zh_CN")
		ctxt, msgID := t.Noop.X("File|", "Open")
		So(t.X(ctxt, msgID), ShouldEqual, "打开文件")
		So(t.X("Project|", "Open"), ShouldEqual, "打开工程")
		So(t.XN("File|", "Open One", "Open %d", 1), ShouldEqual, "打开 %d 个文件")
		So(t.XN("File|", "Open One", "Open %d", 1, 1), ShouldEqual, "打开 1 个文件")
		So(t.XN("Project|", "Open One", "Open %d", 2), ShouldEqual, "打开 %d 个工程")
		So(t.XN("Project|", "Open One", "Open %d", 2, 2), ShouldEqual, "打开 2 个工程")
		So(t.XN64("Project|", "Open One", "Open %d", 2, 2), ShouldEqual, "打开 2 个工程")
	})
}
