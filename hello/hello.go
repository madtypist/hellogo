package main

import (
	"fmt"
	"math/rand"
	"github.com/madtypist/stringutil"
	"time"
)

func main() {
	// need to Seed rand to make it truly (somewhat) random
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Hello, world.")
	fmt.Println(stringutil.Reverse("Hello, Go!"))
	fmt.Println("My favorite number is ", rand.Intn(20))
}