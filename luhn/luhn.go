package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Valid(id string) bool {
	idWithoutSpaces := strings.ReplaceAll(id, " ", "")
	if idWithoutSpaces == "0" {
		return false
	}
	doubleDigitsId := []byte(idWithoutSpaces)
	length := len(idWithoutSpaces) - 1
	sumDigits := 0
	for i := 0; i <= length; i++ {
		reversedIndex := length - i
		n, err := strconv.Atoi(string(idWithoutSpaces[reversedIndex]))

		if err != nil {
			return false
		}

		if i%2 == 1 {
			n *= 2
			if n > 9 {
				n -= 9
			}
			doubleDigitsId[reversedIndex] = byte(n)
		}
		sumDigits += n
	}

	return sumDigits%10 == 0 || sumDigits == 0
}

func main() {
	fmt.Println(Valid("4539 3195 0343 6467"))
	fmt.Println(Valid("0"))
	fmt.Println(Valid(" 0"))
	fmt.Println(Valid("055b 444 285"))
}
