package main

import (
	"fmt"
)

var Person struct {
	name string
	age int

}

func main(){
	name := Person.name
	age := Person.age
	data := []string{name}
	fmt.Println("What is your name?")
	fmt.Scan(&name)
	fmt.Println("What is your age? ")
	fmt.Scan(&age)
	fmt.Println(append(data,name))
}
