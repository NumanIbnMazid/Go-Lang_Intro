package main

import "fmt"

func main() {
	var x = 25
	if (x > 10) {
		fmt.Printf("%d is greater than 10\n", x)
	}

	var age = 24
	if age >= 18 && age <= 30 {
		fmt.Println("I am young!")
	}

	var is_rainning bool = true
	if is_rainning == false {
		fmt.Println("I will go outside.")
	} else {
		fmt.Println("I will stay at home.")
	}

	var number int = 49
	if number >= 100 {
		fmt.Println("The number is greater than or equal 100")
	} else {
		if number >= 50 {
			fmt.Println("The number is greater than or equal 50")
		} else {
			fmt.Println("The number is less than 50")
		}
	}

	if x := 100; x%2 == 0 {
		fmt.Printf("%d is even\n", x)
	} else {
		fmt.Printf("%d is odd\n")
	}
}
