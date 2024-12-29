package medium

import "math"

/**
 * 1014. Best Sightseeing Pair
 * https://leetcode.com/problems/best-sightseeing-pair/description/?envType=daily-question&envId=2024-12-27
 */

func BestSightseeingPair(A []int) int {
	max := float64(0)
	preMax := float64(A[0])
	for j := 1; j < len(A); j++ {
		max = math.Max(max, preMax+float64(A[j]-j))
		preMax = math.Max(preMax, float64(A[j]+j))
	}
	return int(max)
}
