package dp

import "math"

func bl(wgt, val []int, i, c int) int {
	if i == 0 || c == 0 {
		return 0
	}
	if wgt[i-1] > c {
		return bl(wgt, val, i-1, c)
	}
	push := bl(wgt, val, i-1, c-wgt[i-1]) + val[i-1]
	notPush := bl(wgt, val, i-1, c)
	return int(math.Max(float64(push), float64(notPush)))
}

func bl1(wgt, val []int, i, c int) int {
	mem := [][]int{}
	for i := 0; i <= len(wgt); i++ {
		mem = append(mem, make([]int, c+1))
	}
	return bl1Mem(wgt, val, mem, i, c)
}

func bl1Mem(wgt, val []int, mem [][]int, i, c int) int {
	if i == 0 || c == 0 {
		return 0
	}
	if mem[i][c] != 0 {
		return mem[i][c]
	}
	if wgt[i-1] > c {
		return bl1Mem(wgt, val, mem, i-1, c)
	}
	push := bl(wgt, val, i-1, c-wgt[i-1]) + val[i-1]
	notPush := bl(wgt, val, i-1, c)
	mem[i][c] = int(math.Max(float64(push), float64(notPush)))
	return mem[i][c]
}

func Solution(wgt, val []int, i, c int) int {
	if i == 0 || c == 0 {
		return 0
	}
	dp := make([]int, c+1)
	for j := 1; j <= i; j++ {
		for currentC := c; currentC > 0; currentC-- {
			if wgt[j-1]-currentC <= 0 {
				dp[currentC] = int(math.Max(float64(dp[currentC]), float64(dp[currentC-wgt[j-1]]+val[j-1])))
			}
		}
	}
	return dp[c]
}
