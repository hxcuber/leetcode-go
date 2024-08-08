package solutions

func maximumUniqueSubarray(nums []int) int {
	hashmap := make(map[int]int)
	result := 0
	best := 0
	l := 0

	for i := 0; i < len(nums); i++ {
		result += nums[i]
		hashmap[nums[i]]++

		for hashmap[nums[i]] > 1 {
			hashmap[nums[l]]--
			result -= nums[l]
			l++
		}
		best = max(best, result)
	}

	return best
}
