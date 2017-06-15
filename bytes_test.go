package goutils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDecode(t *testing.T) {
	Convey("Decode", t, func() {
		var str string
		Decode([]byte{5, 12, 0, 2, 'a', 'b'}, &str)
		So(str, ShouldEqual, "ab")
	})
}

func TestEncode(t *testing.T) {
	Convey("Encode", t, func() {
		So(Encode("ab"), ShouldResemble, []byte{5, 12, 0, 2, 'a', 'b'})
	})
}
