package u32

import "sort"

// Count is count Uint32.
type Count struct {
	Num uint32
	Sum uint8
}

// SetCount set count.
type SetCount map[uint32]*Count

// Add is SetCount add Sets.
func (m SetCount) Add(sets ...*Set) {
	for _, set := range sets {
		for k := range *set {
			if mv, ok := m[k]; ok {
				mv.Sum++
			} else {
				m[k] = &Count{Num: k, Sum: 1}
			}
		}
	}
}

// Count is SetCount count.
func (m SetCount) Count(removeSets ...*Set) []Count {
	ret := make([]Count, 0)
	set := NewSet()
	set.Union(removeSets...)
	for _, v := range m {
		if !set.Has(v.Num) {
			ret = append(ret, *v)
		}
	}
	sort.Slice(ret, func(i, j int) bool {
		if ret[i].Sum == ret[j].Sum {
			return ret[i].Num < ret[j].Num
		}
		return ret[j].Sum < ret[i].Sum
	})
	return ret
}
