# Functions and Return Types in Go

This project demonstrates how to create and use functions in Go, including different types of parameters and return values.

## What is a Function?

A function is a reusable block of code that performs a specific task. Functions help organize code, avoid repetition, and make programs easier to understand.

## Basic Function Syntax

```go
func functionName(parameters) returnType {
    // function body
    return value
}
```

## 1. Basic Function (No Parameters, No Return)

The simplest type of function:

```go
func greet() {
    fmt.Println("Hello, World!")
}

// Call it
greet() // Output: Hello, World!
```

**Key Points:**
- Use `func` keyword to declare a function
- Function name follows Go naming conventions
- Parentheses `()` even when no parameters
- Curly braces `{}` contain the function body

## 2. Functions with Parameters

Functions can accept input values (parameters):

```go
func greetPerson(name string) {
    fmt.Println("Hello,", name + "!")
}

greetPerson("Alice") // Output: Hello, Alice!
```

**Multiple Parameters:**
```go
func add(a int, b int) int {
    return a + b
}

// Shorthand (when parameters have same type)
func multiply(x, y int) int {
    return x * y
}
```

## 3. Return Values

Functions can return values using the `return` keyword:

```go
func add(a int, b int) int {
    return a + b
}

result := add(10, 20) // result = 30
```

**Important:**
- Specify return type after parameters
- Must return a value of the specified type
- Function execution stops at `return`

## 4. Multiple Return Values

Go functions can return multiple values - a unique and powerful feature!

```go
func divide(dividend, divisor int) (int, int) {
    quotient := dividend / divisor
    remainder := dividend % divisor
    return quotient, remainder
}

// Call and capture both values
q, r := divide(17, 5) // q = 3, r = 2
```

**Common Pattern - Returning Error:**
```go
func safeDivide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

## 5. Named Return Values

You can name return values in the function signature:

```go
func rectangleStats(length, width int) (area int, perimeter int) {
    area = length * width
    perimeter = 2 * (length + width)
    return // Naked return - returns named values
}
```

**Benefits:**
- Self-documenting code
- Can use "naked return" (just `return`)
- Variables are automatically initialized to zero values

**Note:** Naked returns can reduce readability in long functions, use with caution!

## 6. Ignoring Return Values

Use `_` (blank identifier) to ignore values you don't need:

```go
quotient, _ := divide(20, 3) // Ignore remainder
_, remainder := divide(20, 3) // Ignore quotient
```

## 7. Variadic Functions

Functions that accept a variable number of arguments:

```go
func sumAll(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// Call with different numbers of arguments
sumAll(1, 2, 3)       // 6
sumAll(10, 20, 30, 40) // 100
sumAll()              // 0
```

**Key Points:**
- Use `...` before the type
- Creates a slice of that type
- Must be the last parameter
- Can call with any number of arguments (including zero)

## 8. Functions as Variables

In Go, functions are first-class citizens - you can assign them to variables:

```go
func add(a, b int) int {
    return a + b
}

// Assign function to variable
mathFunc := add
result := mathFunc(5, 3) // result = 8
```

## 9. Anonymous Functions

Functions without a name, often used for short operations:

```go
// Assign to variable
square := func(x int) int {
    return x * x
}
fmt.Println(square(5)) // 25

// Immediate execution
result := func(a, b int) int {
    return a - b
}(10, 3) // result = 7
```

## 10. Defer Statement

`defer` postpones function execution until the surrounding function returns:

```go
func demoDefer() {
    defer fmt.Println("This prints last")
    fmt.Println("This prints first")
    fmt.Println("This prints second")
}

// Output:
// This prints first
// This prints second
// This prints last
```

**Common Uses:**
- Closing files: `defer file.Close()`
- Unlocking mutexes: `defer mutex.Unlock()`
- Cleanup operations
- Multiple defers execute in LIFO order (Last In, First Out)

## 11. Recursive Functions

Functions that call themselves:

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

factorial(5) // 5 * 4 * 3 * 2 * 1 = 120
```

**Requirements:**
- Must have a base case (stopping condition)
- Must progress toward the base case
- Can cause stack overflow if too deep

## Function Parameter Rules

### Same Type Shorthand
```go
// Instead of this:
func add(a int, b int, c int) int

// You can write:
func add(a, b, c int) int
```

### Mixed Types
```go
func formatPerson(name string, age int, height float64) string {
    return fmt.Sprintf("%s: %d years, %.2fm", name, age, height)
}
```

## Return Type Rules

### Single Return Type
```go
func add(a, b int) int {
    return a + b
}
```

### Multiple Return Types
```go
func getCoordinates() (int, int) {
    return 10, 20
}

// Or with different types
func getInfo() (string, int, bool) {
    return "Alice", 25, true
}
```

### No Return Type
```go
func printMessage(msg string) {
    fmt.Println(msg)
    // No return statement needed
}
```

## Best Practices

1. **Use descriptive names:** `calculateTotal()` not `calc()`
2. **Keep functions short:** One function, one purpose
3. **Return errors:** Use `(result, error)` pattern for error handling
4. **Use named returns sparingly:** Only when they improve clarity
5. **Document exported functions:** Use comments for public functions
6. **Avoid side effects:** Functions should be predictable

## Common Patterns

### Error Handling Pattern
```go
func readFile(filename string) (string, error) {
    // ... read file
    if err != nil {
        return "", err
    }
    return content, nil
}

// Usage
content, err := readFile("data.txt")
if err != nil {
    // handle error
}
```

### Builder Pattern
```go
func buildMessage(parts ...string) string {
    result := ""
    for _, part := range parts {
        result += part + " "
    }
    return result
}
```

## Running the Program

```bash
# Run the program
go run main.go

# Build executable
go build

# Format code
go fmt main.go
```

## Practice Exercises

Try creating these functions:

1. `subtract(a, b int) int` - Returns a minus b
2. `isEven(n int) bool` - Returns true if n is even
3. `max(a, b int) int` - Returns the larger number
4. `greetMultiple(names ...string)` - Greets multiple people
5. `celsiusToFahrenheit(c float64) float64` - Temperature converter
6. `swap(a, b int) (int, int)` - Returns values in reverse order

## Key Takeaways

- Functions organize and reuse code
- Parameters pass data into functions
- Return values pass data out of functions
- Go supports multiple return values (unique feature!)
- Use `_` to ignore unwanted return values
- Variadic functions accept variable arguments
- Defer executes code before function returns
- Functions can be assigned to variables

## Next Steps

After mastering functions, learn about:
- Packages and imports
- Methods and receivers (functions on types)
- Interfaces
- Error handling patterns
- Closures and higher-order functions

Happy coding! 🚀
