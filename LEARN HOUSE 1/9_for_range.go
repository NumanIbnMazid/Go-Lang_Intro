package main

import "fmt"

func main() {

	// go for range
	langs := []string{"Go", "C", "C++", "Python"}
	
	for i, s := range langs {
		fmt.Println(i, s)
	}

	// for range with string
	/*
	For a string, the for loop iterates over each character's unicode points. The first value is the starting byte index of the rune and the second the rune itself.
	*/

	for i, s := range "Hello World" {
		fmt.Printf("%U represents %c and it is at position %d\n", s, s, i)
	}

	// for range with map
	// range on map iterates over key/value pairs. It can also iterate over just the keys of a map.

	fruits := map[string]string{"A": "Apple", "B": "Banana", "C": "Cherry"}

	for key, value := range fruits {
		fmt.Printf("%s -> %s\n", key, value)
	}

	for key := range fruits{
		fmt.Println("Key: ", key)
	}

	// range with channel
	// it iterates over values sent on the channel until closed.

	ch := make(chan string)
	
	go func() {
		ch <- "H"
		ch <- "E"
		ch <- "L"
		ch <- "L"
		ch <- "O"
		close(ch)
	} ()

	for n := range ch {
		fmt.Println(n)
	}

}
