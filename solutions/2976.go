package solutions

import (
	"math"
)

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	var m [26][26]int64

	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if i != j {
				m[i][j] = math.MaxInt64
			} else {
				m[i][j] = 0
			}
		}
	}

	for i, v := range cost {
		m[original[i]-'a'][changed[i]-'a'] = min(int64(v), m[original[i]-'a'][changed[i]-'a'])
	}

	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				if max(m[i][k], m[k][j]) != math.MaxInt64 {
					m[i][j] = min(m[i][j], m[i][k]+m[k][j])
				}
			}
		}
	}

	result := int64(0)
	for i, si := range source {
		if m[si-'a'][target[i]-'a'] == math.MaxInt64 {
			return -1
		}
		result += m[si-'a'][target[i]-'a']
	}
	return result
}
