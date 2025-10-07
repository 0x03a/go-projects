//go:build third

// input taking here.
package main

import "fmt"

func main() {
	fmt.Println("This is the third main function.")
	var name string
	print("Enter your name: ")
	fmt.Scan(&name) // taking input from user, for all data types
	//fmt.Scan("Enter your name", &name) , can't do this
	// fmt.Scan() function is used to take input from user.
	// It takes the address of the variable as an argument.

	// now we require ampersand (&) before variable name to store the input value in that variable
	// var name string is stored at random memory address like 0xc000010230 in RAM.
	// So we need to provide that address to the Scan function to store the input value in that variable.
	// &name will give the address of the variable name.
	// So the input value will be stored at that address in RAM.

	// Now we can print the value of the variable name.

	fmt.Println("Hello", name)

}
