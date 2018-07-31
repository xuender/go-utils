package goutils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewFileId(t *testing.T) {
	Convey("FileId", t, func() {
		Convey("Write", func() {
			id, _ := NewFileId("")
			id.Write([]byte("xxx"))
			So(id.String(), ShouldEqual, "pouybESLWCKDbbx4xozlvwM=")
		})
		Convey("error", func() {
			_, err := NewFileId("/dfdfdfd")
			So(err, ShouldNotBeNil)
		})
		Convey("String", func() {
			id, err := NewFileId("LICENSE")
			So(err, ShouldBeNil)
			So(id.String(), ShouldEqual, "SEofGUXjne39NwsrA8NO8TEE")
		})
	})
}
