package longestpalindromicsubstring

func isPalindrome(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	var result string
	n := len(s)
	found := false
	for j := n; j > 0 && !found; j-- {
		for i := 0; i+j-1 < n && i < j+i && !found; i++ {
			if isPalindrome(s[i : j+i]) {
				found = true
				result = s[i : j+i]
			}
		}
	}

	return result
}
