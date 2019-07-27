package implement

func Reverse(x int) int {
	return reverse(x)
}

// 数字翻转
func reverse(x int) int {
	var result int
	for {
		if x == 0 {
			break
		}
		yushu := x % 10
		result = result*10 + yushu
		x = x / 10
	}
	if x >= -2<<31 && x <= 2<<31-1 {
		return result
	}
	return 0
}
