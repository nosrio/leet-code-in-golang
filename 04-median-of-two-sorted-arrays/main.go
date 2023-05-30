package medianoftwosortedarrays

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j, n int = 0, 0, 0
	var totalSize int = len(nums1) + len(nums2)
	var halfway int = totalSize/2 + 1
	var last, secondLast = math.MinInt, math.MinInt

	for n < halfway && i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			n++
			val := nums1[i]
			i++
			secondLast = last
			last = val
		} else {
			n++
			val := nums2[j]
			j++
			secondLast = last
			last = val
		}
	}
	for n < halfway && i < len(nums1) {
		n++
		val := nums1[i]
		i++
		secondLast = last
		last = val
	}
	for n < halfway && j < len(nums2) {
		n++
		val := nums2[j]
		j++
		secondLast = last
		last = val
	}
	if totalSize%2 == 0 {
		return float64(last+secondLast) / 2
	}

	return float64(last)
}
