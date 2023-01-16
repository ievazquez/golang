package main

import (
	"fmt"
)


func myFunc() {
	defer fmt.Println("Deferred in the function")
	fmt.Println("Not deferred in the function")
}
func main () {
	
	fmt.Println("Open the file!")
	defer fmt.Println("Statement 1")
	defer fmt.Println("Statement 2")
	defer fmt.Println("Statement 3")
	defer fmt.Println("Statement 4")
	defer fmt.Println("Statement 5")
	fmt.Println("Close the file!")
	myFunc()

	x := 1000

	defer fmt.Println("value of x: ", x)
	x++
	fmt.Println("value after incrementing ", x)
}


