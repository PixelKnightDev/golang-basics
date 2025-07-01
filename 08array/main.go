package main

import "fmt"

func main() {
	fmt.Println("array basics")

	var fruitList [4]string

	fruitList[0] = "tomato"
	fruitList[1] = "apple"
	fruitList[3] = "banana"
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("length is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}

	fmt.Println("vegetables list is: ", vegList)
	fmt.Println("length is: ", len(vegList))
}
