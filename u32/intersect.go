package u32

// A and B is sorted
func IsIntersect(a, b *[]uint32) bool {
	aValue := *a
	bValue := *b
	aLen := len(aValue)
	bLen := len(bValue)
	if aLen == 0 || bLen == 0 {
		return false
	}
	aLast := aValue[aLen-1]
	bLast := bValue[bLen-1]
	if aValue[0] > bLast || bValue[0] > aLast {
		return false
	}
	for _, i := range aValue {
		if i < bValue[0] {
			continue
		}
		if i > bLast {
			break
		}
		for _, f := range bValue {
			if i == f {
				return true
			}
		}
	}
	return false
}
