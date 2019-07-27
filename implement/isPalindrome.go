package implement

func IsPalindrome(x int) bool {
	return isPalindrome(x)
}

//回文数判断
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	var result int
	for {
		if tmp == 0 {
			break
		}
		result = result*10 + tmp%10
		tmp /= 10
	}
	if result == x {

		return true
	}
	return false
}
