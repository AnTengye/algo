package slidingwindow

import "math"

func minWindow(s string, t string) string {
	target := make(map[uint8]int)
	for i := 0; i < len(t); i++ {
		target[t[i]]++
	}
	cur := make(map[uint8]int)
	var ans string
	i, j, checked := 0, 0, 0
	min := math.MaxInt
	for j < len(s) {
		c := s[j]
		j++
		if _, ok := target[c]; ok {
			cur[c]++
			if cur[c] == target[c] {
				checked++
			}
		}
		for checked == len(target) {
			if j-i < min {
				ans = s[i:j]
				min = j - i
			}
			d := s[i]
			i++
			if _, ok := target[d]; ok {
				if cur[d] == target[d] {
					checked--
				}
				cur[d]--
			}
		}
	}
	return ans
}
