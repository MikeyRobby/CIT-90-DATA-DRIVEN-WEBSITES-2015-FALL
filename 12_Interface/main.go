package main

import "fmt"

func main() {
	var collection []interface{}
	x := 7
	y := "MikeyRobby"
	z := 12.21

	collection = append(collection, x)
	collection = append(collection, y)
	collection = append(collection, z)

	fmt.Println(collection)


}
