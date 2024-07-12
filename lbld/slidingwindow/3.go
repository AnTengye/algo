package slidingwindow

import "math"

// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	if sLen == 0 {
		return 0
	}
	check := make(map[uint8]struct{}, sLen)
	i, j := 0, 0
	max := 0
	for i <= j && j < sLen {
		for _, ok := check[s[j]]; ok; _, ok = check[s[j]] {
			delete(check, s[i])
			i++
		}
		check[s[j]] = struct{}{}
		max = int(math.Max(float64(max), float64(j-i+1)))
		j++

	}
	return max
}
