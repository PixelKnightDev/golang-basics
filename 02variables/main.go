package main

import "fmt"

// var jwtToken int = 34444. ok

// jwttttoken := 3444. not allowed
// := walrus token only allowed inside a function no outside

const LoginToken string = "login checked" // starting with capital first letter means variable is public and can be used anywhere

func main() {
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.455332424234
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of the type: %T \n", anotherVariable)

	// implicit type

	var website = "lakers in 5"
	fmt.Println(website)
	fmt.Printf("Variable is of the type: %T \n", website)

	// no var style

	numberoOfUser := 3000000.9
	fmt.Println(numberoOfUser)

	// public

	fmt.Println(LoginToken)
	fmt.Printf("Variable type is: %T", LoginToken)
}
