package utils

// StringSlice string slice
type StringSlice []string

// Add adds all elements of vals in the slice.
func (ss *StringSlice) Add(vals ...string) []string {
	for _, s := range vals {
		*ss = append(*ss, s)
	}
	return *ss
}

// IndexOf returns the index of the matched val, else -1.
func (ss *StringSlice) IndexOf(val string) int {
	for i, v := range *ss {
		if v == val {
			return i
		}
	}
	return -1
}

// Contains returns true if val is found, else false.
func (ss StringSlice) Contains(val string) bool {
	for _, v := range ss {
		if v == val {
			return true
		}
	}
	return false
}

// Delete removes all occurrences of vals in the slice.
func (ss *StringSlice) Delete(vals ...string) []string {
	for _, v := range vals {
		for {
			i := ss.IndexOf(v)
			if i < 0 {
				break
			}
			*ss = append((*ss)[:i], (*ss)[i+1:]...)
		}
	}
	return *ss
}
