package main

import (
	"fmt"
)

func main() {
	// a simple for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println("Sum + ", i, " =", sum)
	}


	// a different loop leaving out explicit declaration of init and post statements
	sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	// same idea, but now done demonstrating that 'for' is 'while' in Go
	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
	}
}