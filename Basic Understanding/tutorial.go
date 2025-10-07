//go:build first

package main // required at the top of every Go source file

import "fmt" // import the fmt package for formatted I/O

func main() { // main function - entry point of the program
	fmt.Println("Hello, World!") // print "Hello, World!" to the console
}

// no as go is a complied language, so we have to first convert it into low level language specific binary file
// then we can run that binary file to see the output

// to compile and run the go program, we can use the following commands in terminal
// go build tutorial.go  // this will create a binary file named tutorial (or tutorial.exe on Windows)
// ./tutorial            // this will run the binary file and we will see the output "Hello, World!" in the console

// or we can use the following command to compile and run the go program in one step
// go run tutorial.go   // this will compile and run the program in one step and we will see the output "Hello, World!" in the console

// Note: The comments in the code are for explanation purposes and are not part of the actual code.
