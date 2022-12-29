package main

import (
	"fmt"
)

type Dog struct {
	Breed string
	Weigth int
}

func main() {

	poodle := Dog{ Breed: "Poodgle", Weigth: 34}

	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)

	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed,poodle.Weigth)
	fmt.Println("done")
}