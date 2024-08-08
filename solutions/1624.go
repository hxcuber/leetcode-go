package solutions

func maxLengthBetweenEqualCharacters(s string) int {
	type pair struct {
		start int
		end   int
	}
	m := make(map[int32]int)

	result := -1
	for i, v := range s {
		if _, ok := m[v-'a']; !ok {
			m[v-'a'] = i
		} else {
			result = max(result, i-m[v-'a']-1)
		}
	}
	return result
}
