package top

func longestPalindrome(s string) string {
	maxLen := 0
	maxString := ""
	for i := 0; i < len(s); i++ {
		if i > len(s)-maxLen/2 {
			break
		} else {
			j := 0
			for ; j+i < len(s) && s[i] == s[j+i]; j++ {
			}
			j = j + i - 1
			k := 0
			for ; i-k >= 0 && (j+k) < len(s); k++ {
				if s[j+k] != s[i-k] {
					break
				}
			}
			if k != 0 {
				k--
			}
			tmpStr := s[i-k : j+k+1]
			if maxLen < len(tmpStr) {
				maxString = tmpStr
				maxLen = len(tmpStr)
			}
		}
	}
	return maxString
}
