# Hello World - Your First Go Program

This is a simple Hello World program to help you understand the basics of Go programming.

## Understanding the Code Structure

Every Go program follows a specific structure. Let's break down each part:

### 1. Package Declaration
```go
package main
```

**What is a Package?**
- A package is a way to group related Go code together
- Every Go file must start with a package declaration
- `package main` is special - it tells Go this is an executable program (not a library)
- Only the `main` package can have a `main()` function that serves as the entry point
- Think of packages like folders that organize your code

### 2. Import Statement
```go
import "fmt"
```

**What is Import?**
- The `import` keyword brings in other packages that you want to use
- `"fmt"` is a standard library package (comes with Go)
- You must import packages before you can use their functions
- If you import multiple packages, you can use this syntax:
  ```go
  import (
      "fmt"
      "time"
  )
  ```

**What is fmt?**
- `fmt` stands for "format"
- It's used for formatted input/output operations
- Common functions include:
  - `fmt.Println()` - Print with a new line
  - `fmt.Print()` - Print without a new line
  - `fmt.Printf()` - Print with formatting (like %s, %d)
  - `fmt.Scan()` - Read input from user

### 3. Main Function
```go
func main() {
    fmt.Println("Hello World")
}
```

**Understanding the Function:**
- `func` is the keyword to declare a function
- `main()` is the special function where program execution begins
- Every executable Go program must have exactly one `main()` function in the `main` package
- The curly braces `{}` contain the function body (the code to execute)
- `fmt.Println("Hello World")` prints "Hello World" to the console and adds a new line

## File Structure (The Proper Order)

When writing Go programs, always follow this order:

1. **Package declaration** - First line, always
2. **Import statements** - After package, before functions
3. **Functions** - Your code logic goes here

```go
// 1. Package (required)
package main

// 2. Imports (if needed)
import "fmt"

// 3. Functions
func main() {
    // Your code here
}
```

## How to Run the Program

Go provides several commands to work with your code:

### `go run` - Run Without Building
```bash
go run main.go
```
- Compiles and runs your program in one step
- Doesn't create an executable file
- Great for quick testing during development
- Output: `Hello World`

### `go build` - Create an Executable
```bash
go build main.go
```
- Compiles your program and creates an executable file
- Creates a binary file named `main` (or `main.exe` on Windows)
- You can then run it directly: `./main`
- Use `go build -o myprogram main.go` to name the executable

### `go fmt` - Format Your Code
```bash
go fmt main.go
```
- Automatically formats your code according to Go standards
- Fixes indentation, spacing, and style issues
- Always run this before sharing your code!
- Go has one official formatting style (no debates!)

### Other Useful Commands

**`go vet`** - Check for Common Mistakes
```bash
go vet main.go
```
- Examines your code for suspicious constructs
- Catches common errors that the compiler might miss

**`go clean`** - Remove Build Files
```bash
go clean
```
- Removes compiled binaries and cached files
- Cleans up your project directory

**`go mod init`** - Initialize a Module (for larger projects)
```bash
go mod init myproject
```
- Creates a `go.mod` file to manage dependencies
- Required for projects with multiple packages or external dependencies

## Quick Reference

| Command | Purpose |
|---------|---------|
| `go run main.go` | Run the program without creating a file |
| `go build` | Compile and create executable |
| `go fmt` | Auto-format code |
| `go vet` | Check for errors |
| `go clean` | Remove build files |

## Try It Yourself!

1. Run the program:
   ```bash
   go run main.go
   ```

2. Build an executable:
   ```bash
   go build main.go
   ./main
   ```

3. Format the code:
   ```bash
   go fmt main.go
   ```

## Next Steps

Now that you understand the basics:
- Try changing the message from "Hello World" to something else
- Add multiple `fmt.Println()` statements
- Experiment with `fmt.Printf("Hello %s\n", "World")`
- Learn about variables and data types

Happy coding! 🚀
