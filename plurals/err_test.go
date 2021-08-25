package plurals

import (
	"fmt"
	"testing"

	"github.com/cockroachdb/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_errorListener_addError(t *testing.T) {
	Convey("add-err", t, func() {
		var err = errors.Errorf("abc=%v", 1)
		e := new(errorListener)
		e.addError(fmt.Errorf("fmt-error"))
		e.addError(errors.New("errors-new"))
		e.addError(errors.Wrapf(err, "wrap message"))
		t.Logf("err=%+v", e.err)
	})
}
