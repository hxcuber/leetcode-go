package solutions

import "slices"

func frequencySortUnoptimised(nums []int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	type pair struct {
		val   int
		count int
	}

	var l []pair
	for k, v := range m {
		l = append(l, pair{k, v})
	}

	slices.SortFunc(l, func(a, b pair) int {
		if a.count == b.count {
			return b.val - a.val
		}
		return a.count - b.count
	})

	var result []int
	for _, v := range l {
		for range v.count {
			result = append(result, v.val)
		}
	}

	return result
}

func frequencySort(nums []int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}

	slices.SortFunc(nums, func(a, b int) int {
		if m[a] == m[b] {
			return b - a
		}
		return m[a] - m[b]
	})

	return nums
}
