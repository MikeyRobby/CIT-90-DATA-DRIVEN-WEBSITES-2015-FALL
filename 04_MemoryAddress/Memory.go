package main

import "fmt"

func main() {
	a := "MikeyRobby"

	fmt.Println(a) // MikeyRobby
	fmt.Println(&a) // Memory Address of variable a

	var b *string = &a // b points to the memory address of variable a, The * means pointer
	fmt.Println(b) // prints the memory address of variable a that is now stored in variable b
	fmt.Println(*b) // Print out the value of b which is MikeyRobby.
}

