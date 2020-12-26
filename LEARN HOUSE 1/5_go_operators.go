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
	println()

	/*
	------- Comparison Operators -------:
	Comparison operators are used to evaluate a comparison between two operands. The result of comparison is a Boolean value that can be only true or false. Comparison operators are also referred as relational operators.

	[a:=20, b:=10]

	Operators	Description				Example
	>		Greater than				a>b returns True
	<		Less than					a,b returns False
	>=		Greater than or Equal to	a>=b returns True
	<=		Less than or equal to		a<=b returns False
	==		Is equal to					a==b returns False
	!=		Not equal to				a!=b returns True
	*/

	var e, f int = 20, 10
	if (e > f) {
		fmt.Println("e is greater than f")
	} else {
		fmt.Println("f is greater than e")
	}

	/*
	------- Logical Operators -------:
	Logical operators are used to combine expressions with conditional statements using Logical (AND, OR, NOT) operators which results in True or False.

	[a holds true or 1, b holds false or 0]

	Operator	Name		Description									Expression
	&&		Logical AND		return true if all expressions are true		(a && b) returns false
	||		Logical OR		return true if any expression is true		(a || b) returns true
	!		Logical NOT		return complement of expression				!a returns false
	*/

	var g, h, i int = 20, 10, 25
	var flag, result bool
	flag = false
	result = (g > h) && (g > i)
	fmt.Printf("\n (g > h) && (g > i) : %t \n", result)
	result = (g > h) || (g > i)
	fmt.Printf("\n (g > h) || (g > i) : %t \n", result)
	result = !flag
	fmt.Printf("\n !flag : %t \n", result)

	/*
	------- Bitwise Operators -------:
	Bitwise operator are used to perform bit level operations over its operand. The bitwise logical and shift operators are only applicable to integers.

	Operator	Description
	&			bitwise AND
	|			bitwise OR
	^			bitwise XOR
	&^			bit clear (AND XOR)
	*/

	var j, k, l int
	j = 50
	k = 10
	l = j & k
	fmt.Println("j & k : ", l)
	l = j | k
	fmt.Println("j | k : ", l)
	l = j ^ k
	fmt.Println("j ^ k : ", l)
	l = j &^ k
	fmt.Println("j &^ k : ", l)

	/*
	------- Other Operators -------:
	Operator 	Name 			Description
	&		Address of			&a generates a pointer to a
	*		Pointer to			*a denotes the variable pointed to by a
	<-		Receive Operator	<-ch is the value received from channel ch
	*/

	var m int = 30
	fmt.Println("&m", &m)

	/*
	------- String Concatenation Operator -------:
	in go, string can be concatenated using the + operator or the += assignment operator.
	*/
	var x, y string = "hello ", "world"
	var concatenated = x + y
	fmt.Println(concatenated)

	var z = "Numan "
	z += "Ibn Mazid"
	fmt.Println(z)
}