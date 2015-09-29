package main

import "fmt"

// My solution

func main(){
	var n int
	fmt.Println("What number would you like to half?: ")
	fmt.Scan(&n)
	if (n%2 == 0){
		fmt.Println(n/2)
		fmt.Println(true)
	}else{
		fmt.Println(n/2)
		fmt.Println(false)
	}
	//main()

}


// This is the solution for the exercise by the instructor Todd Mcleod

/**
func half(n int) (int, bool) {
	return n/2, n%2 == 0
}

func main(){
	h, even := half(2)
	fmt.Println(h, even)
}
**/


