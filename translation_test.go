package t

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTranslation_Load(t *testing.T) {
	Convey("Load", t, func() {
		var tr = NewTranslation("")
		tr.Load("testdata")
	})
}
