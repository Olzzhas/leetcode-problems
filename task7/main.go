package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(reverse(-32))
}

func getChar(str string, index int) rune {
	return []rune(str)[index]
}

func reverse(x int) int {
	var result int

	text := strconv.FormatInt(int64(x), 10)

	lenOfText := len(text)

	var reversedText string

	for i := lenOfText - 1; i > -1; i-- {
		if i == lenOfText-1 && string(getChar(text, lenOfText-1)) == "0" {
			continue
		}

		if string(getChar(text, i)) == "-" {
			continue
		}

		reversedText = fmt.Sprint(reversedText, string(getChar(text, i)))

	}

	fmt.Println(reversedText)

	result, _ = strconv.Atoi(reversedText)

	if x < 0 {
		result = result * -1
	}

	if result > 2147483647 || result < -2147483648 {
		return 0
	}

	return result
}
