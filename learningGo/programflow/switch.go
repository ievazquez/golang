package main


import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	//dow := rand.Intn(6) + 1
	var result string

	switch dow := rand.Intn(6) + 1; dow {
	case 1:
		result = "Sunday"
	case 7:
		result = "Saturday"
	default:
		result = "Weekday"
	}
	//fmt.Printf("Day %v, %v\n", dow, result )
	fmt.Printf("Day %v\n",  result )

	x := - 42

	switch {
	case x < 0:
		result = "< 0"
	case x == 0:
		result = "== 0"	
	default:
		result = "> 0"	
	}
	fmt.Printf("Day %v\n",  result )

}