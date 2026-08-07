package main

import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	kindIsCar := kind == "car"
	kindIsTruck := kind == "truck"
	return kindIsCar || kindIsTruck
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	betterChoice := min(option1, option2)
	betterChoiceMsg := fmt.Sprintf("%s is clearly the better choice.", betterChoice)
	return betterChoiceMsg
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	var realPrice float64
	if age < 3 {
		realPrice = originalPrice * 0.8
	} else if age >= 3 && age < 10 {
		realPrice = originalPrice * 0.7
	} else {
		realPrice = originalPrice * 0.5
	}

	return realPrice
}

func main() {
	fmt.Println(NeedsLicense("car"))
	fmt.Println(NeedsLicense("bike"))
	fmt.Println(NeedsLicense("truck"))
	fmt.Println(ChooseVehicle("Wuling Hongguang", "Toyota Corolla"))
	fmt.Println(ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf"))
	fmt.Println(CalculateResellPrice(1000, 1))
	fmt.Println(CalculateResellPrice(1000, 5))
	fmt.Println(CalculateResellPrice(1000, 15))
}
