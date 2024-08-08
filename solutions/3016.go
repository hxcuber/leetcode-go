package solutions

import (
	"slices"
)

func minimumPushesFirstTry(word string) int {
	m := make([]int, 26)
	for _, v := range word {
		m[v-'a']++
	}

	type Pair struct {
		letter int32
		count  int
	}

	var l []Pair
	for i, v := range m {
		if v != 0 {
			l = append(l, Pair{int32(i), v})
		}
	}

	slices.SortFunc(l, func(a, b Pair) int {
		return b.count - a.count
	})

	count := 0
	result := 0
	for i := 0; i < len(l); i++ {
		result += l[i].count * (count/8 + 1)
		count++
	}

	return result
}

func minimumPushes(word string) int {
	m := make([]int, 26)
	for _, v := range word {
		m[v-'a']++
	}

	slices.Sort(m)

	count := 0
	result := 0
	for i := 0; i < len(m); i++ {
		result += m[len(m)-i-1] * (count/8 + 1)
		count++
	}

	return result
}
