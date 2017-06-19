package u32

import "sort"

// Int32Slice is utin32 sort slice.
type Int32Slice []uint32

func (p Int32Slice) Len() int           { return len(p) }
func (p Int32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort []uint32.
func (p Int32Slice) Sort() { sort.Sort(p) }
