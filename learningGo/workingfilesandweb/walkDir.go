package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func processPath( path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if path != "." {
		if info.IsDir() {
			fmt.Println("Directory: ",path)
		} else {
			fmt.Println("\tFile: ",path)			
		}
	}
	return nil
}


func main() {
	root, _ := filepath.Abs(".")
	fmt.Println("Processing path:", root)

	err := filepath.Walk(root, processPath)

	if err != nil {
		fmt.Println("Error: ", err)
	}
}