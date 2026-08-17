package main

import "fmt"

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitsMap := map[string]int{}
	unitsMap["quarter_of_a_dozen"] = 3
	unitsMap["half_of_a_dozen"] = 6
	unitsMap["dozen"] = 12
	unitsMap["small_gross"] = 120
	unitsMap["gross"] = 144
	unitsMap["great_gross"] = 1728

	return unitsMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]

	if !exists {
		return false
	}

	bill[item] += value
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	unitValue, unitExists := units[unit]
	itemValue, itemExists := bill[item]

	if !itemExists || !unitExists {
		return false
	}

	newValue := itemValue - unitValue

	if newValue < 0 {
		return false
	}

	if newValue == 0 || itemValue == 0 || unitValue == 0 {
		delete(bill, item)
		return true
	}

	bill[item] = newValue

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
	if !exists {
		return 0, false
	}

	return value, exists
}

func main() {
	units := Units()
	fmt.Println(units)
	bill := NewBill()
	fmt.Println(bill)
	// ok := AddItem(bill, units, "carrot", "dozen")
	// fmt.Println(ok)
	fmt.Println(RemoveItem(bill, units, "carrot", "dozen"))
	bill2 := map[string]int{"carrot": 12, "grapes": 3}
	qty, ok := GetItem(bill2, "carrot")
	fmt.Println(qty)
	fmt.Println(ok)
}
