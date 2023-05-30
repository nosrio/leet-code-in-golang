package longestpalindromicsubstring

import "testing"

func TestExample1(t *testing.T) {
	input := "babad"
	output := "bab"
	got := longestPalindrome(input)
	if got != output {
		t.Errorf("longestPalindrome(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample2(t *testing.T) {
	input := "cbbd"
	output := "bb"
	got := longestPalindrome(input)
	if got != output {
		t.Errorf("longestPalindrome(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample3(t *testing.T) {
	input := "forgeeksskeegfor"
	output := "geeksskeeg"
	got := longestPalindrome(input)
	if got != output {
		t.Errorf("longestPalindrome(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample4(t *testing.T) {
	input := "abb"
	output := "bb"
	got := longestPalindrome(input)
	if got != output {
		t.Errorf("longestPalindrome(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample5(t *testing.T) {
	input := "abcdbbfcba"
	output := "bb"
	got := longestPalindrome(input)
	if got != output {
		t.Errorf("longestPalindrome(%v)=%v, wanted: %v", input, got, output)
	}
}
