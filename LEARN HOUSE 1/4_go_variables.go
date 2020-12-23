package main

import "fmt"

func main() {
	// Go is a statically typed programming language. This means that variables always have a specific type associated with it and cannot be changed.

	// Syntax => var <name> <type>
	var hello_str string
	hello_str = "Hello World!"
	fmt.Println(hello_str)

	// Syntax => var <name> <type> = <expression>
	var name string = "Numan Ibn Mazid"
	fmt.Println(name)

	// Shorthand for assignment + declaration => <name> := <expression>
	message := "Learning GoLang"
	fmt.Println(message)

	// Multiple variable declaration
	var a, b int = 10, 20
	fmt.Println(a, b)

	var c, d string
	c, d = "Apple", "Banana"
	fmt.Println(c, d)

	// const variable => cannot be changed or modified
	// syntax => const <name> <type> = <expression>
	const gender string= "Male"
	fmt.Println(gender)
}
