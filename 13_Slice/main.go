package main

import (
	"fmt"
)

func main() {
	s := []string{
		"MikeyRobby",
		"Todd",
		"Go",
		"Fresno",
		"City",
		"College",
	}

	for i, value := range s {
		fmt.Println(i, " - ", value)
	}

	mySlice := []int{2,4,6,8,10,9,7,5,3,1,}

	for i, value := range mySlice {
		fmt.Println(i, value)
	}
}
