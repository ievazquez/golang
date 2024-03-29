package main

import (
	"fmt"
	"stringutil"
)

func doTest() {
	fmt.Println("Do Test")
	sum := add(20, 103)
	fmt.Println("sum ", sum)
	sum = addAll(1, 2, 3, 4, 5)
	fmt.Println("sum ", sum)
}

// func add(value1 int, value2 int) int {
func add(value1, value2 int) int {
	return value1 + value2
}

func addAll(values ...int) int {
	sum := 0
	fmt.Printf("%T\n", values)
	for i := range values {
		sum += values[i]
	}
	return sum
}

func main() {
	doTest()
	fullname, len := FullName("Carl", "Johnson")
	fmt.Printf("Full Name %v number of chars %v\n", fullname, len)

	fullname2, len2 := FullNameNakedReturn("Arthur", "Dent")
	fmt.Printf("Full Name %v number of chars %v\n", fullname2, len2)
}
