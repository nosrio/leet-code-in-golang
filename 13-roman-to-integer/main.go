package roman_to_integer

func romanToInt(s string) int {
	var number int = 0
	var result int = 0
	for i, char := range s {
		switch char {
		case 'I':
			number = 1
			if i+1 < len(s) && (s[i+1] == 'V' || s[i+1] == 'X') {
				number = -1
			}
		case 'V':
			number = 5
		case 'X':
			number = 10
			if i+1 < len(s) && (s[i+1] == 'L' || s[i+1] == 'C') {
				number = -10
			}
		case 'L':
			number = 50
		case 'C':
			number = 100
			if i+1 < len(s) && (s[i+1] == 'D' || s[i+1] == 'M') {
				number = -100
			}
		case 'D':
			number = 500
		case 'M':
			number = 1000
		}
		result += number
	}
	return result
}
