package reverseinteger

import "testing"

func TestExample1(t *testing.T) {
	input := 123
	output := 321
	got := reverse(input)
	if got != output {
		t.Errorf("reverse(%v)=%v, watend: %v", input, got, output)
	}
}

func TestExample2(t *testing.T) {
	input := -123
	output := -321
	got := reverse(input)
	if got != output {
		t.Errorf("reverse(%v)=%v, watend: %v", input, got, output)
	}
}

func TestExample3(t *testing.T) {
	input := 120
	output := 21
	got := reverse(input)
	if got != output {
		t.Errorf("reverse(%v)=%v, watend: %v", input, got, output)
	}
}

func TestExample4(t *testing.T) {
	input := 1534236469
	output := 0
	got := reverse(input)
	if got != output {
		t.Errorf("reverse(%v)=%v, watend: %v", input, got, output)
	}
}
