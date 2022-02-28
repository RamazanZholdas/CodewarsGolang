package main

func Comp(a []int, b []int) bool {
	if a == nil || b == nil || len(a) != len(b) {
		return false
	}

	m := make(map[int]int, len(b))
	for _, number := range b {
		m[number]++
	}

	for _, number := range a {
		if count, ok := m[number*number]; !ok || count == 0 {
			return false
		} else {
			m[number*number]--
		}
	}

	return true
}
