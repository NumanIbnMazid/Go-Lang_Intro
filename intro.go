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

	// add
	var sum = 11 + 12
	fmt.Println("The sum is:", sum)

	var num1 int = 10
	var num2 int = 12
	var result int = num1 + num2
	fmt.Println("The sum is:", result)

	var num3 uint
	num3 = 66
	fmt.Println("The number is:", num3)

	// declare multiple variables in single line
	var num4, num5 int
	num4, num5 = 200, 300
	fmt.Println("The numbers are:", num4, ",", num5)

	// var assignment
	var firstName = "Numan"
	lastName := "Ibn Mazid" // var lastName string = "Ibn Mazid"
	fullname := firstName + " " + lastName
	fmt.Println("Full name is:", fullname)
	age := 26
	fmt.Println("Age is:", age)

	// const variables - cannot be changed
	const gender string = "Male"
	// gender = "female" // will throw error
	fmt.Println("Gender is:", gender)
}
