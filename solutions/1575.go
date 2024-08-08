package solutions

import "leetcode-go"

func countRoutesRecursive(locations []int, start int, finish int, fuel int) int {
	if fuel < 0 {
		return 0
	}

	result := 0
	if start == finish {
		result++
	}

	for i, v := range locations {
		if start != i {
			cost := locations[start] - v
			result += countRoutesRecursive(locations, i, finish, fuel-max(cost, -cost))
		}
	}

	return result % leetcode_go.MOD
}

func countRoutes(locations []int, start int, finish int, fuel int) int {
	dp := make([][]int, len(locations))
	for i := range dp {
		dp[i] = make([]int, fuel+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var memo func(c int, rf int) int
	memo = func(c int, rf int) int {
		if rf < 0 {
			return 0
		}

		if dp[c][rf] != -1 {
			return dp[c][rf]
		}

		result := 0
		if c == finish {
			result++
		}

		for i, v := range locations {
			if c != i {
				cost := locations[c] - v
				result = (result + memo(i, rf-max(cost, -cost))) % leetcode_go.MOD
			}
		}

		dp[c][rf] = result
		return result
	}

	result := memo(start, fuel)
	return result
}
