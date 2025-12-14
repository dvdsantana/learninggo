package main

import (
	"fmt"
	"math"
)

func main() {
	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	sum = math.Round(sum*100) / 100
	fmt.Printf("Float sum rounded %v\n\n", sum)

	fmt.Println("Pi value is, ", math.Pi)

	circleRadios := 12.5
	circleArea := math.Pi * math.Pow(circleRadios, 2)
	fmt.Printf("Circle area:%.2f\n", circleArea)

}