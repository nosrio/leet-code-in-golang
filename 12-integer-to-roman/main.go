package integertoroman

func intToRoman(num int) string {
	var str string
	for x := num; x > 0; {
		switch {
		case x >= 1000:
			x -= 1000
			str += "M"
		case x >= 500:
			if x >= 900 {
				x -= 900
				str += "CM"
			} else {
				x -= 500
				str += "D"
			}
		case x >= 100:
			if x >= 400 {
				x -= 400
				str += "CD"
			} else {
				x -= 100
				str += "C"
			}
		case x >= 50:
			if x >= 90 {
				x -= 90
				str += "XC"
			} else {
				x -= 50
				str += "L"
			}
		case x >= 10:
			if x >= 40 {
				x -= 40
				str += "XL"
			} else {
				x -= 10
				str += "X"
			}
		case x >= 5:
			if x >= 9 {
				x -= 9
				str += "IX"
			} else {
				x -= 5
				str += "V"
			}
		case x >= 1:
			if x >= 4 {
				x -= 4
				str += "IV"
			} else {
				x -= 1
				str += "I"
			}
		}
	}
	return str
}
