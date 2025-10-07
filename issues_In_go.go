/*
------------------------------

Issue1: 'cannot define main in multiple files'
🧩 Case 3: Multiple files with their own func main() in the same folder

If your folder has multiple .go files (e.g., tutorial.go and example.go), and both have:

    func main() {}

then Go will also show this error.

Fix:

Keep only one file with func main().

Other files can have helper functions, but not another main.

Issue1 solution:
1- Remove or rename extra main functions so only one main function exists per folder.
2- Alternatively, separate files with main functions into different folders.
3-
Sure 👍 Here’s the **build tags method** in **3 simple points:**

1. 🏷️ Add a special comment at the **top of each file**, like

       //go:build first

   or

       //go:build second

2. 💡 Each file can have its own func main() — only **one tag’s file** will be compiled when you run it.

3. 🚀 Run a specific one using:

       go run -tags=first .

   or

       go run -tags=second .

This way, you can have multiple main functions in the same folder, but only one will be active at a time based on the tag you choose when running the program.

note: You only need to run "go mod init" once per project folder.
After that, you can use any number of build tags or go run commands without repeating it.
go.mod -> Created file that tracks Go version & dependencies

In my case :

(pyenv) inshal@inshal-HP-Laptop-15s-fq5xxx:~/FYP/Go language Learning/Go projects/Beginner Project/Project1/tutorial1$ go mod init tutorial1
go: creating new go.mod: module tutorial1
go: to add module requirements and sums:
        go mod tidy

------------------------------

issue2: 'if a variable is declared and not used, the compiler will throw an error' (unused variable error)
solution: use the variable or remove it

------------------------------

issue3: 'again and again specifying data type for each variable is tedious'
Solution: use type inference with :=

e.g., instead of
var a int = 10
you can write
a := 10  (here Go infers that a is of type int)

------------------------------
*/