package solutions

func kthDistinct(arr []string, k int) string {
	m := make(map[string]int)
	for _, v := range arr {
		m[v]++
	}

	d := 0
	for _, v := range arr {
		if m[v] == 1 {
			d++
			if d == k {
				return v
			}
		}
	}
	return ""
}
