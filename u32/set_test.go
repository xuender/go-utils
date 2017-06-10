package u32

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSetAdd(t *testing.T) {
	Convey("Add", t, func() {
		set := NewSet(3, 6)
		So(set.Numbers(), ShouldResemble, []uint32{3, 6})
	})
}

func TestSetNumbers(t *testing.T) {
	Convey("Numbers", t, func() {
		Convey("[]uint32{3, 6}", func() {
			set := Set{}
			set.Add([]uint32{3, 6}...)
			So(set.Numbers(), ShouldResemble, []uint32{3, 6})
		})
		Convey("[]uint32{136, 99999, 630, 9}", func() {
			set := Set{}
			set.Add([]uint32{136, 630, 6}...)
			set.Add([]uint32{630, 6}...)
			set.Add([]uint32{99999, 9}...)
			So(set.Numbers(), ShouldResemble, []uint32{6, 9, 136, 630, 99999})
		})
	})
}

func TestSetClear(t *testing.T) {
	Convey("Clear", t, func() {
		set := NewSet(3, 4, 5)
		set.Clear()
		So(len(set.Numbers()), ShouldEqual, 0)
	})
}

func TestSetCopy(t *testing.T) {
	Convey("Copy", t, func() {
		old := NewSet(3)
		new := old.Copy()
		So(new.Numbers()[0], ShouldEqual, 3)
	})
}

func TestSetComplement(t *testing.T) {
	Convey("Complement", t, func() {
		a := NewSet(3, 4, 5)
		b := NewSet(4, 5, 6, 7)
		a.Complement(b)
		So(a.Numbers(), ShouldResemble, []uint32{6, 7})
	})
}

func TestSetEmpty(t *testing.T) {
	Convey("Empty", t, func() {
		a := NewSet()
		So(a.Empty(), ShouldEqual, true)
		a.Add(3)
		So(a.Empty(), ShouldEqual, false)
		b := NewSet(3)
		a.Minus(b)
		So(a.Empty(), ShouldEqual, true)
	})
}

func TestSetEqual(t *testing.T) {
	Convey("Equal", t, func() {
		a := NewSet([]uint32{3, 60, 9000}...)
		b := NewSet([]uint32{60, 3, 9000}...)
		So(a.Equal(b), ShouldEqual, true)
		b.Add(4)
		So(a.Equal(b), ShouldEqual, false)
	})
}

func TestSetUnion(t *testing.T) {
	Convey("Union", t, func() {
		set := NewSet(6)
		Convey("3", func() {
			new := NewSet(3)
			set.Union(new)
			So(set.Numbers(), ShouldResemble, []uint32{3, 6})
		})
	})
}

func TestSetMinus(t *testing.T) {
	Convey("Minus", t, func() {
		set := NewSet(1, 2, 1000, 90, 9000)
		new := NewSet(2, 1000, 80, 91, 9000)
		set.Minus(new)
		So(set.Numbers(), ShouldResemble, []uint32{1, 90})
	})
}

func TestSetRemove(t *testing.T) {
	Convey("Remove", t, func() {
		set := NewSet(1, 2, 1000, 90, 9000)
		set.Remove(2, 1000, 80, 91, 9000)
		So(set.Numbers(), ShouldResemble, []uint32{1, 90})
	})
}

func TestSetIntersect(t *testing.T) {
	Convey("Intersect", t, func() {
		set := NewSet(1, 2, 1000, 90, 9000)
		new := NewSet(2, 1000, 80, 91, 9000)
		set.Intersect(new)
		So(set.Numbers(), ShouldResemble, []uint32{2, 1000, 9000})
	})
}

func TestSetRetain(t *testing.T) {
	Convey("Retain", t, func() {
		set := NewSet(1, 2, 1000, 90, 9000)
		set.Retain(2, 1000, 80, 91, 9000)
		So(set.Numbers(), ShouldResemble, []uint32{2, 1000, 9000})
	})
}

func TestSetContain(t *testing.T) {
	Convey("Contain", t, func() {
		a := NewSet(3, 60, 9, 9000)
		b := NewSet(3, 60, 9000)
		So(a.Contain(b), ShouldEqual, true)
		So(b.Contain(a), ShouldEqual, false)
	})
}
