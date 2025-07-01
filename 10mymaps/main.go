package main

import "fmt"

func main() {
	fmt.Println("maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println(languages)
	fmt.Println("long for js is: ", languages["JS"])

	// delete(languages, "RB")
	fmt.Println(languages)

	for _, value := range languages {
		fmt.Printf("value is %v \n", value)
	}
}
