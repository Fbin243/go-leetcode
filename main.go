package main

import (
	"fmt"
	"go_leetcode/leetcode"
)

func main() {
	// result := medium.BestSightseeingPair([]int{8, 1, 5, 2, 6})
	// result := medium.LengthOfLIS_DP([]int{10, 9, 2, 5, 3, 7, 101, 18})
	// result := leetcode.MaxScore("011101")
	result := leetcode.VowelStrings([]string{"aba", "bcb", "ece", "aa", "e"}, [][]int{{0, 2}, {1, 4}, {1, 1}})

	fmt.Printf("\nResult: %d\n", result)
}
