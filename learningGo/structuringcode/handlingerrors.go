package main

import (
	"fmt"
	"os"
	"errors"
)

func main() {
	f, err := os.Open("filename.txt")

	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(err.Error())
	}

	myError := errors.New("Error!!!!")
	fmt.Println(myError.Error())


	attendance := map [string] bool {
		"Ann": true,
		"Mike":true,
		"John":false,
	}
	attended, found := attendance["Mike"]
	fmt.Println("Mike ", attended, found)

	attended, found = attendance["John"]
	fmt.Println("John ", attended, found)

	attended, found = attendance["Kevin"]
	fmt.Println("Kevin ", attended, found)

}