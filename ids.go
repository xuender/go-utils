package utils

// IDS slice ID
type IDS []ID

// Empty returns true if IDS is empty, else false.
func (s *IDS) Empty() bool {
	return len(*s) == 0
}

// Add ids
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

// Delete removes all occurrences of ids in the slice.
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

// Intersect returns the slice of intersecting values.
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
