package main

import "fmt"

// Use fmt.SprintF to string print to a variable the number 394 in hex
// Then use fmt.Println to print that variable to standard output


func main(){
	n := fmt.Sprintf("%#x", 394)
	fmt.Println(n)
}
