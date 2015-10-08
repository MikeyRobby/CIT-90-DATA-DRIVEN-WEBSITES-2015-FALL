package main

import "fmt"

// Use var notation to initialize a variable to zero value of type rune
// assign a rune to the variable
// Print the type of the variable

func main() {
	var r rune
	r = 30
	fmt.Printf("%T", r)

	// The reason the type is of int32 is because rune is an alias for int32
}
