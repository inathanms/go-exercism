package main

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	workingCarsPerHour := float64(productionRate) * (successRate / 100)
	return workingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerMinute := int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
	return workingCarsPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	carsOutiseAGroup := carsCount % 10
	groupsOfTenCars := (carsCount - carsOutiseAGroup) / 10
	cost := uint((carsOutiseAGroup * 10000) + (groupsOfTenCars * 95000))
	return cost
}

func main() {
	fmt.Println(CalculateWorkingCarsPerHour(1547, 90))
	fmt.Println(CalculateWorkingCarsPerMinute(1105, 90))
	fmt.Println(CalculateCost(37))
	fmt.Println(CalculateCost(21))
}
