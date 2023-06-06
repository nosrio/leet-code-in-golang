package stringtointeger

func myAtoi(s string) int {
	var out, sign int = 0, 1
	var converting_string bool = false
	limit := 2147483648

	for _, char := range s {
		if converting_string == false && int(char) == ' ' {
			continue
		}

		if converting_string == false && char == '-' {
			sign = -1
			converting_string = true
			continue
		}
		if converting_string == false && char == '+' {
			converting_string = true
			continue
		}

		if int(char) < 48 || int(char) > 57 {
			break
		}
		converting_string = true
		out = out*10 + int(char) - 48
		if out >= limit {
			out = limit
			if sign == 1 {
				out -= 1
			}
			break
		}
	}

	return out * sign
}
