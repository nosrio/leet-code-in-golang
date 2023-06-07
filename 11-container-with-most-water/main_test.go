package containerwithmostwater

import "testing"

func TestExample1(t *testing.T) {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	output := 49
	got := maxArea(input)
	if got != output {
		t.Errorf("maxArea(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample2(t *testing.T) {
	input := []int{1, 1}
	output := 1
	got := maxArea(input)
	if got != output {
		t.Errorf("maxArea(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample3(t *testing.T) {
	input := []int{1, 2, 1}
	output := 2
	got := maxArea(input)
	if got != output {
		t.Errorf("maxArea(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample4(t *testing.T) {
	input := []int{1, 2, 3, 4}
	output := 4
	got := maxArea(input)
	if got != output {
		t.Errorf("maxArea(%v)=%v, wanted: %v", input, got, output)
	}
}

func TestExample5(t *testing.T) {
	input := []int{2, 3, 4, 5, 18, 17, 6}
	output := 17
	got := maxArea(input)
	if got != output {
		t.Errorf("maxArea(%v)=%v, wanted: %v", input, got, output)
	}
}
