package main

import (
	"errors"
	"fmt"
)

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Please, enter a positive number.")
	}

	steps := 0

	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}

		steps += 1
	}

	return steps, nil
}

func main() {
	fmt.Println(CollatzConjecture(12))
}
