package main

import (
	"fmt"
)

func printer(str ...string){  // Creating a function to loop through strings
	for k, v := range str{
		fmt.Println(k,v)
	}
}

func main() {

	mySlice := []string{"MikeyRobby","Todd","Daniel", "Rio", } // A slice of names
	fmt.Println(mySlice) // Prints out the slice

	printer(mySlice...) // Loops through mySlice and prints out the index and the value
}
