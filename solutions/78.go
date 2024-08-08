package solutions

func subsets(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return [][]int{nil}
	}

	results := make([][]int, 0)
	ss := subsets(nums[1:])
	for _, s := range ss {
		sCopy := make([]int, len(s), len(s)+1)
		copy(sCopy, s)
		sCopy = append(sCopy, nums[0])
		results = append(results, sCopy)
	}
	results = append(results, ss...)
	return results
}
