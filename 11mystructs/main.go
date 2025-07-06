package main

import "fmt"

func main() {
	fmt.Println("structs")

	// no inheritance in golang: no super or parent

	hitesh := User{"Hitesh", "hiesh@gmail.com", true, 19}
	fmt.Println(hitesh)

	fmt.Printf("Hitesh deatils are: %+v\n", hitesh)

	fmt.Printf("age is %v", hitesh.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
