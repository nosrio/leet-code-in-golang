package stringtointeger

import "testing"

func TestExample1(t *testing.T) {
	input := "42"
	output := 42
	got := myAtoi(input)
	if got != output {
		t.Errorf("myAtoi(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample2(t *testing.T) {
	input := "   -42"
	output := -42
	got := myAtoi(input)
	if got != output {
		t.Errorf("myAtoi(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample3(t *testing.T) {
	input := "4193 with words"
	output := 4193
	got := myAtoi(input)
	if got != output {
		t.Errorf("myAtoi(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample4(t *testing.T) {
	input := "words and 987"
	output := 0
	got := myAtoi(input)
	if got != output {
		t.Errorf("myAtoi(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample5(t *testing.T) {
	input := "-91283472332"
	output := -2147483648
	got := myAtoi(input)
	if got != output {
		t.Errorf("myAtoi(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample6(t *testing.T) {
	input := "21474836460"
	output := 2147483647
	got := myAtoi(input)
	if got != output {
		t.Errorf("myAtoi(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample7(t *testing.T) {
	input := "00000-42a1234"
	output := 0
	got := myAtoi(input)
	if got != output {
		t.Errorf("myAtoi(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample8(t *testing.T) {
	input := "   +0 123"
	output := 0
	got := myAtoi(input)
	if got != output {
		t.Errorf("myAtoi(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample9(t *testing.T) {
	input := "+1"
	output := 1
	got := myAtoi(input)
	if got != output {
		t.Errorf("myAtoi(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample10(t *testing.T) {
	input := "9223372036854775808"
	output := 2147483647
	got := myAtoi(input)
	if got != output {
		t.Errorf("myAtoi(%v)=%v, wanted: %v", input, got, output)
	}
}
