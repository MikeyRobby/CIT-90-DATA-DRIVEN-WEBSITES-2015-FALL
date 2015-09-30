package main

import (
	"fmt"
)

func person(Fname string, age int)(string, int){
	fmt.Println("What is your name? ")
	fmt.Scan(&Fname)
	println("What is your age? ")
	fmt.Scan(&age)
	fmt.Println(Fname, "is ", age, " years old.")
	return Fname, age
}


func main(){
	person("", 0)


}








