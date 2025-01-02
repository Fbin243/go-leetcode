package leetcode

/**
 * 2559. Count Vowel Strings in Ranges
 * https://leetcode.com/problems/count-vowel-strings-in-ranges/description/?envType=daily-question&envId=2025-01-02
 */

import "fmt"

func VowelStrings(words []string, queries [][]int) []int {
	prefix_sum := make([]int, len(words)+2)
	for i, word := range words {
		if word[0] == word[len(word)-1] && isVowel(word[0]) {
			prefix_sum[i+1] = prefix_sum[i] + 1
		} else {
			prefix_sum[i+1] = prefix_sum[i]
		}
	}

	fmt.Printf("prefix_sum: %v", prefix_sum)

	output := make([]int, len(queries))
	for i, query := range queries {
		output[i] = prefix_sum[query[1]+1] - prefix_sum[query[0]]
	}

	return output
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
