package top

func letterCombinations(digits string) []string {
	phone := map[uint8][]string{
		50: {"a", "b", "c"},
		51: {"d", "e", "f"},
		52: {"g", "h", "i"},
		53: {"j", "k", "l"},
		54: {"m", "n", "o"},
		55: {"p", "q", "r", "s"},
		56: {"t", "u", "v"},
		57: {"w", "x", "y", "z"},
	}
	if len(digits) == 0 {
		return []string{}
	}
	totalLen := 1
	for i := range digits {
		totalLen *= len(phone[digits[i]])
	}
	result := make([]string, totalLen)
	for i := range result {
		last := i
		tmp := ""
		//比如说此时输入的23457
		//那么就有00000五位待确定
		//其中每种最大的情况位33334
		//依次处理每一种
		for j := 0; j < len(digits); j++ {
			c := digits[j]
			pos := last % len(phone[c])
			tmp += phone[c][pos]
			last = last / len(phone[c])
		}
		result[i] = tmp
	}
	return result
}
