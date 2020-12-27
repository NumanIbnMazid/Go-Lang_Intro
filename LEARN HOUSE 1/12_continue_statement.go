package main

import "fmt"

func main() {
	/*
	The continue statement gives you way to skip over the current iteration of any loop. When a continue statement is encountered in the loop, the Go interpreter ignores the rest of the statements in the loop body for current iteration and returns the program execution to the very first statement in the loop body. It does not terminate the loop rather continues to the next iteration.
	*/

	var counter = 0

	for counter < 10 {
		counter++

		if (counter == 5) {
			fmt.Println("5 is skipped")
			continue
			fmt.Println("This won't be printed too")
		}
		fmt.Printf("Number is %d\n", counter)
	}
	fmt.Println("out of loop")

	// labeled continue
	/*
	The standard unlabeled continue statement is used to skip the current iteration of the nearest enclosing loop. In GoLang, there is another form of continue (labeled continue) statement is used to skip the iteration of specified loop (can be outer loop).
	*/

	var i, j int = 0, 0

	outerfor:
		for i < 3 {
			j = 0
			i++

			for j < 3 {
				j++
				fmt.Printf("i is %d and j is %d\n", i, j)
				if (i == 2) {
					fmt.Println("I am skipped")
					continue outerfor
				}
				fmt.Println("I am printed")
			}
		}
}