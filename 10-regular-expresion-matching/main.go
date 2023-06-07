package regularexpresionmatching

func isMatch(s string, p string) bool {
	var lenY, lenX int = len(s) + 1, len(p) + 1

	dp := make([][]bool, lenY)
	for y := range dp {
		dp[y] = make([]bool, lenX)
	}

	dp[0][0] = true
	for x := 1; x < lenX; x++ {
		if p[x-1] == '*' {
			dp[0][x] = dp[0][x-2]
		}
	}
	for y := 1; y < lenY; y++ {
		for x := 1; x < lenX; x++ {
			switch p[x-1] {
			case '*':
				if p[x-2] == s[y-1] || p[x-2] == '.' {
					dp[y][x] = dp[y-1][x] || dp[y][x-1] || dp[y][x-2]
				} else {
					dp[y][x] = dp[y][x-2]
				}
			case '.', s[y-1]:
				dp[y][x] = dp[y-1][x-1]
			}
		}
	}
	return dp[len(s)][len(p)]
}
