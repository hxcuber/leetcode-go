package solutions

import (
	"slices"
)

// Doesn't work
func maxDistanceFirstTry(position []int, m int) int {
	slices.Sort(position)

	maximum := position[len(position)-1]
	minimum := position[0]
	maxDiff := maximum - minimum

	if m == 2 {
		return maxDiff
	}

	maxMinDiff := maxDiff / (m - 1)
	next := maximum - maxMinDiff*(m-2)
	nextPos, _ := slices.BinarySearch(position, next)

	if position[nextPos]-next > next-position[nextPos-1] && nextPos != 1 {
		nextPos--
	}

	nextPos = min(len(position)-1-(m-2), nextPos)
	return min(position[nextPos]-minimum, maxDistanceFirstTry(position[nextPos:], m-1))
}
