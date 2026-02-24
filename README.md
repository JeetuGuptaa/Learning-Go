# Go Learning Journey

Welcome to my Go programming learning repository! This is where I document my progress as I learn the Go programming language.

## About This Repository

This repository contains various Go projects and examples that I'm building while learning. Each folder represents a different topic or concept in Go.

## Projects

### 1. [Hello World](1.%20helloworld/README.md)
My first Go program! This project covers:
- Basic Go program structure
- Understanding packages and imports
- The `fmt` package for output
- Essential Go commands (`go run`, `go build`, `go fmt`)
- File structure and organization

### 2. [Types and Variables](2.%20types-and-variables/README.md)
Learning about data types and variable declarations in Go:
- Basic types (string, int, float64, bool)
- Different ways to declare variables (`var`, `:=`)
- Type conversion and zero values
- Constants and multiple declarations
- Integer sizes and floating point types

### 3. [Functions and Return Types](3.%20functions-and-return-types/README.md)
Understanding how to create and use functions:
- Basic function syntax and parameters
- Return values (single and multiple)
- Named return values and naked returns
- Variadic functions (variable arguments)
- Anonymous functions and defer statements
- Recursion and functions as variables

### 4. [Arrays, Slices, and Loops](4.%20arrays-slices-loops/README.md)
Working with collections and iteration:
- Arrays (fixed-length) vs Slices (dynamic)
- Creating, appending, and slicing operations
- For loops (traditional, while-style, for-range)
- Break and continue statements
- Common patterns (sum, max, filter, etc.)

### 5. [Pointers](5.%20pointers/README.md)
Understanding memory addresses and pointers:
- What are pointers and why they exist
- The `&` operator (address-of) and `*` operator (dereference)
- Pass by value vs pass by reference
- Pointers with functions
- Nil pointers and safe usage
- When to use (and not use) pointers
- Common mistakes and how to avoid them
- Pointers vs reference types (slices, maps)

### 6. [Maps](6.%20maps/README.md)
Working with key-value pairs and hash tables:
- What are maps and when to use them
- Creating maps (make, literals, nil maps)
- Adding, accessing, updating, and deleting elements
- The comma-ok idiom for checking key existence
- Iterating over maps (random order!)
- Valid key types and value types
- Maps are reference types
- Practical examples (counters, grouping, sets, caches)

### 7. [Custom Types and Methods](7.%20custom-types-methods/README.md)
Creating custom types and attaching behavior:
- Defining structs (custom types)
- Creating and initializing structs
- Receiver functions (methods)
- Value receivers vs pointer receivers
- Constructor functions and embedded structs
- Method chaining and practical examples

### 8. [File I/O](8.%20file-io/README.md)
Reading and writing files to disk:
- Simple file operations (`os.ReadFile`, `os.WriteFile`)
- File handles and proper cleanup
- Buffered I/O for better performance
- File flags and permissions
- Appending to files and line-by-line reading
- Copying files and checking file existence

### 9. [Testing](9.%20testing/README.md)
Writing tests and benchmarks for Go code:
- Understanding `go mod init` and modules
- Test file structure and naming conventions
- Basic tests and assertions (`t.Error`, `t.Fatal`)
- Table-driven tests (best practice pattern)
- Test coverage and running tests
- Benchmarking for performance measurement
- Test helpers and parallel tests
- Why you can't name a module "main"

### 10. [Interfaces](10.%20interfaces/README.md)
Mastering Go's most powerful feature for abstraction:
- What interfaces are and why they're essential
- Implicit interface implementation (no keywords!)
- Polymorphism and writing flexible code
- Interface composition (embedding interfaces)
- Empty interface (`interface{}`) and type assertions
- Type switches for handling different types
- Pointer vs value receivers with interfaces
- Common standard library interfaces (io.Reader, io.Writer, error)
- Best practices (small interfaces, accept interfaces/return structs)
- Real-world examples (payment processing, shapes, counters)

### 11. [Goroutines and Channels](11.%20goroutines-channels/README.md)
Concurrent programming in Go:
- Goroutines (lightweight threads)
- Channels for communication between goroutines
- Buffered vs unbuffered channels
- Channel operations (send, receive, close, range)
- Select statement for multiplexing
- Worker pool pattern
- Common concurrency patterns (fan-out, fan-in, pipeline)
- Best practices for concurrent programming

### 12. [HTTP/REST APIs](12.%20http-rest-apis/README.md)
Building web servers and REST APIs with Go's standard library:
- Creating HTTP servers with `net/http`
- RESTful endpoint design (GET, POST, DELETE)
- JSON marshaling and unmarshaling with struct tags
- Request handling and response writing
- HTTP status codes and error handling
- Middleware patterns (logging, CORS)
- HTTP client for making requests
- Route handling and URL path parsing
- Best practices for API design

## Getting Started with Go

If you're new to Go, start with the Hello World project above. It provides a comprehensive introduction to:
- What packages are
- How imports work
- Basic Go syntax
- Running and building Go programs

## Useful Resources

- [Official Go Documentation](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://go.dev/tour/)

## Common Go Commands

```bash
# Run a Go program
go run filename.go

# Build an executable
go build filename.go

# Format code (Go's automatic formatter)
go fmt filename.go

# Check for common mistakes
go vet filename.go

# Clean build files
go clean

# Initialize a Go module (required for testing and imports)
go mod init <module-name>

# Run tests
go test

# Run tests with verbose output
go test -v

# Run tests with coverage
go test -cover

# Run benchmarks
go test -bench=.
```

## Progress Tracker

- ✅ Hello World - Basic program structure
- ✅ Types and Variables - Data types and declarations
- ✅ Functions and Return Types - Function fundamentals
- ✅ Arrays, Slices, and Loops - Collections and iteration
- ✅ Pointers - Memory addresses and references
- ✅ Maps - Key-value pairs and hash tables
- ✅ Custom Types and Methods - Structs and receiver functions
- ✅ File I/O - Reading and writing files to disk
- ✅ Testing - Writing tests, benchmarks, and using go modules
- ✅ Interfaces - Polymorphism, abstraction, and flexible design
- ✅ Goroutines and Channels - Concurrent programming with goroutines
- ✅ HTTP/REST APIs - Building web servers and REST APIs
- 🔄 More topics coming as I learn...

---

*Last updated: 2026-02-24*
