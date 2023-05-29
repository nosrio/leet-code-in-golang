package two_sum

func twoSum(nums []int, target int) []int {
	var i, j, num, summand int = 0, 0, 0, 0
	var outer, inner []int
	var found bool = false
	outer = make([]int, len(nums))
	inner = make([]int, len(nums)-1)
	outer = nums

	for i, num = range outer {
		inner = outer[i+1 : len(nums)]
		for j, summand = range inner {
			if num+summand == target {
				found = true
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}

	return []int{i, j + i + 1}
}
