package main

import (
	"fmt"
	"go_leetcode/medium"
)

func main() {
	// result := medium.BestSightseeingPair([]int{8, 1, 5, 2, 6})
	LIS := medium.LengthOfLIS_DP([]int{10, 9, 2, 5, 3, 7, 101, 18})
	fmt.Printf("Result: %d\n", LIS)
}
