package main

import "fmt"

func main() {
	n := 456720
		fmt.Printf("%d\t\t %#x\t\t %b\t\t ", n, n, n)
	fmt.Println(&n)
}
