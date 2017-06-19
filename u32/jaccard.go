package u32

// Jaccard (A, B) = |A intersect B| / |A union B| * 1000
func Jaccard(a, b []uint32) int {
	aLen := len(a)
	bLen := len(b)
	if aLen == 0 || bLen == 0 {
		return 0
	}
	tmpMap := make(map[uint32]bool, aLen+bLen)
	var aOrB, aAndB int
	for _, i := range a {
		if _, ok := tmpMap[i]; !ok {
			tmpMap[i] = true
			aOrB++
		}
	}
	for _, i := range b {
		v, ok := tmpMap[i]
		if ok {
			if v {
				aAndB++
			} else {
				continue
			}
		} else {
			aOrB++
		}
		tmpMap[i] = false
	}
	return aAndB * 1000 / aOrB
}
