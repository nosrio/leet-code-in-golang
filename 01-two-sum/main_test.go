package two_sum

import "testing"

func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestExample1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	output := []int{0, 1}
	got := twoSum(nums, target)
	if !testEq(got, output) {
		t.Errorf("twoSum(%v,%d) = %v; want %v", nums, target, got, output)
	}
}

func TestExample2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	output := []int{1, 2}
	got := twoSum(nums, target)
	if !testEq(got, output) {
		t.Errorf("twoSum(%v,%d) = %v; want %v", nums, target, got, output)
	}
}

func TestExample3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	output := []int{0, 1}
	got := twoSum(nums, target)
	if !testEq(got, output) {
		t.Errorf("twoSum(%v,%d) = %v; want %v", nums, target, got, output)
	}
}
