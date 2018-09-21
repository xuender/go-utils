package goutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewFileId(t *testing.T) {
	Convey("FileId", t, func() {
		Convey("Write", func() {
			id, _ := NewFileID("")
			id.Write([]byte("xxx"))
			So(id.String(), ShouldEqual, "a68bb26c448b5822836dbc78c68ce5bf03")
		})
		Convey("error", func() {
			_, err := NewFileID("/dfdfdfd")
			So(err, ShouldNotBeNil)
		})
		Convey("String", func() {
			id, err := NewFileID("LICENSE")
			So(err, ShouldBeNil)
			So(id.String(), ShouldEqual, "484a1f1945e39dedfd370b2b03c34ef13104")
		})
	})
}
