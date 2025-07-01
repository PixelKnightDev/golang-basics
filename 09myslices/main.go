package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("slices basics")

	var fruitList = []string{"apple", "tomato", "peach"}
	fmt.Printf("The type of fruitList is %T\n", fruitList)
	fmt.Println("the size is ", len(fruitList))

	fruitList = append(fruitList, "mango", "banana")

	fmt.Println(fruitList)
	fmt.Printf("the type of fruist list is: %T\n", fruitList)

	// fruitList = append(fruitList[1:3])
	//print part of slice

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 343
	highScores[2] = 543
	highScores[3] = 123

	fmt.Println(highScores)

	highScores = append(highScores, 783, 1000, 873)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

}
