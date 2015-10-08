package main

import "fmt"

//Use fmt.PrintF to print the type of a variable

func main(){
	n := 394
	s := "MikeyRobby"

	fmt.Printf("%T\n%T\n",n,n,) //Type int
	fmt.Printf("%T\n%T", s,s) // Type string
}
