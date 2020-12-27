package main

import "fmt"

func main() {
	/*
	In go lang go to statements are used to alter the normal execution of a program and transfer control to a labeled statement in the same program. The label is an identifier, it can be any plain text and can be set anywhere in the Go program above or below to go to statement. When a go to statement is encountered, compiler transfers the control to a label; and begin execution from there.
	*/

	// Election Example
	var age int

	election:
		fmt.Print("Enter your age ")
		fmt.Scanln(&age)

	if (age <= 17) {
		fmt.Println("You are not eligible to vote!")
		goto election
	} else {
		fmt.Println("You are eligible to vote!")
	}

	// Choice question example

	question1:
		var q1 int
		fmt.Print("1/1 = ")
		fmt.Scanln(&q1)

	if q1 == 1 {
		fmt.Println("Question 1 passed!")
		goto question2
	} else {
		fmt.Println("Wrong answer!")
		goto question1
	}

	question2:
		var q2 bool
		fmt.Print("Go Lang is statically typed. true or false?")
		fmt.Scanln(&q2)

	if q2 == true {
		fmt.Println("Question 2 passed!")
	} else {
		fmt.Println("Wrong answer!")
		goto question2
	}

}
