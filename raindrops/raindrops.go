package main

import (
	"fmt"
	"strconv"
)

func Convert(number int) string {
	raindropSound := ""
	isNotDivisibleBy3And5And7 := true
	if number%3 == 0 {
		raindropSound += "Pling"
		isNotDivisibleBy3And5And7 = false
	}
	if number%5 == 0 {
		raindropSound += "Plang"
		isNotDivisibleBy3And5And7 = false
	}
	if number%7 == 0 {
		raindropSound += "Plong"
		isNotDivisibleBy3And5And7 = false
	}

	if isNotDivisibleBy3And5And7 {
		return strconv.Itoa(number)
	}

	return raindropSound
}

func main() {
	fmt.Println(Convert(52))
	fmt.Println(Convert(14))
	fmt.Println(Convert(20))
	fmt.Println(Convert(105))
}
