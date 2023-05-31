package zigzagconversion

func convert(s string, numRows int) string {
	var str string
	if numRows == 1 {
		return s
	}
	step := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += step {
			str += string(s[j])
			if i != 0 && i != numRows-1 && (j+step-2*i) < len(s) {
				str += string(s[j+step-2*i])
			}
		}
	}
	return str
}
