package main

import (
	"fmt"
)

func main() {
	fmt.Println("Methods")
	dog := Dog{"Poodle", "Woof"}
	dog.Speak()

	dog.Sound = "Bark"
	dog.Speak()

	println(dog.SpeakThreeTimes())
}

type Dog struct {
	Breed string
	Sound string
}

func (d Dog) Speak() {
	fmt.Println("The", d.Breed, "says", d.Sound)
}

func (d Dog) SpeakThreeTimes() string {
	return fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
}