package main

import "fmt"

func main() {
	println("Functions")
	doSomething()
}

func doSomething() {
	println("Doing something")
	value1 := 5
	value2 := 10
	result := addValues(value1, value2)
	fmt.Println("The sum of", value1, "and", value2, "is", result)

	value3 := 15
	result, count, average := addAllValues(value1, value2, value3)
	fmt.Println("The sum of", value1, "and", value2, "and", value3, "is", result)
	fmt.Println("The count of", value1, "and", value2, "and", value3, "is", count)
	fmt.Println("The average of", value1, "and", value2, "and", value3, "is", average)
	
	values := []int{1, 2, 3, 4, 5}
	result, count, average = addAllValues(values...)
	fmt.Println("The sum of", values, "is", result)
	fmt.Println("The count of", values, "is", count)
	fmt.Println("The average of", values, "is", average)
}

func addValues(a, b int) int {
	return a + b
}

func addAllValues(values...int) (int, int, float64) {
	total := 0
	for _, value := range values {
		total += value
	}
	count := len(values)
	average := float64(total) / float64(count)
	return total, count, average
}
