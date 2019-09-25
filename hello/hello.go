package main

import (
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
	"github.com/madtypist/stringutil"
	"time"
)

// variables declared at the package level
var foo, bar string = "something", "more"

func areaOfCircle(r float64) int {
	//let's return a nice round number
	return int(math.Pi * math.Pow(r, 2))
}

func main() {
	// need to Seed rand to make it truly (somewhat) random
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Hello, world.")
	fmt.Println(stringutil.Reverse("Hello, Go!"))
	fmt.Println("My favorite number is", rand.Intn(20))

	// Printf fun
	fmt.Printf("I've got %g problems, but finding the area of a circle ain't one.\n", math.Pi)
	fmt.Println(areaOfCircle(5))

	// function level variables
	// These things don't need a type initializer - they will take on the type of 
	// whatever is assigned to them
	var x,y = 45.6, true

	// We can do var blocks as well
	var (
		z	bool = true
		maxInt	uint64  = 1<<64-1
		tricky  complex128 = cmplx.Sqrt(-5 + 12i)
	)


	// Inside a function, the := short assignment statement can be used in place of a 
	// var declaration with implicit type
	thing1, thing2 := 13, "bananas"

	fmt.Println(foo, bar, x, y, thing1, thing2, tricky)

	// You can use %T to figure out variable type
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}