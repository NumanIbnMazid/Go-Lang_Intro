package main

import "fmt"

func calc(num1, num2 int) (int, int) {
	var res = num1 + num2
	var res2 = num1 - num2
	return res, res2
}

func add(num1, num2 int) int {
	var res = num1 + num2
	return res
}

func multiDivi(num1, num2 int) (multRes, divRes int) {
	multRes = num1 * num2
	divRes = num1 / num2
	return
}

func main() {
	num1 := 11
	num2 := 12

	result, result2 := calc(num1, num2)
	fmt.Println("The results are:", result, result2)

	addRes := add(2, 3)
	fmt.Println("Add result is:", addRes)

	multRes, divRes := multiDivi(50, 5)
	fmt.Println("Multiplication and Division results are:", multRes, divRes)

}
