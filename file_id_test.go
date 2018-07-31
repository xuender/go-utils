package goutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewFileId(t *testing.T) {
	Convey("NewFileId", t, func() {
		Convey("String", func() {
			id := NewFileId()
			id.Write([]byte("xxx"))
			So(id.String(), ShouldEqual, "pouybESLWCKDbbx4xozlvwM=")
		})
	})
	Convey("NewFileIdByFile", t, func() {
		Convey("error", func() {
			_, err := NewFileIdByFile("/dfdfdfd")
			So(err, ShouldNotBeNil)
		})
		Convey("String", func() {
			id, err := NewFileIdByFile("LICENSE")
			So(err, ShouldBeNil)
			So(id.String(), ShouldEqual, "SEofGUXjne39NwsrA8NO8TEE")
		})
	})
}
