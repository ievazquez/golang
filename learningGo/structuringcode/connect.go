package main


import "fmt"

var isConnected bool = false


func connect() {
	isConnected = true
	fmt.Println("Connected to database!")
}

func disconnect() {
	isConnected = false
	fmt.Println("Disconnected from database!")
}

func doSomething() {
	connect()
	fmt.Println("Deferring disconnect!")
	defer disconnect()
	fmt.Printf("Connection open: %v\n",isConnected)
	fmt.Println("Do something")
}

func main() {
	fmt.Printf("Connection open: %v\n",isConnected)
	doSomething()
	fmt.Printf("Connection open: %v\n",isConnected)
}

