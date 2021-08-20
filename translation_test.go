package t

import (
	"embed"
	"io/fs"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//go:embed testdata
var fsys embed.FS

func TestTranslation_Load(t *testing.T) {
	Convey("Load", t, func() {
		var tr = NewTranslation("")
		tr.Load("testdata")
		So(tr.Langs, ShouldResemble, []string{"zh_CN"})
	})
	Convey("FS", t, func() {
		Convey("fs", func() {
			fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
				t.Logf("walk: path=%v d=%T err=%v", path, d, err)
				return err
			})
		})
		Convey("os", func() {
			fs.WalkDir(os.DirFS("testdata"), ".", func(path string, d fs.DirEntry, err error) error {
				t.Logf("walk: path=%v d=%T err=%v", path, d, err)
				return err
			})
		})
	})
	Convey("LoadFS", t, func() {
		var tr = NewTranslation("")
		tr.LoadFS(fsys)
		So(tr.Langs, ShouldResemble, []string{"zh_CN"})
	})
}
