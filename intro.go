package main

import (
	"fmt"
)

func helloWorld() {
	fmt.Print("Hello World!")
	fmt.Println("Hello in Go!")
}

func main() {
	var word1 = "Apple"
	var word2 = word1
	word2 = "banana"
	fmt.Println(word1, word2)

	// hello world function call
	helloWorld()
}
