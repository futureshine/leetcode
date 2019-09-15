package implement

func ThreeSum(nums []int) [][]int {
	return threeSum(nums)
}

func quickSort(nums *[]int, left, right int) {
	if left >= right {
		return
	}
	var realNums = *nums
	if realNums[left] < realNums[right] {
		tmp := realNums[left]
		realNums[left] = realNums[right]
		realNums[right] = tmp
	}
	var pivot = realNums[right]
	var indexLeft, indexRight = left, right - 1
	for {
		if indexLeft >= indexRight {
			break
		}
		if realNums[indexLeft] >= pivot && realNums[indexRight] <= pivot {
			var tmp = realNums[indexLeft]
			realNums[indexLeft] = realNums[indexRight]
			realNums[indexRight] = tmp
			indexLeft++
			indexRight--
		} else if realNums[indexLeft] < pivot {
			indexLeft++
		} else if realNums[indexRight] > pivot {
			indexRight--
		}
	}
	if realNums[indexLeft] >= pivot {
		realNums[right] = realNums[indexLeft]
		realNums[indexLeft] = pivot
		quickSort(nums, left, indexLeft-1)
		quickSort(nums, indexLeft+1, right)
	} else {
		realNums[right] = realNums[indexLeft+1]
		realNums[indexLeft+1] = pivot
		quickSort(nums, left, indexLeft)
		quickSort(nums, indexLeft+2, right)
	}

}

func threeSum(nums []int) [][]int {
	var length = len(nums)
	quickSort(&nums, 0, length-1)
	var result = [][]int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		} else {
			var left, right = i + 1, length - 1
			var flag = false
			var mark = 0
			for {
				var tmp = []int{}
				if left >= right || (i >= 1 && nums[i] == nums[i-1]) {
					break
				}
				sum := nums[i] + nums[left] + nums[right]
				if sum == 0 {
					if (!flag) || (flag && nums[left] != mark) {
						tmp = append(tmp, nums[i], nums[left], nums[right])
						flag = true
						mark = nums[left]
						result = append(result, tmp)
					}
					left++
					right--
				} else if sum < 0 {
					left++
				} else {
					right--
				}
			}
		}

	}
	return result
}
