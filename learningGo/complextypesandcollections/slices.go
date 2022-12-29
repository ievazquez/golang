package main

import ( 
	"fmt"
	"sort"
)

func main() {
	var colors = [] string {"Blue", "Red", "Green"};

	fmt.Println(colors)
	colors = append(colors, "Pink")

	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors) - 1])
	fmt.Println(colors)
	
	numbers := make( [] int, 5 , 5)
	numbers[0] = 11
	numbers[1] = 2
	numbers[2] = 30
	numbers[3] = 42
	numbers[4] = 5

	//numbers[5] = 6
	numbers = append(numbers, 6)
	fmt.Println(numbers)

	fmt.Println( cap(numbers))

	sort.Ints(numbers)
	fmt.Println(numbers)

}