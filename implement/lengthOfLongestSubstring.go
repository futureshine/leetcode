package implement

func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring(s)
}

func larger(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLongestSubstring(s string) int {
	var length = len(s)
	var (
		result = 0
		i      = 0
		j      = 0
	)
	var tmp = map[byte]int{}
	for {
		if i < length && j < length {
			if tmp[s[j]] != 1 {
				tmp[s[j]] = 1
				j++
				result = larger(result, j-i)
			} else {
				delete(tmp, s[i])
				i++
			}
		} else {
			break
		}
	}
	return result
}
