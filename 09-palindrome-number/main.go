package palindromenumber

func isPalindrome(x int) bool {
	var reversed int = 0

	for n := x; n > 0; n /= 10 {
		residuo := n % 10
		reversed = reversed*10 + residuo
	}

	return x == reversed
}
