package implement

func ThreeSumClosest(nums []int, target int) int {
	return threeSumClosest(nums, target)
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}

func threeSumClosest(nums []int, target int) int {
	quickSort(&nums, 0, len(nums)-1)
	var offset = 1<<32 - 1
	var result int
	for i := 0; i < len(nums)-2; i++ {
		var left = i + 1
		var right = len(nums) - 1
		for {
			if left >= right {
				break
			}
			var sum = nums[i] + nums[left] + nums[right]
			var tmpOffset = abs(sum - target)
			if tmpOffset == 0 {
				return sum
			}
			if tmpOffset < offset {
				offset = tmpOffset
				result = sum
			}
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
