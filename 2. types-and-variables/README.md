# Types and Variables in Go

This project demonstrates the fundamental concepts of variables and data types in Go.

## Basic Types in Go

Go is a **statically typed** language, which means every variable has a specific type that cannot change.

### 1. String Type

Strings represent text data and are enclosed in double quotes.

```go
var name string = "John"
var message string
message = "Hello"
greeting := "Hi there!" // Short declaration
```

**Key Points:**
- Use double quotes `"text"` (not single quotes)
- Default zero value: `""` (empty string)
- Strings are immutable (cannot be changed after creation)

### 2. Integer Types

Integers are whole numbers (no decimal points).

```go
var age int = 25
var count int // Default value is 0
score := 100  // Short declaration
```

**Integer Sizes:**
- `int` - Platform dependent (32 or 64 bit)
- `int8` - 8 bit (-128 to 127)
- `int16` - 16 bit (-32,768 to 32,767)
- `int32` - 32 bit (-2 billion to 2 billion)
- `int64` - 64 bit (very large range)
- `uint` - Unsigned (positive only), platform dependent
- `uint8`, `uint16`, `uint32`, `uint64` - Unsigned variants

**Default zero value:** `0`

### 3. Floating Point Types

Floating point numbers have decimal points.

```go
var price float64 = 19.99
var temperature float32 = 36.6
pi := 3.14159 // Defaults to float64
```

**Types:**
- `float32` - 32 bit floating point
- `float64` - 64 bit floating point (more precise, recommended)

**Default zero value:** `0.0`

### 4. Boolean Type

Booleans represent true/false values.

```go
var isActive bool = true
var hasAccess bool // Default is false
isStudent := false
```

**Possible values:** Only `true` or `false`  
**Default zero value:** `false`

## Variable Declaration Methods

Go provides three ways to declare variables:

### Method 1: Full Declaration with Type
```go
var name string = "John"
```
- Use `var` keyword
- Specify the variable name
- Specify the type
- Assign a value

### Method 2: Type Inference
```go
var name = "John"
```
- Compiler automatically determines type from the value
- Still uses `var` keyword

### Method 3: Short Declaration `:=`
```go
name := "John"
```
- **Most common method**
- Only works inside functions (not at package level)
- Type is automatically inferred
- Shorter and cleaner syntax

### Declaration Without Assignment
```go
var name string
// name is "" (empty string)

var age int
// age is 0
```
- Variables get their "zero value" by default
- Zero values: `0` for numbers, `""` for strings, `false` for booleans

## Multiple Variable Declarations

### Method 1: Same Line
```go
var name, city string = "John", "NYC"
x, y, z := 1, 2, 3
```

### Method 2: Grouped Declaration
```go
var (
    name string = "John"
    age  int    = 25
    city string = "NYC"
)
```

## Constants

Constants are values that cannot be changed after declaration.

```go
const PI = 3.14159
const CompanyName = "TechCorp"
const MaxConnections = 100
```

**Rules for Constants:**
- Use `const` keyword
- Must be assigned when declared
- Cannot be changed later
- Cannot use `:=` syntax
- Usually written in UPPERCASE or PascalCase

## Type Conversion

Go requires **explicit** type conversion (no automatic conversion).

```go
var intNum int = 42
var floatNum float64 = float64(intNum) // Convert int to float64

var pi float64 = 3.14
var wholeNum int = int(pi) // Converts to 3 (truncates decimal)
```

**Important:** You must explicitly convert between types, even between `int` and `int64`!

## Zero Values (Default Values)

When you declare a variable without assigning a value, it gets a zero value:

| Type | Zero Value |
|------|------------|
| `string` | `""` (empty string) |
| `int`, `int8`, `int16`, `int32`, `int64` | `0` |
| `uint`, `uint8`, `uint16`, `uint32`, `uint64` | `0` |
| `float32`, `float64` | `0.0` |
| `bool` | `false` |

## Important Rules

1. **Variable names must start with a letter** (or underscore)
2. **Variables must be used** - Unused variables cause compilation errors
3. **Case matters** - `name` and `Name` are different variables
4. **Exported vs Unexported:**
   - Variables starting with uppercase are "exported" (public)
   - Variables starting with lowercase are "unexported" (private to package)

## Common Mistakes

### ❌ Wrong: Can't change type
```go
var name string = "John"
name = 123 // ERROR! Can't assign int to string variable
```

### ✅ Correct: Convert types
```go
var name string = "John"
var age int = 25
message := name + " is " + string(age) // Need conversion
```

### ❌ Wrong: := outside function
```go
package main

name := "John" // ERROR! Can't use := at package level
```

### ✅ Correct: Use var at package level
```go
package main

var name = "John" // Correct
```

## Running the Program

```bash
# Run the program
go run main.go

# Build executable
go build main.go

# Format code
go fmt main.go
```

## Practice Exercises

Try modifying `main.go` to:
1. Create variables for your own name, age, and favorite number
2. Declare constants for important values (like your birth year)
3. Practice type conversions between int and float64
4. Create a boolean variable and print it
5. Declare multiple variables in one line

## Next Steps

After mastering types and variables, you can learn:
- Functions and parameters
- Control structures (if/else, loops)
- Arrays and slices
- Structs and custom types

Happy coding! 🎯
