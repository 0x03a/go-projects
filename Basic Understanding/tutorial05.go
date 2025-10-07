//go:build six

// functions in go
package main

import "fmt"

/*

Note:
Where	Allowed	Example
✅ Global function	   Yes	  func add(a,b int){}
✅ Local anonymous func	Yes,  inside main()	add := func(a,b int){}
❌ Global anonymous func	 No	   add := func(a,b int){} (outside main)


*/

func add(a int, b int) {
	print("Addistion", a+b, "\n")
}

func subtract(a int, b int) {
	print("Subtraction", a-b, "\n")
}

func multiply(a int, b int) {
	print("Multiplication", a*b, "\n")
}

func divide(a int, b int) {
	print("Division", a/b, "\n")
}

// function with return type
func addWithReturn(a int, b int) int {
	return a + b
}

// function with multiple return types
func addAndSubtract(a int, b int) (int, int) {
	return a + b, a - b
}

// function as a parameter
func executeFunction(f func()) {
	f()
}

func main() {
	fmt.Println("This is the sixth main function.")
	i := 5
	j := 5
	add(5, 5)
	subtract(i, j)
	multiply(i, j)
	divide(i, j)

	// function with return type
	result := addWithReturn(i, j)
	fmt.Println("Result of addition with return:", result)

	// function with multiple return types
	sum, difference := addAndSubtract(i, j)
	fmt.Println("Sum:", sum, "Difference:", difference)

	// anonymous function
	func() {
		fmt.Println("This is an anonymous function.")
	}() // immediately invoked

	// function as a variable
	f := func() {
		fmt.Println("This is a function assigned to a variable.")
	}
	f() // invoking the function

	// function as a parameter
	executeFunction(f)
}
