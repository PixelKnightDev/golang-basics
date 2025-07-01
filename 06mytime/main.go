package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time study for golang")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))
	//these are all variables for just the position and what you want
	//follow the exact=> 02 for date 01 for month, 15:04:05 for time and capital Monday for days of the week

	createdDate := time.Date(2020, time.July, 12, 6, 30, 0, 0, time.UTC)

	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("02-01-2006 15:04:05 Monday"))

}
