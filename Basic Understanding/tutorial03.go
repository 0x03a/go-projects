//go:build four

package main

import "fmt"

func main() {
	fmt.Println("This is the fourth main function.")
	var age int
	print("Enter your age: ")
	fmt.Scan(&age)
	fmt.Println(age <= 10)

	// if age is less than or equal to 10, then print "You are a child"
	if age <= 10 {
		fmt.Println("You are a child")
	} else if age > 10 && age <= 18 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are an adult")
		return // exit from the main function
		// any code written after return statement will not be executed
	}

	// switch case
	// can also be used instead of if-else
	// in each case block, we can have multiple statements without using curly braces

	switch {
	case age <= 10:
		fmt.Println("You are a child1")
		fmt.Println("You are a child2")

		fmt.Println("You are a child3")
	case age > 10 && age <= 18:
		fmt.Println("You are a teenager")
	default:
		fmt.Println("You are an adult")
	}

	fmt.Printf("What is better, the RTX 3080 or the RTX 4090? ")
	var answer1 string
	var answer2 string
	fmt.Scan(&answer1, &answer2)
	if answer1+" "+answer2 == "RTX 4090" {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Wrong answer!")
	}

	// type conversion
	var x int = 10
	var y float64 = 5.5
	// var z int = x + y // this will throw an error as we cannot add int and float64 directly
	var z float64 = float64(x) + y // type conversion of x to float64
	fmt.Println("z:", z)

}

// ✅ Postfix increment and decrement operators
// x++
// x--
// Note: There are NO prefix increment or decrement operators in Go
// ++x ❌
// --x ❌

// ✅ Logical operators
// && (AND)
// || (OR)
// !  (NOT)

// ✅ Relational operators
// == (equal to)
// != (not equal to)
// >  (greater than)
// <  (less than)
// >= (greater than or equal to)
// <= (less than or equal to)

// ✅ Arithmetic operators
// + (addition)
// - (subtraction)
// * (multiplication)
// / (division)
// % (modulus)

// ✅ Assignment operators
// =
// +=
// -=
// *=
// /=
// %=

// ✅ Compound assignment examples
// x += 10 // x = x + 10
// x -= 10 // x = x - 10
// x *= 10 // x = x * 10
// x /= 10 // x = x / 10
// x %= 10 // x = x % 10

// ❌ Note: There is NO ternary operator in Go
// x = (condition) ? value1 : value2 // invalid in Go

// Note: The comments in the code are for explanation purposes and are not part of the actual code.
