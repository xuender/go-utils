package goutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func read(line string) {

}

func TestReadLines(t *testing.T) {
	Convey("Read Lines", t, func() {
		Convey("error file", func() {
			So(ReadLines("/aaaaa", read), ShouldNotBeNil)
		})
		Convey("read file", func() {
			So(ReadLines("LICENSE", read), ShouldBeNil)
		})
	})
}
