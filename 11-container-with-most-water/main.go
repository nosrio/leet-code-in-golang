package containerwithmostwater

func maxArea(height []int) int {
	var area int = 0
	var n int = len(height)
	var h int

	for i, j := 0, n-1; i < j; {
		h = height[j]
		if height[i] < height[j] {
			h = height[i]
		}
		b := j - i
		area_t := b * h
		if area_t > area {
			area = area_t
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return area
}
