package leetcode

/**
 * 1422. Maximum Score After Splitting a String
 * https://leetcode.com/problems/maximum-score-after-splitting-a-string/?envType=daily-question&envId=2025-01-01
 */

import (
	"fmt"
	"math"
)

func MaxScore(s string) int {
	n := len(s)
	totalOnes := 0
	for _, c := range s {
		totalOnes += int(c - '0')
	}

	maxVal := 0.0
	ones := 0
	for i := n - 1; i >= 0; i-- {
		ones += int(s[i] - '0')
		maxVal = math.Max(maxVal, float64(ones+i-totalOnes+ones))
	}

	fmt.Printf("n: %f", maxVal)
	fmt.Printf("\n totalOnes: %d", totalOnes)

	return int(maxVal)
}
