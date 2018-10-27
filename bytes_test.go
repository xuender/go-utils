package utils

import (
	"fmt"
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

func ExampleEncode() {
	bs, _ := Encode("ab")
	fmt.Printf("%x\n", bs)

	// Output:
	// 050c00026162
}

func ExampleDecode() {
	var str string
	Decode([]byte{5, 12, 0, 2, 'a', 'b'}, &str)
	fmt.Println(str)

	// Output:
	// ab
}

func ExamplePrefixBytes() {
	fmt.Println(PrefixBytes([]byte("abc"), 'i', 'd', '-'))

	// Output:
	// [105 100 45 97 98 99]
}

func ExampleJoin() {
	a := []byte("aa")
	b := []byte("bb")
	fmt.Println(Join([]byte(","), a, b))

	// Output:
	// [97 97 44 98 98]
}

func ExampleConcat() {
	a := []byte("aa")
	b := []byte("bb")
	fmt.Println(Concat(a, b))

	// Output:
	// [97 97 98 98]
}

func ExamplePrefixUint32() {
	fmt.Println(PrefixUint32([]byte("id-"), 3))

	// Output:
	// [105 100 45 3 0 0 0]
}

func Example() {
	fmt.Println(PrefixUint32([]byte("id-"), 3))

	// Output:
	// [105 100 45 3 0 0 0]
}
