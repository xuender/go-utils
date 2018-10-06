package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func read(line string) {

}
func readBuf(bs []byte) {

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

func TestReadBuf(t *testing.T) {
	Convey("Read Buf", t, func() {
		Convey("error file", func() {
			So(ReadBuf("/aaaaa", readBuf), ShouldNotBeNil)
		})
		Convey("read file", func() {
			So(ReadBuf("LICENSE", readBuf), ShouldBeNil)
		})
	})
}
