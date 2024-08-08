package solutions

import "leetcode-go"

func numberOfCombinations(nums string) int {
	if nums == "" || nums[0] == '0' || len(nums) == 1 {
		return 0
	}

	prev := numberOfCombinations(nums[1:])

	if nums[0] <= nums[1] {
		prev++
	} else {
		prev--
	}

	return (1 + prev) % leetcode_go.MOD
}
