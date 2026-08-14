package main

import "fmt"

// Gauss
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// https://prepp.in/question/the-sum-of-the-square-of-the-first-ten-natural-num-6a1fb1a2f0b97f5cdf293b6e
func SumOfSquares(n int) int {
	return ((n * (n + 1)) * (2*n + 1)) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

func main() {
	fmt.Println(SquareOfSum(10))
	fmt.Println(SumOfSquares(10))
	fmt.Println(Difference(10))
}
