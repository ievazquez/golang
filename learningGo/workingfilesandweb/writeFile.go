package main

import (
	"fmt"
	"io"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	content := "Hello World!"

	file, err := os.Create("./fromString.txt")
	checkError(err)
	defer file.Close()
	ln, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("All done with fle of %v\n", ln)
}