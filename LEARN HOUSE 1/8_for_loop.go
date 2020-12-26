package main

import "fmt"

func main() {

	// for loop
	sum := 0
	for i := 1; i < 5; i ++ {
		sum += 1
	}
	fmt.Println(sum)

	// for loop as while loop
	power := 1
	for power < 5 {
		power *= 2
	}
	fmt.Println(power)

	// infinite loop
	sum2 := 0
	for {
		sum2++ // repeated forever
		fmt.Println("Executing...")
	}
	fmt.Println(sum2)
}
