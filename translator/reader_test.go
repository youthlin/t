package translator

import (
	"io"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/youthlin/t/errors"
)

func Test_newReader(t *testing.T) {
	Convey("newReader", t, func() {
		Convey("nil-input", func() {
			r := newReader(nil)
			So(r, ShouldNotBeNil)
			So(r.lines, ShouldEqual, []string(nil))
			So(r.lineNo, ShouldEqual, -1)
			So(r.totalLine, ShouldEqual, 0)
		})
		Convey("empty-input", func() {
			r := newReader([]string{})
			So(r, ShouldNotBeNil)
			So(r.lines, ShouldResemble, []string{})
			So(r.lineNo, ShouldEqual, -1)
			So(r.totalLine, ShouldEqual, 0)
		})
		Convey("one-line", func() {
			r := newReader([]string{"hello"})
			So(r, ShouldNotBeNil)
			So(r.lines, ShouldResemble, []string{"hello"})
			So(r.lineNo, ShouldEqual, -1)
			So(r.totalLine, ShouldEqual, 1)
		})
	})
}

func Test_reader_currentLine(t *testing.T) {
	Convey("currentLine", t, func() {
		Convey("new-and-currentLine", func() {
			r := newReader([]string{"hello"})
			_, err := r.currentLine()
			So(err, ShouldNotBeNil)
		})
		Convey("next-and-currentLine", func() {
			r := newReader([]string{"hello"})
			line, err := r.nextLine()
			So(err, ShouldBeNil)
			So(line, ShouldEqual, "hello")
			str, err := r.currentLine()
			So(err, ShouldBeNil)
			So(str, ShouldEqual, line)
		})
	})
}
func Test_reader_nextLine(t *testing.T) {
	Convey("nextLine", t, func() {
		Convey("nil", func() {
			r := newReader(nil)
			_, err := r.nextLine()
			So(err, ShouldNotBeNil)
			So(errors.Is(err, io.EOF), ShouldBeTrue)
		})
		Convey("not-empty", func() {
			r := newReader([]string{"hello"})
			line, err := r.nextLine()
			So(err, ShouldBeNil)
			So(line, ShouldEqual, "hello")
			line, err = r.nextLine()
			So(line, ShouldEqual, "")
			So(errors.Is(err, io.EOF), ShouldBeTrue)
		})
	})
}
func Test_reader_unGetLine(t *testing.T) {
	Convey("unget", t, func() {
		Convey("nil", func() {
			r := newReader(nil)
			err := r.unGetLine()
			So(err, ShouldNotBeNil)
		})
		Convey("non-empty", func() {
			r := newReader([]string{"hello"})
			line, err := r.nextLine()
			So(err, ShouldBeNil)
			err = r.unGetLine()
			So(err, ShouldBeNil)
			line2, err := r.nextLine()
			So(err, ShouldBeNil)
			So(line, ShouldEqual, line2)
		})
	})
}
