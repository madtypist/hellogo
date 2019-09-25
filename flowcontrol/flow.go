package main

import (
	"fmt"
)

func main() {
	// a simple for loop
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println("Sum + ", i, " =", sum)
		sum += i
	}

	fmt.Println("Final sum =", sum)
}