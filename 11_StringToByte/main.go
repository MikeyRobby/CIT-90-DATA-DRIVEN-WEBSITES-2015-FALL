package main

import (
	"fmt"
)

func main() {
	s := "MikeyRobby"
	fmt.Println(string([]byte{'M','I', 'K', 'E', 'Y', 'R', 'O', 'B', 'B', 'Y',}))
	fmt.Println([]byte(s))
}

