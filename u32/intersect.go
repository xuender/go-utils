package u32

// A and B is sorted
func IsIntersect(a, b []uint32) bool {
	aLen := len(a)
	bLen := len(b)
	if aLen == 0 || bLen == 0 {
		return false
	}
	aLast := a[aLen-1]
	bLast := b[bLen-1]
	if a[0] > bLast || b[0] > aLast {
		return false
	}
	for _, i := range a {
		if i < b[0] {
			continue
		}
		if i > bLast {
			break
		}
		for _, f := range b {
			if i == f {
				return true
			}
		}
	}
	return false
}
