package implement

var resut [][]bool

// IsMatch 正则匹配
func IsMatch(s string, p string) bool {
	return isMatchDynamicBottomToTop(s, p)
}

// 回溯的方式
func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}
	var firstMatch = s != "" && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	}
	return firstMatch && isMatch(s[1:], p[1:])

}

// 动态规划自顶而下
func isMatchDynamicUpToBottom(s string, p string) bool {
	var result = make([]int, (len(s)+1)*(len(p)+1))
	return returnDynamic(0, 0, s, p, result)
}

func returnDynamic(i int, j int, text string, pattern string, result []int) bool {
	var ans int
	if result[i*(len(pattern)+1)+j] != 0 {
		ans = result[i*(len(pattern)+1)+j]
	} else {
		if j == len(pattern) {
			if i == len(text) {
				ans = 1
			} else {
				ans = -1
			}
		} else {
			firstMatch := i < len(text) && (text[i] == pattern[j] || pattern[j] == '.')
			if j+1 < len(pattern) && pattern[j+1] == '*' {
				if returnDynamic(i, j+2, text, pattern, result) || firstMatch && returnDynamic(i+1, j, text, pattern, result) {
					ans = 1
				} else {
					ans = -1
				}
			} else {
				if firstMatch && returnDynamic(i+1, j+1, text, pattern, result) {
					ans = 1
				} else {
					ans = -1
				}
			}
		}
		result[i*(len(pattern)+1)+j] = ans
	}
	if ans == 1 {
		return true
	}

	return false
}

//动态规划，自底而上
func isMatchDynamicBottomToTop(s string, p string) bool {
	var result = make([]int, (len(s)+1)*(len(p)+1))
	result[len(s)*(len(p)+1)+len(p)] = 1
	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			if result[i*(len(p)+1)+j] == 0 {
				var ans bool
				var firstMatch = i < len(s) && (s[i] == p[j] || p[j] == '.')
				if j+1 < len(p) && p[j+1] == '*' {
					ans = result[i*(len(p)+1)+j+2] == 1 || (firstMatch && result[(i+1)*(len(p)+1)+j] == 1)
				} else {
					ans = firstMatch && result[(i+1)*(len(p)+1)+j+1] == 1
				}
				if ans {
					result[i*(len(p)+1)+j] = 1
				} else {
					result[i*(len(p)+1)+j] = -1
				}
			}
		}
	}
	if result[0] == 1 {
		return true
	}
	return false
}
