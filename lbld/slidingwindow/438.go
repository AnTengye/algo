package slidingwindow

func findAnagrams(s string, p string) []int {
	ans := []int{}
	target := make(map[uint8]int)
	for i := 0; i < len(p); i++ {
		target[p[i]]++
	}
	cur := make(map[uint8]int)
	i, j, checked := 0, 0, 0
	for j < len(s) {
		c := s[j]
		j++
		if _, ok := target[c]; ok {
			cur[c]++
			if cur[c] == target[c] {
				checked++
			}
		}
		for j-i >= len(p) {
			if checked == len(target) {
				ans = append(ans, i)
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
