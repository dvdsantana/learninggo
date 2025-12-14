package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Switch")

	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	var result string

	switch dayNumber := int(weekday); dayNumber {
	case 1:
		result = "Monday"
	case 2:
		result = "Tuesday"
	case 3:
		result = "Wednesday"
	case 4:
		result = "Thursday"
	case 5:
		result = "Friday"
	default:
		result = "Weekend"
	}
	fmt.Println(result)

	x := 0
	switch {
	case x < 0:
		fmt.Println("x is less than 0")
		fallthrough
	case x == 0:
		fmt.Println("x is equal to 0")
		fallthrough
	default:
		fmt.Println("x is greater than 0")
	}
}