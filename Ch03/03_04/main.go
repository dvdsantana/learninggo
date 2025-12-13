package main

import (
	"fmt"
	"sort"
)

func main() {
	// Slice
	//var colors = []string{"Red", "Green", "Blue"}
	var colors = make([]string, 0, 3)
	colors = append(colors, "Red")
	colors = append(colors, "Green")
	colors = append(colors, "Blue")
	fmt.Println(colors)
	colors = append(colors, "Yellow")
	fmt.Println(colors)

	colors = remove(colors, 2)
	fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println(colors)

}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}