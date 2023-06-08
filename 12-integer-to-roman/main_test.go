package integertoroman

import "testing"

func TestExample1(t *testing.T) {
	input := 3
	output := "III"
	got := intToRoman(input)

	if got != output {
		t.Errorf("intToRoman(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample2(t *testing.T) {
	input := 58
	output := "LVIII"
	got := intToRoman(input)

	if got != output {
		t.Errorf("intToRoman(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample3(t *testing.T) {
	input := 1994
	output := "MCMXCIV"
	got := intToRoman(input)

	if got != output {
		t.Errorf("intToRoman(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample4(t *testing.T) {
	input := 1999
	output := "MCMXCIX"
	got := intToRoman(input)

	if got != output {
		t.Errorf("intToRoman(%v)=%v, wanted: %v", input, got, output)
	}
}
