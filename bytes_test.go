package utils

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDecode(t *testing.T) {
	Convey("Decode", t, func() {
		var str string
		err := Decode([]byte{5, 12, 0, 2, 'a', 'b'}, &str)
		So(err, ShouldEqual, nil)
		So(str, ShouldEqual, "ab")
		Convey("Error", func() {
			var nstr string
			err := Decode(nil, &nstr)
			So(err.Error(), ShouldEqual, "EOF")
			So(nstr, ShouldEqual, "")
		})
	})
}

func TestEncode(t *testing.T) {
	Convey("Encode", t, func() {
		bs, _ := Encode("ab")
		So(bs, ShouldResemble, []byte{5, 12, 0, 2, 'a', 'b'})
	})
}
