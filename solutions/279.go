package solutions

import "math"

func numSquaresRecursive(n int) int {
	if n == 0 {
		return 0
	}

	minimum := math.MaxInt
	for i := 1; i*i <= n; i++ {
		result := 1 + numSquaresRecursive(n-i*i)
		if result < minimum {
			minimum = result
		}
	}

	return minimum
}

func numSquaresDP(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt
	}

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			result := 1 + dp[i-j*j]
			if result < dp[i] {
				dp[i] = result
			}
		}
	}

	return dp[n]
}
