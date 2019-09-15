package implement

func IntToRoman(num int) string {
	return intToRoman(num)
}

func pow(index int) int {
	var result = 1
	for i := 0; i < index; i++ {
		result *= 10
	}
	return result
}

func joinStringByCounts(raw string, count int) string {
	var result = ""
	for i := 0; i < count; i++ {
		result += raw
	}
	return result
}

func intToRoman(num int) string {
	var result = ""
	var mapString = make(map[int]string)
	mapString[1] = "I"
	mapString[5] = "V"
	mapString[10] = "X"
	mapString[50] = "L"
	mapString[100] = "C"
	mapString[500] = "D"
	mapString[1000] = "M"
	mapString[0] = ""
	var numberIndex = [4]int{}
	var highIndex = 1000
	for i := 0; i < 4; i++ {
		numberIndex[i] = (num / highIndex) % 10
		highIndex /= 10
	}
	for i := 0; i < 4; i++ {
		if numberIndex[i] != 0 {
			currentIndex := 1000 / pow(i)
			if numberIndex[i] < 4 {
				result += joinStringByCounts(mapString[currentIndex], numberIndex[i])
			} else if numberIndex[i] == 4 {
				result += mapString[currentIndex] + mapString[5*currentIndex]
			} else if numberIndex[i] > 4 && numberIndex[i] < 9 {
				result += mapString[5*currentIndex] + joinStringByCounts(mapString[currentIndex], numberIndex[i]-5)
			} else {
				result += mapString[currentIndex] + mapString[10*currentIndex]
			}
		}
	}
	return result
}
