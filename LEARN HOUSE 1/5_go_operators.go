// In go most of the operators are binary that means they took two operands but a few of them takes only one operand.

/*
Available operators in GoLang:
	~ Arithmetic Operators
	~ Assignment Operators
	~ Comparison (Relational) Operators
	~ Logical Operators
	~ Bitwise Operators
*/

package main

import "fmt"

func main() {
	/*
	------- Arithmetic Operators -------:
	Arithmetic Operators are used to perform arithmetic operations like addition, subtraction, multiplication, division, %modulus, exponent etc.

	EXAMPLE: a := 20 and b := 10

	Operator Name			Description							Example
	+		Addition		Addition of given operands			a+b returns 30
	- 		Subtraction		Subtraction of second from first	a-b returns 10
	* 		Multiply		Multiplication of given operands	a*b returns 200
	/		Division		Returns Quotient after division		a/b returns 2
	% 		Modulus			Returns reminder after division		a%b returns 0
	*/

	var a, b = 20, 10
	fmt.Println("Addition: ", a + b)
	fmt.Println("Subtraction: ", a - b)
	fmt.Println("Multiply: ", a * b)
	fmt.Println("Division: ", a / b)
	fmt.Println("Modulus: ", a % b)

	/*
	------- Assignment Operators -------:
	Assignment operators are used to assign values to a variable.

	Operator Description					Expression
	=		Assignment Operator				a=b
	+=		Add and assign					a+=b : a=a+b
	-=		Subtract and assign				a-=b : a=a-b
	*= 		Multiply and assign				a*=b : a=a*b
	/=		Divide and assign				a/=b : a=a/b
	%=		Mod and assign					a%=b : a=a%b
	<<=		Left shift and assign			a<<=5 : a=a<<5
	>>= 	Right shift and assign			a>>=5 : a=a<<5
	&=		Bitwise and assign				a&=5 : a=a&5
	^=		Bitwise exclusive OR and assign	a^=5 : a=a^5
	|= 		Bitwise inclusive OR and assign	a=a|5
	*/

	var c, d int = 30, 5
	var temp = 0

	c += d
	temp = c
	fmt.Println("30+=5 : ", c)

	c -= d
	fmt.Printf("%d-=5 : %d", temp, c)
	temp = c

	println()
	c *= d
	fmt.Printf("%d*=5 : %d", temp, c)
	temp = c

	println()
	c /= d
	fmt.Printf("%d/=5 : %d", temp, c)
	temp = c

	println()
	c %= d
	fmt.Printf("%d%s=5 : %d", temp, "%", c)
	temp = c

	println()
	c, temp = 2, 2
	c <<= d
	fmt.Printf("%d<<=5 : %d", temp, c)
	temp = c

	println()
	c, temp = 265, 265
	c >>= d
	fmt.Printf("%d>>=5 : %d", temp, c)
	temp = c

	println()
	c &= d
	fmt.Printf("%d&=5 : %d", temp, c)
	temp = c

	println()
	c ^= d
	fmt.Printf("%d^=5 : %d", temp, c)
	temp = c

	println()
	c |= d
	fmt.Printf("%d|=5 : %d", temp, c)
	temp = c
	
}