package mo

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRead(t *testing.T) {
	Convey("Read", t, func() {
		f, err := os.Open("../../testdata/zh_CN.mo")
		So(err, ShouldBeNil)
		So(f, ShouldNotBeNil)
		mo, err := Read(f)
		So(err, ShouldBeNil)
		So(mo.Lang(), ShouldEqual, "zh_CN")
	})
}
