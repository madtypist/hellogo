package main

import (
	"fmt"
	"math"
	"math/rand"
	"github.com/madtypist/stringutil"
	"time"
)

// variables declared at the package level
var foo, bar string = "something", "more"

func areaOfCircle(r float64) float64 {
	return math.Pi * math.Pow(r, 2)	
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

	fmt.Println(foo, bar, x, y)
}