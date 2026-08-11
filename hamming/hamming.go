package main

import (
	"errors"
	"fmt"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("The strings have diferrent lengths.")
	}

	length := len(a)
	hammingDistance := 0
	for i := range length {
		if a[i] != b[i] {
			hammingDistance += 1
		}
	}

	return hammingDistance, nil
}

func main() {
	fmt.Println(Distance("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT"))
}
