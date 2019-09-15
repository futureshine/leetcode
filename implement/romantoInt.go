package implement

func RomanToInt(s string) int {
	return romanToInt(s)
}

func romanToInt(s string) int {
	var intList = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	var stringList = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var numIndex = 0
	var startIndex = 0
	var sum = 0
	for {
		if startIndex == len(s) {
			break
		}
		if startIndex+len(stringList[numIndex]) <= len(s) && stringList[numIndex] == s[startIndex:startIndex+len(stringList[numIndex])] {
			sum += intList[numIndex]
			startIndex += len(stringList[numIndex])
		} else {
			numIndex++
		}
	}
	return sum
}
