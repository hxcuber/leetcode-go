package solutions

func majorityElement(nums []int) int {
	m := make(map[int]int)
	maj := 0
	for _, n := range nums {
		m[n]++
		if m[n] > m[maj] {
			maj = n
		}
	}

	return maj
}
