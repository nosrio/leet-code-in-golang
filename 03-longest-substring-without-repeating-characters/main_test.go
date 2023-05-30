package lswrc

import "testing"

func TestExample1(t *testing.T) {
	input := "abcabcbb"
	output := 3
	got := lengthOfLongestSubstring(input)
	if got != output {
		t.Errorf("lengthOfLongestSubstring(%v) = %v; want %v", input, got, output)
	}
}

func TestExample2(t *testing.T) {
	input := "bbbbb"
	output := 1
	got := lengthOfLongestSubstring(input)
	if got != output {
		t.Errorf("lengthOfLongestSubstring(%v) = %v; want %v", input, got, output)
	}
}

func TestExample3(t *testing.T) {
	input := "pwwkew"
	output := 3
	got := lengthOfLongestSubstring(input)
	if got != output {
		t.Errorf("lengthOfLongestSubstring(%v) = %v; want %v", input, got, output)
	}
}

func TestExample4(t *testing.T) {
	input := ""
	output := 0
	got := lengthOfLongestSubstring(input)
	if got != output {
		t.Errorf("lengthOfLongestSubstring(%v) = %v; want %v", input, got, output)
	}
}

func TestExample5(t *testing.T) {
	input := " "
	output := 1
	got := lengthOfLongestSubstring(input)
	if got != output {
		t.Errorf("lengthOfLongestSubstring(%v) = %v; want %v", input, got, output)
	}
}

func TestExample6(t *testing.T) {
	input := "aa"
	output := 1
	got := lengthOfLongestSubstring(input)
	if got != output {
		t.Errorf("lengthOfLongestSubstring(%v) = %v; want %v", input, got, output)
	}
}

func TestExample7(t *testing.T) {
	input := "au"
	output := 2
	got := lengthOfLongestSubstring(input)
	if got != output {
		t.Errorf("lengthOfLongestSubstring(%v) = %v; want %v", input, got, output)
	}
}
