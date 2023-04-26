package main

import (
	"fmt"
)

// You can also declare multiple variables of different types:
var (
	x int
	y float32
	z string
)

func main() {
	//Declare a single variable and initialize
	var a int = 5
	// You can also declare multiple variables at once:
	var b, c, d = 5, 10, 15
	// You can also leave the type and initialize the variable
	// with its value(Using the shorthand notation with the:= )
	e, f, g := 20, 25, 30

	fmt.Println(a, b, c, d, e, f, g, x, y, z)
}
