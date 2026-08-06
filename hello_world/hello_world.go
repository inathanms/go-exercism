// https://exercism.org/tracks/go/exercises/hello-world

package main

import "fmt"

// HelloWorld greets the world!
func HelloWorld() string {
	return "Hello, World!"
}

func main() {
	result := HelloWorld()
	fmt.Println(result)
}
