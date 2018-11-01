package utils

import "sort"

// IDS is ID slice.
type IDS []ID

// Empty returns true if ID slice is empty, else false.
func (s *IDS) Empty() bool {
	return len(*s) == 0
}

// Add push ids to ID slice.
func (s *IDS) Add(ids ...ID) []ID {
	*s = append(*s, ids...)
	return *s
}

// IndexOf returns the index of the matched id, else -1.
func (s *IDS) IndexOf(id ID) int {
	for f, v := range *s {
		if id == v {
			return f
		}
	}
	return -1
}

// Contains returns true if id is found, else false.
func (s *IDS) Contains(id ID) bool {
	for _, v := range *s {
		if id == v {
			return true
		}
	}
	return false
}

// Delete removes all occurrences of ids in the ID slice.
func (s *IDS) Delete(ids ...ID) []ID {
	for _, v := range ids {
		for {
			f := s.IndexOf(v)
			if f < 0 {
				break
			}
			*s = append((*s)[:f], (*s)[f+1:]...)
		}
	}
	return *s
}

// DeleteIndex removes all occurrences of index in the ID slice.
func (s *IDS) DeleteIndex(index ...int) []ID {
	size := len(index)
	if size == 0 {
		return *s
	}
	if size > 1 {
		sort.Sort(sort.Reverse(sort.IntSlice(index)))
	}
	for _, i := range index {
		*s = append((*s)[:i], (*s)[i+1:]...)
	}
	return *s
}

// Intersect returns the slice of intersecting id.
// Union see Add
// Except see Delete
func (s *IDS) Intersect(ids ...ID) []ID {
	b := IDS(ids)
	del := IDS{}
	for _, i := range *s {
		if !b.Contains(i) {
			del.Add(i)
		}
	}
	s.Delete(del...)
	return *s
}

// Reverse returns the reverse order for ID slice.
func (s *IDS) Reverse() []ID {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
	return *s
}

// Distinct returns the new duplicate free slice.
func (s *IDS) Distinct() []ID {
	m := map[ID]bool{}
	index := []int{}
	for i, id := range *s {
		if _, has := m[id]; has {
			index = append(index, i)
		} else {
			m[id] = true
		}
	}
	s.DeleteIndex(index...)
	return *s
}
