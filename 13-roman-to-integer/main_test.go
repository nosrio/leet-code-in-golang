package roman_to_integer

import "testing"

func TestExample1(t *testing.T) {
	str := "III"
	result := 3

	got := romanToInt(str)
	if got != result {
		t.Errorf("romanToInt(%s) = %d; want %d", str, got, result)
	}
}

func TestExample2(t *testing.T) {
	str := "LVIII"
	result := 58

	got := romanToInt(str)
	if got != result {
		t.Errorf("romanToInt(%s) = %d; want %d", str, got, result)
	}
}

func TestExample3(t *testing.T) {
	str := "MCMXCIV"
	result := 1994

	got := romanToInt(str)
	if got != result {
		t.Errorf("romanToInt(%s) = %d; want %d", str, got, result)
	}
}
