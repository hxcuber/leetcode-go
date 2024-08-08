package solutions

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]] = append(m[nums[i]], i)
	}

	for _, v := range m {
		for i := 0; i < len(v)-1; i++ {
			if v[i]-v[i+1] <= k && v[i]-v[i+1] >= -k {
				return true
			}
		}
	}

	return false
}
