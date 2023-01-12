package t

import (
	"io/fs"
	"os"
	"testing"
	"path/filepath"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_asFS(t *testing.T) {
	Convey("asFS", t, func() {
		Convey("dir", func() {
			fsys := asFS("testdata")
			dir, err := fsys.Open(".")
			So(err, ShouldBeNil)
			So(dir, ShouldNotBeNil)
			fi, err := dir.Stat()
			So(err, ShouldBeNil)
			So(fi.IsDir(), ShouldBeTrue)

			entry, err := fs.ReadDir(fsys, ".")
			So(err, ShouldBeNil)
			So(len(entry) > 0, ShouldBeTrue)

			fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
				t.Logf("dir as FS: path=%v, err=%v entry: name=%v, isDir=%v\n", path, err, d.Name(), d.IsDir())
				return err
			})
		})
		Convey("file", func() {
			fsys := asFS("testdata/zh_CN.po")
			file, err := fsys.Open(".")
			So(err, ShouldBeNil)
			So(file, ShouldNotBeNil)

			bytes, err := fs.ReadFile(fsys, "")
			So(err, ShouldBeNil)
			So(len(bytes) > 0, ShouldBeTrue)

			fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
				t.Logf("file as FS: path=%v, err=%v entry: name=%v, isDir=%v\n", path, err, d.Name(), d.IsDir())
				return err
			})
		})
	})

	// Join 会去除后面的点
	t.Logf("%v", filepath.Join("testdata/zh_CN.mo", "."))

	// os.DirFS Open 时，是直接用 / 连接的
	f := os.DirFS("testdata/zh_CN.mo")
	fs.WalkDir(f, ".", func(path string, d fs.DirEntry, err error) error {
		// path=. | d= <nil> | err=stat testdata/zh_CN.mo/.: not a directory
		t.Logf("path=%v | d= %v | err=%v", path, d, err)
		return err
	})
	// adFS Open 时，用的 Join
	f = asFS("testdata/zh_CN.mo")
	fs.WalkDir(f, ".", func(path string, d fs.DirEntry, err error) error {
		// path=. | d= zh_CN.mo | err=<nil>
		t.Logf("path=%v | d= %v | err=%v", path, d.Name(), err)
		return err
	})
}
