package main

import "fmt"


type Dog struct {
	Breed string
	Weight int
	Sound string
}

func (dog Dog ) Speak() {
	fmt.Println(dog.Sound)
}

func main() {

	poodle := Dog("Poodle", 37, "Woof")
	fmt.Println(poodle)
}