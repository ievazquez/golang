package main

import (
	"fmt"
	"io"
	"io/ioutil"
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

	bytes := [] byte(content)
	ioutil.WriteFile("./fromBytes.txt", bytes, 0644)
	
}