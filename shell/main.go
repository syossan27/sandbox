package main

import "fmt"

func main() {
	var input string
	fmt.Print("What's your name?")
	fmt.Scanln(&input)

	fmt.Printf("Hello, %s!\n", input)
}
