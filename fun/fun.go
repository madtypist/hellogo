package main

import (
	"fmt"
)

/*
	This function takes in two strings and returns them in swapped order
*/
func swap(x, y string)(string, string) {
	return y,x
}

func main() {
	// Let's have some fun!
	a, b := "hello", "world"
	fmt.Println("Original strings:", a, b)
	c, d := swap(a,b)
	fmt.Println("New strings:", c, d)
}