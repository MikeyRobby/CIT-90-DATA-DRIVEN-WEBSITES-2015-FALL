package main

import "fmt"

func main(){
	m := make(map[string]string)
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m["MikeyRobby"])
}

func changeMe(x map[string]string){
	x["MikeyRobby"] = "Student"
	x["Todd"] = "Teacher"
	fmt.Println(x)
}
