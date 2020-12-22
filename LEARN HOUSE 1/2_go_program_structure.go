/*
Every Go program is consists of following building blocks:
- Documentations Block
- Preprocessor Statements
- The main() function
- Local Variable Declarations
- Program Statements
- User Defined Functions
*/

/*
======= Documentations Block =======
Documentation block is the header of any Go program structure which is mainly consists of set of comments, that is used to provide the information about the program written like name of the program, date of creation, date of last modification, licensing or copyrights information, author name, utility of program and any other information that programmer wish to put for the references.
*/

/*
- Documentation Block -
Program Name: Go Program Structure
Version: 1.0
Description: Basic Go Program about Go Program Structure
Utility:
Author: Numan Ibn Mazid
License:
Copyright: Numan Ibn Mazid
Date Created: 2020-12-22
Date Last Modified: 2020-12-22
*/


// ======= Package Declaration =======
package main
/*
In Package Declaration we declare the package name in which the program would included. Go program runs in packages, so it is required to define a package. Every package has a name and path associated with it.
*/

// ======= Preprocessor Statements =======
import "fmt"
/*
This is the section where we define all the preprocessor statements. The "import" preprocessor statements tells the compiler to import required packages prior to compilation of any Go program.
*/

// ======= The main() function =======
/*
This is the most vital part of each and every Go program, it is mandatory for the Go program to have the main() function. The Go program execution starts with the main() function. Go program cannot be executed without the main() function.
*/
func main() {
	// ======= Local Variable Declaration =======
	var msg string = "Hello world"

	// ======= Statements and Expressions =======
	/*
	This is the section where we place our main logic of the program which included the executable statements, that tell the computer to perform a specific action. Program statement can be an input-output statements, arithmetic statements, control statements, simple assignment statements and any other statements and it also includes comments that are enclosed within block comments.
	*/
	if 2 == 2 {
		fmt.Println(msg)
	}

	/*
	User Defined function Call
	*/
	who_am_i()
}

// ======= User defined functions =======
func who_am_i() {
	fmt.Println("I am Numan Ibn Mazid")
}
/*
This is where we put all the user defined sub-program or custom functions created to perform a specific task and called inside the main() function.
*/