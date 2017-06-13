package u32

import "sort"

type Count struct {
	Num uint32
	Sum uint8
}

type SetCount map[uint32]*Count

func (m SetCount) Add(sets ...*Set) {
	for _, set := range sets {
		for k := range *set {
			if mv, ok := m[k]; ok {
				mv.Sum += 1
			} else {
				m[k] = &Count{Num: k, Sum: 1}
			}
		}
	}
}

func (m SetCount) Count() []Count {
	ret := make([]Count, len(m))
	i := 0
	for _, v := range m {
		ret[i] = *v
		i++
	}
	sort.Slice(ret, func(i, j int) bool {
		if ret[i].Sum == ret[j].Sum {
			return ret[i].Num < ret[j].Num
		}
		return ret[j].Sum < ret[i].Sum
	})
	return ret
}
