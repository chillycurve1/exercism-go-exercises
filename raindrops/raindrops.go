package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	var result string
	div3 := number % 3
	div5 := number % 5
	div7 := number % 7

	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		return strconv.Itoa(number)
	}

	if div3 == 0 {
		result = result + "Pling"
	}
	if div5 == 0 {
		result = result + "Plang"
	}
	if div7 == 0 {
		result = result + "Plong"
	}
	return result
}
