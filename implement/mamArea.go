package implement

// 计算盛最多水的容器
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	var area = 0
	for i := 0; i < len(height)-1; i++ {
		for j := i; j < len(height); j++ {
			var tmpArea int
			if height[i] < height[j] {
				tmpArea = height[i] * (j - i)
			} else {
				tmpArea = height[j] * (j - i)
			}
			if tmpArea > area {
				area = tmpArea
			}
		}
	}
	return area
}

// 计算盛最多水的容器，一次扫描，从两侧开始向中间扫，每次移动较短的板子
func maxAreaOn(height []int) int {
	if len(height) < 2 {
		return 0
	}
	var area = 0
	var i = 0
	var j = len(height) - 1
	for {
		if j <= i {
			break
		}
		var tmpArea int
		if height[j] > height[i] {
			tmpArea = height[i] * (j - i)
			i++
		} else {
			tmpArea = height[j] * (j - i)
			j--
		}
		if tmpArea > area {
			area = tmpArea
		}
	}
	return area
}
