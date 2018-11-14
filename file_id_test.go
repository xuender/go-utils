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
			So(id.String(), ShouldEqual, "a68d8af56c8b5822836dbc798864a9ff03")
		})
		Convey("error", func() {
			_, err := NewFileID("/dfdfdfd")
			So(err, ShouldNotBeNil)
		})
		Convey("String", func() {
			id, err := NewFileID("LICENSE")
			So(err, ShouldBeNil)
			So(id.String(), ShouldEqual, "d63d326f93123e0400d1766f34d379813104")
		})
	})
}
