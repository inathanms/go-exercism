package main

import "fmt"

func PreparationTime(layers []string, preparationTimePerLayerInMinutes int) int {
	if preparationTimePerLayerInMinutes == 0 {
		preparationTimePerLayerInMinutes = 2
	}

	return len(layers) * preparationTimePerLayerInMinutes
}

func Quantities(layers []string) (int, float64) {
	quantityOfNoodles := 0
	quantityOfSauce := float64(0)

	for i := range len(layers) {
		switch layers[i] {
		case "sauce":
			quantityOfSauce += 0.2
		case "noodles":
			quantityOfNoodles += 50
		}
	}

	return quantityOfNoodles, quantityOfSauce
}

func AddSecretIngredient(friendsList, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secretIngredient
}

func ScaleRecipe(amountsNeededForTwoPortions []float64, portionsToCook int) []float64 {
	amountsNeededForNPortions := make([]float64, len(amountsNeededForTwoPortions))
	scalePortions := float64(portionsToCook) / float64(2)
	for i := range len(amountsNeededForTwoPortions) {
		fmt.Println(amountsNeededForTwoPortions[i], scalePortions)
		amountsNeededForNPortions[i] = amountsNeededForTwoPortions[i] * scalePortions
	}

	return amountsNeededForNPortions
}

func main() {
	layers := []string{"sauce", "noodles", "sauce", "meat", "mozzarella", "noodles"}
	fmt.Println(PreparationTime(layers, 3))
	fmt.Println(PreparationTime(layers, 0))
	fmt.Println(Quantities(layers))
	friendsList := []string{"noodles", "sauce", "mozzarella", "kampot pepper"}
	myList := []string{"noodles", "meat", "sauce", "mozzarella", "?"}
	AddSecretIngredient(friendsList, myList)
	fmt.Println(myList)
	quantities := []float64{0.6, 300, 1, 0.5, 50, 0.1, 100}
	fmt.Println(ScaleRecipe(quantities, 3))
}
