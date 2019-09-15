package implement

func LongestCommonPrefix(str []string) string {
	return longestCommonPrefix(str)
}

// 简单方法
func longestCommonPrefix(str []string) string {
	if len(str) == 0 {
		return ""
	}
	var index = 0
	for {
		flag := true
		var mark uint8
		if index < len(str[0]) {
			mark = str[0][index]
		} else {
			break
		}
		for _, i := range str {
			if !(index < len(i) && i[index] == mark) {
				flag = false
			}
		}
		if flag {
			index++
		} else {
			break
		}
	}
	return str[0][0:index]
}
