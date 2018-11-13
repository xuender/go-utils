package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewFileId(t *testing.T) {
	Convey("FileId", t, func() {
		Convey("Write", func() {
			id, _ := NewFileID("")
			id.Write([]byte("xxx"))
			So(id.String(), ShouldEqual, "bff7b2198029378f03")
		})
		Convey("error", func() {
			_, err := NewFileID("/dfdfdfd")
			So(err, ShouldNotBeNil)
		})
		Convey("String", func() {
			id, err := NewFileID("LICENSE")
			So(err, ShouldBeNil)
			So(id.String(), ShouldEqual, "a0a0d77cad2926d13104")
		})
	})
}
