package slidingwindow

func checkInclusion(s1 string, s2 string) bool {
	target := make(map[uint8]int)
	for i := 0; i < len(s1); i++ {
		target[s1[i]]++
	}
	cur := make(map[uint8]int)
	i, j, checked := 0, 0, 0
	for j < len(s2) {
		c := s2[j]
		j++
		if _, ok := target[c]; ok {
			cur[c]++
			if cur[c] == target[c] {
				checked++
			}
		}
		for j-i >= len(s1) {
			if checked == len(target) {
				return true
			}
			d := s2[i]
			i++
			if _, ok := target[d]; ok {
				if cur[d] == target[d] {
					checked--
				}
				cur[d]--
			}
		}
	}
	return false
}
