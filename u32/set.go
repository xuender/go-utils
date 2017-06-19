package u32

import (
	"reflect"
	"sort"
)

// Set is uint32 set.
type Set map[uint32]bool

// Add is add uint32 to set.
func (a Set) Add(nums ...uint32) {
	for _, n := range nums {
		a[n] = true
	}
}

// Clear set.
func (a *Set) Clear() {
	*a = map[uint32]bool{}
}

// Complement set.
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

// Contain set.
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

// Copy set.
func (a Set) Copy() Set {
	r := map[uint32]bool{}
	for i := range a {
		r[i] = true
	}
	return r
}

// Empty set.
func (a *Set) Empty() bool {
	return len(*a) == 0
}

// Equal set.
func (a *Set) Equal(b *Set) bool {
	return reflect.DeepEqual(*a, *b)
}

// Has is nums one or more in set.
func (a Set) Has(nums ...uint32) bool {
	for _, i := range nums {
		if _, ok := a[i]; ok {
			return true
		}
	}
	return false
}

// HasAll is all in set.
func (a Set) HasAll(nums ...uint32) bool {
	for _, i := range nums {
		if _, ok := a[i]; !ok {
			return false
		}
	}
	return true
}

// Hit nums.
func (a Set) Hit(nums ...uint32) int {
	ret := 0
	for _, i := range nums {
		if _, ok := a[i]; ok {
			ret++
		}
	}
	return ret
}

// Intersect set.
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

// Jaccard (A, B) = |A intersect B| / |A union B| * 1000
func (a *Set) Jaccard(sets ...*Set) int {
	i := a.Copy()
	i.Intersect(sets...)
	u := a.Copy()
	u.Union(sets...)
	return i.Len() * 1000 / u.Len()
}

// Len set.
func (a *Set) Len() int {
	return len(*a)
}

// Minus sets.
func (a Set) Minus(sets ...*Set) {
	for _, set := range sets {
		for i := range *set {
			delete(a, i)
		}
	}
}

// Numbers set to []uint32.
func (a Set) Numbers() []uint32 {
	ret := make([]uint32, 0, len(a))
	for n := range a {
		ret = append(ret, n)
	}
	sort.Sort(Int32Slice(ret))
	return ret
}

// Remove nums.
func (a Set) Remove(nums ...uint32) *Set {
	for _, i := range nums {
		delete(a, i)
	}
	return &a
}

// Retain nums.
func (a Set) Retain(nums ...uint32) {
	a.Intersect(NewSet(nums...))
}

// Union set.
func (a Set) Union(sets ...*Set) {
	for _, set := range sets {
		for i := range *set {
			a[i] = true
		}
	}
}

// NewSet is new Set by nums.
func NewSet(nums ...uint32) *Set {
	s := make(Set, len(nums))
	s.Add(nums...)
	return &s
}
