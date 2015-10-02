package main

import "fmt"


func person(Fname string, age int) (int, bool){
	fmt.Println("What is your name?")
	fmt.Scan(&Fname)
	fmt.Println("How old are you?")
	fmt.Scan(&age)
	dogYear := age * 7
	if (age > 25){
		fmt.Println(Fname, " is ", dogYear, " in dog years and is old")
		return dogYear,true
	}else if (age < 25){
		fmt.Println(Fname, " is ", dogYear, " in dog years and is not old")
		return dogYear,false
	}
	return dogYear, true

}


func main(){
	person("", 0)
}
