package main

import "fmt"

func main(){
	var name string
	fmt.Println("What is your name: ")
	fmt.Scan(&name)


		switch name {
		case "Ricky":
			fmt.Println("Hello, Ricky")
		case "Mikey":
			fmt.Println("Yo! Mikey!!")
		case "Joanna":
			fmt.Println("GO AWAY!!")
		case "Marcus":
			fmt.Println("Hey buddy!")
		default:
			fmt.Println("You are no one...")
		}

}
