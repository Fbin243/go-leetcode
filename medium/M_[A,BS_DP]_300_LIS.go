package medium

/**
 * 300. Longest Increasing Subsequence
 * https://leetcode.com/problems/longest-increasing-subsequence/description/
 */

func LengthOfLIS_DP(nums []int) int {
	LIS := make([]int, len(nums))
	res := 0
	for i := len(nums) - 1; i >= 0; i-- {
		LIS[i] = 1
		for j := i; j < len(nums); j++ {
			if nums[i] < nums[j] {
				LIS[i] = max(LIS[i], LIS[j]+1)
			}
		}
		res = max(LIS[i], res)
	}

	return res
}