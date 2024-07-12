package dp

import "math"

func Solution1(wgt, val []int, i, c int) int {
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

func Solution2(coins []int, i, a int) int {
	if a == 0 {
		return -1
	}
	dp := make([][]int, i+1)
	for j := 0; j <= i; j++ {
		dp[j] = make([]int, a+1)
	}
	for j := 1; j <= a; j++ {
		dp[0][j] = a + 1
	}
	for j := 1; j <= i; j++ {
		for k := 1; k <= a; k++ {
			if coins[j-1] > k {
				dp[j][k] = dp[j-1][k]
			} else {
				dp[j][k] = int(math.Min(float64(dp[j-1][k]), float64(dp[j][k-coins[j-1]]+1)))
			}
		}
	}
	if dp[i][a] == a+1 {
		return -1
	}
	return dp[i][a]
}

// 优化空间
func solution2(coins []int, i, a int) int {
	if a == 0 {
		return -1
	}
	dp := make([]int, a+1)
	for j := 1; j <= a; j++ {
		dp[j] = a + 1
	}
	for j := 1; j <= i; j++ {
		for k := 1; k <= a; k++ {
			if coins[j-1]-k > 0 {
				dp[k] = dp[k]
			} else {
				dp[k] = int(math.Min(float64(dp[k]), float64(dp[k-coins[j-1]]+1)))
			}
		}
	}
	if dp[a] == a+1 {
		return -1
	}
	return dp[a]
}

func Solution3(coins []int, i, a int) int {
	if a == 0 {
		return 1
	}
	dp := make([][]int, i+1)
	for j := 0; j <= i; j++ {
		dp[j] = make([]int, a+1)
		dp[j][0] = 1
	}
	for j := 1; j <= a; j++ {
		dp[0][j] = 0
	}
	for j := 1; j <= i; j++ {
		for k := 1; k <= a; k++ {
			if coins[j-1] > k {
				dp[j][k] = dp[j-1][k]
			} else {
				dp[j][k] = dp[j-1][k] + dp[j][k-coins[j-1]]
			}
		}
	}
	return dp[i][a]
}

func Solution4(s, t []string) int {
	sLen := len(s)
	tLen := len(t)
	dp := make([][]int, sLen+1)
	for i := 0; i <= sLen; i++ {
		dp[i] = make([]int, tLen+1)
		dp[i][0] = i
	}
	for j := 1; j <= tLen; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= sLen; i++ {
		for j := 1; j <= tLen; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Min(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1])), float64(dp[i-1][j-1]))) + 1
			}
		}
	}
	return dp[sLen][tLen]
}
