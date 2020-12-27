package main

import "fmt"

func main() {
	/*
	Break statement gives the way to break or terminate the execution of innermost "for", "switch", or "select" statement that containing it, and transfers the execution to the next statement following it.
	*/

	var count = 0

	for count <= 10 {
		count++

		if (count == 5) {
			break
		}
		fmt.Printf("Inside loop %d\n", count)
	}
	fmt.Println("Out of loop")


	// labeled break
	/*
	The standard unlabeled break statement is used to terminate the nearest enclosing statement. I GoLang, there is another form of break (labeled break) statement is used to terminate specified "for", "switch" or "select" statement. In GoLang, Label is an identifier which is followed by a colon (:) sign.
	*/
	count = 0
	var innerCounter = 0

	outer :
	for count <= 10 {
		count++

		inner:
			for innerCounter <= 5 {
				innerCounter++

				if innerCounter >= 3 {
					break inner
				}
				fmt.Printf("Inner inside loop! %d\n", innerCounter)
			}

		if (count == 5) {
			break outer
		}
		fmt.Printf("Inside loop %d\n", count)
	}
	fmt.Println("Out of loop")

}