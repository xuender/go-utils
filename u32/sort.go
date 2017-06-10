package u32

import "sort"

type Int32Slice []uint32

func (p Int32Slice) Len() int           { return len(p) }
func (p Int32Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int32Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Int32Slice) Sort()              { sort.Sort(p) }

type CountSlice []Count

func (p CountSlice) Len() int { return len(p) }
func (p CountSlice) Less(i, j int) bool {
	if p[i].Sum == p[j].Sum {
		return p[i].Num < p[j].Num
	}
	return p[j].Sum < p[i].Sum
}
func (p CountSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p CountSlice) Sort()         { sort.Sort(p) }
