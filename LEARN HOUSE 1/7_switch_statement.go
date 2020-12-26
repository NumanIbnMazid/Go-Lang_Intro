package main

import "fmt"

func main() {
	var dayOfWeek = 5

	// switch statement
	switch dayOfWeek {
	case 1:
		fmt.Println("Today is Monday.")
		case 2:
		fmt.Println("Today is Tuesday.")
		case 3:
		fmt.Println("Today is Wednesday.")
		case 4:
		fmt.Println("Today is Thursday.")
		case 5:
		fmt.Println("Today is Friday.")
		case 6:
		fmt.Println("Today is Saturday.")
		case 7:
		fmt.Println("Today is Sunday.")
		default:
		fmt.Println("Invalid Weekday!")
	}

	// shorthand
	switch weekDay := 5; weekDay {
	case 1, 2, 3, 4, 5:
		fmt.Println("It is weekday!")
	case 6, 7:
		fmt.Println("It's Weekend!")
		default:
			fmt.Println("Invalid Day!")
	}

	// fallthrough
	/*
	Execute all the following statements after a match found!
	*/
	var x = 2

	switch x {
	case 1:
		fmt.Println("1")
		fallthrough
		case 2:
			fmt.Println("2")
			fallthrough
			case 3:
				fmt.Println("3")
	}

	// No Expression
	var marks = 55
	switch {
	case marks > 60:
		fmt.Println("Passed with 'A' grade")
		case marks < 60 && marks >= 50:
		fmt.Println("Passed with 'B' grade")
		case marks < 50 && marks >= 35:
		fmt.Println("Passed with 'C' grade")
		default:
			fmt.Println("You are Fail!")
	}
}
