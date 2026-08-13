package main

import "fmt"

func IsLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 {
		return true
	}

	return year%100 == 0 && year%400 == 0
}

func main() {
	fmt.Println(IsLeapYear(1997))
	fmt.Println(IsLeapYear(1900))
	fmt.Println(IsLeapYear(2000))
}
