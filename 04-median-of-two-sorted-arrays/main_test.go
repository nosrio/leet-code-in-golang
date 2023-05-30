package medianoftwosortedarrays

import "testing"

func TestExample1(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	output := float64(2.0)
	got := findMedianSortedArrays(nums1, nums2)
	if got != output {
		t.Errorf("findMedianSortedArrays(%v,%v)=%v, wanted: %v", nums1, nums2, got, output)
	}
}

func TestExample2(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	output := float64(2.5)
	got := findMedianSortedArrays(nums1, nums2)
	if got != output {
		t.Errorf("findMedianSortedArrays(%v,%v)=%v, wanted: %v", nums1, nums2, got, output)
	}
}

func TestExample3(t *testing.T) {
	nums1 := []int{}
	nums2 := []int{1}
	output := float64(1)
	got := findMedianSortedArrays(nums1, nums2)
	if got != output {
		t.Errorf("findMedianSortedArrays(%v,%v)=%v, wanted: %v", nums1, nums2, got, output)
	}
}

func TestExample4(t *testing.T) {
	nums1 := []int{}
	nums2 := []int{1, 2}
	output := float64(1.5)
	got := findMedianSortedArrays(nums1, nums2)
	if got != output {
		t.Errorf("findMedianSortedArrays(%v,%v)=%v, wanted: %v", nums1, nums2, got, output)
	}
}

func TestExample5(t *testing.T) {
	nums1 := []int{1, 4, 5}
	nums2 := []int{3, 3}
	output := float64(3)
	got := findMedianSortedArrays(nums1, nums2)
	if got != output {
		t.Errorf("findMedianSortedArrays(%v,%v)=%v, wanted: %v", nums1, nums2, got, output)
	}
}

func TestExample6(t *testing.T) {
	nums1 := []int{1}
	nums2 := []int{1}
	output := float64(1)
	got := findMedianSortedArrays(nums1, nums2)
	if got != output {
		t.Errorf("findMedianSortedArrays(%v,%v)=%v, wanted: %v", nums1, nums2, got, output)
	}
}
