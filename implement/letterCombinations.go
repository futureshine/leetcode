package implement

func LetterCombinations(digits string) []string {
	return letterCombinations(digits)
}

func letterCombinations(digits string) []string {
	var numberMap = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	var result = []string{}
	if digits == "" {
		return []string{}
	}
	getCombinations("", digits, &result, numberMap)
	return result
}

func getCombinations(combinations, letters string, result *[]string, numberMap map[byte]string) {
	if len(letters) != 0 {
		var nextDigit = byte(letters[0])
		var letterString = numberMap[nextDigit]
		for _, i := range letterString {
			getCombinations(combinations+string(i), letters[1:], result, numberMap)
		}

	} else {
		*result = append(*result, combinations)
	}
}
