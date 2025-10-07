//go:build five

// loops in go + arrays and slices
package main

import "fmt"

func main() {
	fmt.Println("This is the fifth main function.")
	// for loop

	/*

			You have a **syntax error** in your Go `for` loop.
		Here’s the **correct version** 👇

		```go
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		```

		### Explanation:

		* `:=` is used for declaring and initializing a variable in Go (short declaration).
		* The syntax for a `for` loop in Go is:

		  ```go
		  for init; condition; post {
		      // code
		  }
		  ```

		✅ **Correct:** `for i := 0; i < 5; i++ { ... }`
		❌ **Incorrect:** `for var i int = 0; i < 5; i++ { ... }` — `var` cannot be used inside `for` like that.


	*/

	// 🌀 1. for loop
	// ✅ Initialization → Condition → Increment
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//Runs while i < 5.

	//🔁 2. While-style loop

	// Go doesn’t have a while keyword — you use for with only a condition:
	// Equivalent to a “while” loop in other languages.
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// 3. Infinite loop
	// Runs endlessly unless you use break or return.
	for {
		fmt.Println("Running forever!")
	}

}
