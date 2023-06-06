package palindromenumber

import "testing"

func TestExample1(t *testing.T) {
	input := 121
	output := true
	got := isPalindrome(input)
	if got != output {
		t.Errorf("isPalindrome(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample2(t *testing.T) {
	input := -121
	output := false
	got := isPalindrome(input)
	if got != output {
		t.Errorf("isPalindrome(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample3(t *testing.T) {
	input := 10
	output := false
	got := isPalindrome(input)
	if got != output {
		t.Errorf("isPalindrome(%v)=%v, wanted: %v", input, got, output)
	}
}
