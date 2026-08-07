package main

import "fmt"

// Package weather defines the set of functions and variables.
// package weather

var (
	// CurrentCondition: Variable to define current condition.
	CurrentCondition string
	// CurrentLocation: Variable to define current location.
	CurrentLocation string
)

// Forecast function.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

func main() {
	fmt.Println()
}
