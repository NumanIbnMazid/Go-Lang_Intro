package main

import (
	"fmt"
	"time"
)

func main() {

	// if condition
	num := 1
	if num == 1 {
		fmt.Println("One")
	} else if num == 2{
		fmt.Println("Two")
	} else {
		fmt.Println("None")
	}

	// switch case
	num2 := 2
	switch num2 {
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		default:
			fmt.Println("None")
	}

	// day check
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
		case today + 0:
			fmt.Println("Today")
		case today + 1:
			fmt.Println("Tomorrow")
		case today + 2:
			fmt.Println("In two days")
		default:
			fmt.Println("Too far away.")
	}

	fmt.Println("Time:", time.Now())
}

