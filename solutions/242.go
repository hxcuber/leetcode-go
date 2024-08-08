package solutions

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	m := map[rune]int{}
	for _, v := range s {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}

	for _, v := range t {
		if _, ok := m[v]; ok {
			m[v]--
			if m[v] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
