# Testing in Go

This lesson covers testing, benchmarking, and test organization in Go.

## Why Testing Matters

- Ensures code works as expected
- Catches bugs early
- Makes refactoring safer
- Documents expected behavior
- Builds confidence in your code

## Go Modules (`go mod init`)

### What is `go mod`?

Go modules are Go's dependency management system. A module is a collection of related Go packages that are versioned together.

### Why Do You Need `go mod init`?

```bash
go mod init <module-name>
```

**Required for:**
1. **Testing** - `go test` needs a module to run
2. **Importing local packages** - Organize code across multiple files
3. **Dependency management** - Track external libraries
4. **Publishing code** - Share your code with others

**Example:**
```bash
# Initialize a module
go mod init calculator

# Or with a full path (for GitHub projects)
go mod init github.com/username/calculator
```

This creates a `go.mod` file:
```go
module calculator

go 1.23.0
```

### Important Rule

**Never name your module "main"** - `package main` is special (executable programs) and cannot be imported by the test framework.

### Important: Package Names Must Match

**Your test file and source file must use the same package name:**

```go
// calculator.go
package calculator  // ✅

// calculator_test.go
package calculator  // ✅ Same package name
```

**Wrong:**
```go
// calculator.go
package calculator

// calculator_test.go  
package main  // ❌ Package mismatch - tests will fail!
```

### Go Version in go.mod

The `go 1.23.0` line specifies the minimum Go version for your module. You can check your Go version with:

```bash
go version
```

The version in `go.mod` should match or be lower than your installed Go version.

## Test File Structure

### Naming Convention

- Test files must end with `_test.go`
- Example: `calculator.go` → `calculator_test.go`

### Test Function Rules

1. **Must start with "Test"**: `func TestAdd(t *testing.T)`
2. **Must accept `*testing.T`**: For reporting test failures
3. **Must be exported**: Capital first letter

```go
// ✅ Correct
func TestAdd(t *testing.T) { }
func TestMultiply(t *testing.T) { }

// ❌ Wrong
func testAdd(t *testing.T) { }      // lowercase
func AddTest(t *testing.T) { }       // wrong order
func TestAdd() { }                   // no parameter
```

## Basic Test Example

```go
package calculator

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    
    if result != expected {
        t.Errorf("Add(2, 3) = %d; expected %d", result, expected)
    }
}
```

## Testing Methods

### 1. `t.Error()` and `t.Errorf()`
- Reports failure but continues test
- Use for non-critical failures

```go
if result != expected {
    t.Errorf("got %d, expected %d", result, expected)
}
```

### 2. `t.Fatal()` and `t.Fatalf()`
- Reports failure and stops test immediately
- Use when subsequent tests would fail anyway

```go
if config == nil {
    t.Fatal("config cannot be nil")
}
```

### 3. `t.Log()` and `t.Logf()`
- Print information (only shown with `-v` flag)
- Useful for debugging

```go
t.Logf("Testing with input: %d", input)
```

## Table-Driven Tests (Best Practice!)

Instead of writing many similar tests, use a table:

```go
func TestMultiply(t *testing.T) {
    // Define test cases as a slice of structs
    tests := []struct {
        name     string  // Descriptive name for this test case
        a        int     // First input
        b        int     // Second input
        expected int     // Expected result
    }{
        {"positive numbers", 3, 4, 12},  // Each row is a test case
        {"with zero", 5, 0, 0},
        {"negative numbers", -2, 3, -6},
    }
    
    // Run each test case
    for _, tt := range tests {
        // t.Run creates a subtest with the given name
        t.Run(tt.name, func(t *testing.T) {
            result := Multiply(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("got %d, expected %d", result, tt.expected)
            }
        })
    }
}
```

**What the output looks like:**
```
=== RUN   TestMultiply
=== RUN   TestMultiply/positive_numbers
=== RUN   TestMultiply/with_zero
=== RUN   TestMultiply/negative_numbers
--- PASS: TestMultiply (0.00s)
    --- PASS: TestMultiply/positive_numbers (0.00s)
    --- PASS: TestMultiply/with_zero (0.00s)
    --- PASS: TestMultiply/negative_numbers (0.00s)
```

**Benefits:**
- Easy to add new test cases (just add a row!)
- Clear test names with `t.Run()`
- Less code duplication
- Better test output showing which specific case failed

## Understanding `tt := tt` in Parallel Tests

When using parallel tests with table-driven tests, you'll see this pattern:

```go
for _, tt := range tests {
    tt := tt  // ⚠️ Important! Capture the range variable
    t.Run(tt.name, func(t *testing.T) {
        t.Parallel()
        // Use tt here
    })
}
```

**Why is `tt := tt` needed?**

- The loop variable `tt` is reused on each iteration
- When tests run in parallel, they might execute after the loop finishes
- Without capturing, all parallel tests would use the last value of `tt`
- `tt := tt` creates a new variable for each iteration

**Note:** In Go 1.22+, this is no longer required as loop variables are automatically captured per iteration, but it's still good practice for compatibility.

## What Test Failures Look Like

**When a test fails:**

```go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 6 {  // Wrong! Add(2, 3) = 5, not 6
        t.Errorf("Add(2, 3) = %d; expected 6", result)
    }
}
```

**Output:**
```
=== RUN   TestAdd
    calculator_test.go:10: Add(2, 3) = 5; expected 6
--- FAIL: TestAdd (0.00s)
FAIL
```

The output shows:
- Which test failed
- The file and line number
- Your custom error message
- Duration of the test

## Running Tests

### Basic Commands

```bash
# Run all tests in current directory
go test

# Run tests with verbose output
go test -v

# Run specific test
go test -run TestAdd

# Run tests matching pattern
go test -run Test.*Multiply

# Run all tests in package and subpackages
go test ./...
```

### Test Coverage

```bash
# Run tests with coverage
go test -cover

# Generate detailed coverage report
go test -coverprofile=coverage.out

# View coverage in browser
go tool cover -html=coverage.out
```

### Example Output

```
$ go test -v
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestMultiply
=== RUN   TestMultiply/positive_numbers
=== RUN   TestMultiply/with_zero
--- PASS: TestMultiply (0.00s)
    --- PASS: TestMultiply/positive_numbers (0.00s)
    --- PASS: TestMultiply/with_zero (0.00s)
PASS
ok      calculator      0.002s
```

## Benchmarking

Measure performance of your code:

```go
func BenchmarkFibonacci(b *testing.B) {
    // b.N is automatically adjusted by Go
    // It runs the function many times to get accurate timing
    for i := 0; i < b.N; i++ {
        Fibonacci(10)  // Function to benchmark
    }
}
```

**Important:** Benchmark functions:
- Must start with `Benchmark`
- Take `*testing.B` parameter (not `*testing.T`)
- Must loop `b.N` times (Go determines this value)
- Should NOT call `b.Log()` or allocate memory in the loop (affects results)

**Run benchmarks:**
```bash
# Run all benchmarks (excludes regular tests)
go test -bench=.

# Run specific benchmark
go test -bench=BenchmarkFibonacci

# With memory allocation stats (recommended!)
go test -bench=. -benchmem

# Control benchmark duration
go test -bench=. -benchtime=10s
```

**Output:**
```
BenchmarkFibonacci-8    1000000    1234 ns/op    0 B/op    0 allocs/op
```

**Understanding the output:**
- `BenchmarkFibonacci` = benchmark name
- `-8` = number of CPU cores used (GOMAXPROCS)
- `1000000` = number of iterations run (automatically determined)
- `1234 ns/op` = average time per operation in nanoseconds
- `0 B/op` = bytes allocated per operation (with `-benchmem`)
- `0 allocs/op` = number of allocations per operation (with `-benchmem`)

**How b.N works:**
Go automatically increases `b.N` until the benchmark runs long enough to get accurate timing (usually 1 second). Don't set `b.N` yourself!

## Test Helpers

Use `t.Helper()` to create reusable test utilities:

```go
func assertEqual(t *testing.T, got, expected int) {
    t.Helper() // Error points to caller, not this line
    if got != expected {
        t.Errorf("got %d, expected %d", got, expected)
    }
}

func TestWithHelper(t *testing.T) {
    result := Add(5, 7)
    assertEqual(t, result, 12) // Error shows this line if fails
}
```

## Parallel Tests

Run tests concurrently for faster execution:

```go
func TestAddParallel(t *testing.T) {
    t.Parallel() // This test runs in parallel with others
    
    result := Add(1, 1)
    if result != 2 {
        t.Error("failed")
    }
}
```

## Test Setup and Teardown

Use `TestMain` for setup/teardown:

```go
func TestMain(m *testing.M) {
    // Setup
    fmt.Println("Setting up tests...")
    
    // Run tests
    code := m.Run()
    
    // Teardown
    fmt.Println("Cleaning up...")
    
    os.Exit(code)
}
```

## Best Practices

1. **Use table-driven tests** for multiple scenarios
2. **Test edge cases** (zero, negative, empty, nil)
3. **Write clear test names** that describe what's being tested
4. **Keep tests focused** - one concept per test
5. **Use `t.Helper()`** for test utilities
6. **Test behavior, not implementation**
7. **Run tests before committing code**
8. **Aim for high coverage** (but 100% isn't always necessary)

## Common Patterns

### Testing for Errors

```go
func TestDivideByZero(t *testing.T) {
    result, err := SafeDivide(10, 0)
    if err == nil {
        t.Error("expected error for divide by zero")
    }
}
```

### Testing Panics

```go
func TestPanic(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Error("expected panic")
        }
    }()
    
    SomeFunctionThatPanics()
}
```

## Project Structure Example

```
calculator/
├── go.mod              # Module definition
├── calculator.go       # Main code
├── calculator_test.go  # Tests for calculator.go
└── README.md
```

**For larger projects:**
```
myapp/
├── go.mod
├── main.go
├── user.go
├── user_test.go        # Tests for user.go
├── payment.go
├── payment_test.go     # Tests for payment.go
└── utils/
    ├── helpers.go
    └── helpers_test.go # Tests for helpers.go
```

**Test Organization Tips:**
- Keep test files alongside the code they test
- One `_test.go` file per source file
- Group related tests together
- Use descriptive test names: `TestUserCreationWithValidEmail`
- Use table-driven tests to avoid duplicating test code

## Running the Examples

```bash
# Navigate to the testing folder
cd "7. testing"

# Initialize module (already done)
go mod init calculator

# Run all tests
go test

# Run with verbose output
go test -v

# Run with coverage
go test -cover

# Run benchmarks
go test -bench=.
```

## Key Takeaways

1. **`go mod init`** is required before testing
2. **Test files end with `_test.go`**
3. **Test functions start with `Test` and take `*testing.T`**
4. **Use table-driven tests** for cleaner, more maintainable code
5. **`t.Error` continues**, **`t.Fatal` stops**
6. **Benchmarks** measure performance
7. **Coverage** shows what code is tested
8. Testing makes your code reliable and maintainable!

## Additional Resources

- [Official Testing Package Docs](https://pkg.go.dev/testing)
- [Table Driven Tests](https://go.dev/wiki/TableDrivenTests)
- [Go Testing Best Practices](https://go.dev/doc/effective_go#testing)
