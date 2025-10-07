//go:build second

package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// integer variables
	var a int = 10
	var b int = 20 // handles all +ive and -ive whole numbers

	var c uint = 30 // handles only +ive whole numbers
	// var d int32 = 40         // handles whole numbers up to 2 billion
	// var e int64 = 9000000000 // handles whole numbers up to 9 quintillion
	// float variables

	var f float32 = 5.5   // handles decimal numbers up to 7 digits
	var g float64 = 10.99 // handles decimal numbers up to 15 digits

	// boolean variable
	var isTrue bool = true // handles true or false values

	// printing variables
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("f:", f)
	fmt.Println("g:", g)
	fmt.Println("isTrue:", isTrue)

	// performing arithmetic operations

	var sum int = a + b
	fmt.Println("Sum of", a, "and", b, "is", sum)

	// string variable
	var name string = "Inshal"
	fmt.Println("My name is", name)

	name3 := "Inshal3" // type inference
	fmt.Println("My name is", name3)

	// embedding variables in strings
	fmt.Printf("hello  %v, How are u doing? val is: %v \n", name, a)

}
