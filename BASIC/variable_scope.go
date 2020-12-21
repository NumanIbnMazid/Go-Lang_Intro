package main

import "fmt"

// package level variable
var outside = "ABC"
var b = 10

func demo() {
	// function level variable
	var a = 9
	fmt.Println(a)
}

func main() {
	demo()
	// fmt.Println(a) // ./variable_scope.go:12:14: undefined: a
	fmt.Println("Outside :", outside)
	// function level variable
	var b = 7
	fmt.Println("b redefined (pre: 10, next: 7):", b)
}
