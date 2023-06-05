package reverseinteger

func reverse(x int) int {
	var out int
	var sign int
	limit := 2147483648
	if x > 0 {
		sign = 1
	} else {
		sign = -1
	}

	for num := x * sign; num > 0; num = num / 10 {
		residuo := num % 10
		out = out*10 + residuo
	}
	if out > limit {
		return 0
	}
	return out * sign
}
