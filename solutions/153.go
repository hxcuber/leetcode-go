package solutions

//func findMin(nums []int) int {
//	for i := 0; i < len(nums)-1; i++ {
//		if nums[i] > nums[i+1] {
//			return nums[i+1]
//		}
//	}
//	return nums[0]
//}

func findMin(nums []int) int {
	if nums[0] < nums[len(nums)-1] {
		return nums[0]
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nums[0]
}
