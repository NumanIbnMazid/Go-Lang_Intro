/*
Go comes with the following built-in data-types:
	- Numbers
	- Boolean
	- String

- Numbers: Holds numeric values. Go supports teh following numerical data types:
	* Integer: Used to store whole numbers. Go supports several integer types varying internal sizes for storing signed and unsigned integers.
		# Signed Integers
		~ type	- Size					- Range
		~ int8	- 8 bits				- 128-127
		~ int16	- 16 bits				- -2^15 to 2^15
		~ int32 - 32 bits				- -2^31 to 2^31 - 1
		~ int64	- 64 bits				- -2^63 to 2^63
		~ int 	- Platform Independent	- Platform Dependent

		# Unsigned Integers
		~ type		- Size					- Range
		~ uint8		- 8 bits				- 0 - 127
		~ uint16	- 16 bits				- 0 - 2^16 - 1
		~ uint32	- 32 bits				- 0 - 2^32 - 1
		~ uint64	- 64 bits				- 0 - 2^64 - 1
		~ uint		- Platform Dependent	- Platform Dependent

		It is better to use int data type unless you have a specific reason to use sized or unsigned integer types.
		Uninitialized integer types have a default value of 0.
		In Go lang, octal numbers can be declared using 0 prefix and hexadecimal numbers using 0x prefix.

	* Other Integer Types
		~ byte : It is alias for and equivalent to unit8
		~ rune : It is alias for and equivalent to int32, used to represent characters
		~ uintptr : It is used to hold memory address pointers

	- Golang do not supports a char data type, instead it have byte and rune to represent character values. This helps to distinguish characters from integer values.

	Both byte and rune data types are basically holds integers, the byte data type is represented with a ASCII value while the rune data type is represented with Unicode value encoded in UTF-8 format.
	For example, a byte variable with value 'a' is converted to the integers its ASCII equivalent 97.Similarly, a rune variable with a unicode value ' ' is converted to the its corresponding unicode equivalent U+2665, here LJ+ stands for unicode and the numbers are hexadecimal, which is essentially an integer.

	* Golang Float:
		A Float type is used to store numbers that contain a decimal component (real numbers). Go supports float32 and float64 represented with 32-bit and 64-bit in memory respectively. An uninitialized float has default value of 0 and . An untyped floating point values is considered as float64. So in cases of an untyped floating point variable declaration with an initial value without specifying a type explicitly, the compiler will automatically assign float64 type to it.

	* Complex Numbers:
		Go supports two additional types for representing complex numbers (numbers with imaginary parts) complex64 and complexl 28. Each of the type uses float32 and float64 respectively, to represent their real and imaginary parts.

- Boolean:
		The Boolean data type is used to represent the truth values, which can be either True or False. Boolean are commonly used in decision making statements.

- String
	String variable is used to hold series or sequence of characters - letters, numbers, and special characters. Strings are immutable. Strings can either be declared using double quotes "Hello World" or back ticks Hello World .
*/

package main

import "fmt"

func main() {
	// integer
	var varByte byte = 'a'
	var varRune rune = '❤'
	fmt.Printf("%c = %d and %c = %U \n", varByte, varByte, varRune, varRune)

	// Float
	var num1 float32 = 10.1514
	var num2 = 5014.645 // Type inferred as float64
	fmt.Println(num1, num2)

	// Complex
	var num3 = 3 + 7i // Type inferred as `complex128`
	fmt.Println(num3)
	
	// Boolean
	var flag = true
	fmt.Println(flag)

	// String
	// It does not allows newlines,
	var fstr = "Hello World"
	// Can span multiple lines. Escape
	var sstr = `Hello World,
This is multiline Text string.`
	fmt.Println(fstr)
	fmt.Println(sstr)
}
