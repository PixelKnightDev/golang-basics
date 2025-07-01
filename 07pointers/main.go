package main

import "fmt"

func main() {
	fmt.Println("pointers basics")

	// var ptr *int

	// fmt.Println("Value of ptr is: ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

}
