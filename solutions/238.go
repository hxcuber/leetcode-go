package solutions

func productExceptSelf(nums []int) []int {
	size := len(nums)
	forward := make([]int, size)
	backward := make([]int, size)

	forward[0] = nums[0]
	backward[size-1] = nums[size-1]
	for i := 1; i < size; i++ {
		forward[i] = forward[i-1] * nums[i]
		backward[size-i-1] = backward[size-i] * nums[size-i-1]
	}

	answer := make([]int, size)

	answer[0] = backward[1]
	answer[size-1] = forward[size-2]
	for i := 1; i < size-1; i++ {
		answer[i] = forward[i-1] * backward[i+1]
	}

	return answer
}
