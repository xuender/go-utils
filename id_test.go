package goutils

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestID(t *testing.T) {
	Convey("NewID", t, func() {
		id := NewID('O')
		Convey("ID", func() {
			So(len(id), ShouldEqual, 18)
		})
		Convey("String", func() {
			i := new(ID)
			err := i.Parse("I-Cekw67uyMpBGZLRP2HFVbe")
			So(err, ShouldBeNil)
			So(i[0], ShouldEqual, 'I')
		})
		Convey("Parse", func() {
			So(len(id.String()), ShouldEqual, 24)
		})
		Convey("JSON Marshal", func() {
			b, _ := json.Marshal(id)
			So(len(id), ShouldEqual, 18)
			So(len(b), ShouldEqual, 26)
		})
		Convey("JSON Unarshal", func() {
			i := new(ID)
			err := json.Unmarshal([]byte(`"I-Cekw67uyMpBGZLRP2HFVbe"`), i)
			So(err, ShouldBeNil)
			So(len(i), ShouldEqual, 18)
			So(i[0], ShouldEqual, 'I')
			So(i.String(), ShouldEqual, "I-Cekw67uyMpBGZLRP2HFVbe")
		})
	})
}
