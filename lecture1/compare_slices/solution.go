package compare_slices

func Compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[int]int, len(a))
	for _, val := range a {
		m[val]++
	}
	for _, val := range b {
		m[val]--
	}

	for _, key := range m {
		if key != 0 {
			return false
		}
	}
	return true
}
