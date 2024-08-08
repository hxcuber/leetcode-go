package solutions

func addRungs(rungs []int, dist int) int {
	withZero := append([]int{0}, rungs...)
	answer := 0

	for i := 0; i < len(withZero)-1; i++ {
		if withZero[i+1]-withZero[i] > dist {
			answer += (withZero[i+1] - withZero[i]) / dist
		}
	}
	return answer
}
