package u32

import (
	"reflect"
	"sort"
)

type Set map[uint32]bool

func NewSet(nums ...uint32) *Set {
	s := make(Set, len(nums))
	s.Add(nums...)
	return &s
}

func (a Set) Add(nums ...uint32) {
	for _, n := range nums {
		a[n] = true
	}
}

func (a *Set) Clear() {
	*a = map[uint32]bool{}
}

func (a *Set) Complement(full *Set) {
	r := a.Copy()
	a.Clear()
	av := *a
	for i := range *full {
		if _, ok := r[i]; !ok {
			av[i] = true
		}
	}
}

func (a Set) Copy() Set {
	r := map[uint32]bool{}
	for i := range a {
		r[i] = true
	}
	return r
}

func (a *Set) Count() int {
	return len(*a)
}

func (a *Set) Empty() bool {
	return len(*a) == 0
}

func (a Set) Has(nums ...uint32) bool {
	for _, i := range nums {
		if _, ok := a[i]; !ok {
			return false
		}
	}
	return true
}

func (a Set) Numbers() []uint32 {
	ret := make([]uint32, 0, len(a))
	for n := range a {
		ret = append(ret, n)
	}
	sort.Sort(Int32Slice(ret))
	return ret
}

func (a *Set) Equal(b *Set) bool {
	return reflect.DeepEqual(*a, *b)
}

func (a Set) Union(sets ...*Set) {
	for _, set := range sets {
		for i := range *set {
			a[i] = true
		}
	}
}

func (a Set) Minus(sets ...*Set) {
	for _, set := range sets {
		for i := range *set {
			delete(a, i)
		}
	}
}

func (a Set) Remove(nums ...uint32) {
	for _, i := range nums {
		delete(a, i)
	}
}

func (a Set) Intersect(sets ...*Set) {
	for _, set := range sets {
		sv := *set
		for i := range a {
			if _, ok := sv[i]; !ok {
				delete(a, i)
			}
		}
	}
}

func (a Set) Retain(nums ...uint32) {
	a.Intersect(NewSet(nums...))
}

func (a *Set) Contain(sets ...*Set) bool {
	av := *a
	for _, set := range sets {
		for i := range *set {
			if _, ok := av[i]; !ok {
				return false
			}
		}
	}
	return true
}
