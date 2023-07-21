package reversenumber

import "strconv"

func ReverseNumber(num int) int {

	str := strconv.Itoa(num)

	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	reversed, _ := strconv.Atoi(reversedStr)

	return reversed
}
