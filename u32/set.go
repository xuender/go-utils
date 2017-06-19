package u32

import (
	"reflect"
	"sort"
)

// Set uint32集合.
type Set map[uint32]bool

// Add 增加uint32.
func (a Set) Add(nums ...uint32) {
	for _, n := range nums {
		a[n] = true
	}
}

// Clear 清空集合.
func (a *Set) Clear() {
	*a = map[uint32]bool{}
}

// Complement 补集.
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

// Contain 全包含..
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

// Copy 复制.
func (a Set) Copy() Set {
	r := map[uint32]bool{}
	for i := range a {
		r[i] = true
	}
	return r
}

// Empty 是否为空.
func (a *Set) Empty() bool {
	return len(*a) == 0
}

// Equal 集合比较.
func (a *Set) Equal(b *Set) bool {
	return reflect.DeepEqual(*a, *b)
}

// Has 是否包含任意一个uint32..
func (a Set) Has(nums ...uint32) bool {
	for _, i := range nums {
		if _, ok := a[i]; ok {
			return true
		}
	}
	return false
}

// HasAll 是否包含全部uint32.
func (a Set) HasAll(nums ...uint32) bool {
	for _, i := range nums {
		if _, ok := a[i]; !ok {
			return false
		}
	}
	return true
}

// Hit 命中次数.
func (a Set) Hit(nums ...uint32) int {
	ret := 0
	for _, i := range nums {
		if _, ok := a[i]; ok {
			ret++
		}
	}
	return ret
}

// Intersect 交集.
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

// Jaccard 相似度 (A, B) = |A intersect B| / |A union B| * 1000
func (a *Set) Jaccard(sets ...*Set) int {
	i := a.Copy()
	i.Intersect(sets...)
	u := a.Copy()
	u.Union(sets...)
	return i.Len() * 1000 / u.Len()
}

// Len 长度.
func (a *Set) Len() int {
	return len(*a)
}

// Minus 减去.
func (a Set) Minus(sets ...*Set) {
	for _, set := range sets {
		for i := range *set {
			delete(a, i)
		}
	}
}

// Numbers 返回[]uint32.
func (a Set) Numbers() []uint32 {
	ret := make([]uint32, 0, len(a))
	for n := range a {
		ret = append(ret, n)
	}
	sort.Sort(Int32Slice(ret))
	return ret
}

// Remove 删除.
func (a Set) Remove(nums ...uint32) *Set {
	for _, i := range nums {
		delete(a, i)
	}
	return &a
}

// Retain 交集.
func (a Set) Retain(nums ...uint32) {
	a.Intersect(NewSet(nums...))
}

// Union 并集.
func (a Set) Union(sets ...*Set) {
	for _, set := range sets {
		for i := range *set {
			a[i] = true
		}
	}
}

// NewSet 新建uint32集合.
func NewSet(nums ...uint32) *Set {
	s := make(Set, len(nums))
	s.Add(nums...)
	return &s
}
