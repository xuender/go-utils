package goutils

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCheckError(t *testing.T) {
	Convey("CheckError", t, func() {
		Convey("ok", func() {
			CheckError(nil)
		})
		Convey("error", func() {
			So(func() {
				CheckError(errors.New("error"))
			}, ShouldPanic)
		})
	})
}
