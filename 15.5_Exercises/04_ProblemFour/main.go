package main

import "fmt"

// Create a variable of type int initialized to a zero value
// Use fmt.Scan to receive input from the user @ the terminal
// Store the input in the variable of type int which you created

func main(){
	var a int
	fmt.Println("What is your age?")
	fmt.Scan(&a)
	fmt.Println("You are ", a, "years old!")
}