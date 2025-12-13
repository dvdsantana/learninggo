package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")
	poodle := Dog{"Poodle", 34}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v, Weight: %v", poodle.Breed, poodle.Weight)
	poodle.Weight = 35
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
}

type Dog struct {
	Breed string
	Weight int
}