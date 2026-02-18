# Pointers in Go

This lesson explains pointers from scratch - what they are, why they exist, and how to use them safely.

## What is a Pointer?

A **pointer** is a variable that stores the **memory address** of another variable.

### Memory Visualization

```
Memory Address    Variable    Value
--------------    --------    -----
0x1000           age         25
0x1008           agePtr      0x1000  (points to age)
```

Think of it like this:
- Regular variable: A box containing a value
- Pointer: A box containing the address of another box

## Why Do Pointers Exist?

### 1. Avoid Copying Large Data
Copying large structs is expensive. Passing a pointer is cheap (just 8 bytes).

### 2. Modify Original Data
Functions normally work on copies. Pointers let you modify the original.

### 3. Share Data
Multiple parts of your program can reference the same data.

### 4. Optional Values
`nil` pointer indicates "no value" (common pattern in Go).

## The Two Pointer Operators

### The `&` Operator: Get Address

The `&` operator gets the **memory address** of a variable.

```go
age := 25
agePointer := &age  // agePointer now holds the address of age

fmt.Println(age)        // 25
fmt.Println(&age)       // 0x1040a124 (memory address)
fmt.Println(agePointer) // 0x1040a124 (same address)
```

**Think of it as:** "Give me the address of this variable"

### The `*` Operator: Two Uses!

The `*` operator has **two different meanings** depending on context:

#### 1. Type Declaration (Creating Pointer Type)

```go
var ptr *int  //ptr is a pointer to an int
var p *string // p is a pointer to a string
```

**Pattern:** `*` before a type means "pointer to that type"

#### 2. Dereferencing (Accessing the Value)

```go
x := 42
ptr := &x
value := *ptr  // Get the value ptr points to

fmt.Println(ptr)   // 0x1040a124 (address)
fmt.Println(*ptr)  // 42 (value at that address)
```

**Think of it as:** "Give me the value at this address"

## Basic Pointer Example

```go
// Step by step pointer usage
age := 25                    // Regular variable
var agePointer *int          // Declare pointer (currently nil)
agePointer = &age            // Store address of age

fmt.Println(age)             // 25
fmt.Println(agePointer)      // 0x1040a124 (address)
fmt.Println(*agePointer)     // 25 (value via pointer)

*agePointer = 30             // Modify via pointer
fmt.Println(age)             // 30 (original changed!)
```

## Pointer Types

Every type can have a pointer version:

```go
var intPtr *int           // Pointer to int
var strPtr *string        // Pointer to string
var boolPtr *bool         // Pointer to bool
var slicePtr *[]int       // Pointer to slice
```

## Pass by Value vs Pass by Reference

### Without Pointers (Pass by Value)

Go passes **copies** by default:

```go
func addMoney(balance float64, amount float64) {
    balance += amount  // Only modifies the COPY
}

func main() {
    balance := 100.0
    addMoney(balance, 50.0)
    fmt.Println(balance)  // Still 100.0 (unchanged!)
}
```

**What happens:**
1. `addMoney` gets a **copy** of balance
2. Copy is modified
3. Original remains unchanged

### With Pointers (Pass by Reference)

```go
func addMoney(balance *float64, amount float64) {
    *balance += amount  // Modifies the ORIGINAL
}

func main() {
    balance := 100.0
    addMoney(&balance, 50.0)
    fmt.Println(balance)  // 150.0 (changed!)
}
```

**What happens:**
1. `addMoney` gets the **address** of balance
2. Uses `*` to modify value at that address
3. Original is modified

## Nil Pointers

A pointer that doesn't point to anything is `nil`:

```go
var ptr *int              // Declared but not initialized = nil
fmt.Println(ptr == nil)   // true

// Dereferencing nil crashes your program!
// value := *ptr          // PANIC! Do NOT do this

// Always check before dereferencing
if ptr != nil {
    value := *ptr
    fmt.Println(value)
}
```

**Safe pattern:**
```go
var ptr *int

// Check before using
if ptr == nil {
    fmt.Println("No value available")
} else {
    fmt.Println("Value:", *ptr)
}
```

## Creating Pointers: Two Ways

### Method 1: Using `&` with Existing Variable

```go
age := 25
ptr := &age  // ptr points to age
```

### Method 2: Using `new()` (Less Common)

```go
ptr := new(int)  // Allocates memory, returns pointer
*ptr = 25        // Set the value
```

Most Go code uses method 1.

## Pointers with Functions: Complete Example

```go
package main

import "fmt"

// Takes value - cannot modify original
func doubleValue(x int) {
    x = x * 2  // Only modifies copy
}

// Takes pointer - can modify original
func doublePointer(x *int) {
    *x = *x * 2  // Modifies original
}

func main() {
    num := 5
    
    doubleValue(num)
    fmt.Println(num)  // Still 5
    
    doublePointer(&num)
    fmt.Println(num)  // Now 10
}
```

## When to Use Pointers

### ✅ Use Pointers When:

1. **Function needs to modify the original value**
   ```go
   func increment(x *int) {
       *x++
   }
   ```

2. **Avoiding expensive copies** (large structs)
   ```go
   type LargeStruct struct {
       data [10000]int
   }
   
   func process(ls *LargeStruct) {
       // Efficient - only passes pointer
   }
   ```

3. **Indicating optional/missing values**
   ```go
   var config *Config  // nil means "no config"
   if config != nil {
       // Use config
   }
   ```

### ❌ Don't Use Pointers When:

1. **Working with small values** (int, bool, small structs)
   ```go
   func add(a int, b int) int {  // ✓ Good (no need for pointers)
       return a + b
   }
   ```

2. **Working with slices or maps** (already reference types)
   ```go
   func appendItem(s []int, item int) {
       s = append(s, item)  // Wait, this won't work! See below
   }
   ```

## Special Case: Slices and Maps

**Slices and maps are already reference types!**

```go
func modifySlice(s []int) {
    s[0] = 999  // ✓ This works! Modifies original
}

func appendSlice(s []int) {
    s = append(s, 100)  // ✗ This won't affect original!
    // Because append might create new underlying array
}

// To append and modify original, return the slice:
func appendSlice(s []int) []int {
    return append(s, 100)
}
```

**Key point:** You rarely need `*[]int` (pointer to slice). Just use `[]int`.

## Common Mistakes and How to Avoid Them

### Mistake 1: Dereferencing Nil Pointer

```go
// ❌ BAD
var ptr *int
value := *ptr  // PANIC!

// ✅ GOOD
var ptr *int
if ptr != nil {
    value := *ptr
}
```

### Mistake 2: Forgetting the `&`

```go
// ❌ BAD
func needsPointer(x *int) {}

num := 5
needsPointer(num)  // ERROR: cannot use int as *int

// ✅ GOOD
needsPointer(&num)
```

### Mistake 3: Forgetting the `*` to Dereference

```go
// ❌ BAD
func double(x *int) {
    x = x * 2  // ERROR: can't multiply pointer
}

// ✅ GOOD
func double(x *int) {
    *x = *x * 2  // Dereference to access value
}
```

## Pointer Syntax Quick Reference

```go
var ptr *int        // Declare pointer to int
ptr = &x            // ptr now points to x
value := *ptr       // Get value ptr points to
*ptr = 42           // Set value ptr points to
ptr == nil          // Check if pointer is nil
```

## Memory Safety

Go is **memory safe**:
- You cannot do pointer arithmetic (unlike C/C++)
- You cannot create invalid pointers
- Garbage collector manages memory automatically

**This means pointers in Go are safe and easy to use!**

## Visual Summary

```go
x := 42              // x stores the value 42
ptr := &x            // ptr stores the address of x
fmt.Println(x)       // 42 (the value)
fmt.Println(&x)      // 0x1040a124 (address of x)
fmt.Println(ptr)     // 0x1040a124 (same address)
fmt.Println(*ptr)    // 42 (value via pointer)

*ptr = 100           // Change value via pointer
fmt.Println(x)       // 100 (x changed!)
```

## Running the Examples

```bash
cd "5. pointers"
go run main.go
```

This will run all 8 examples demonstrating pointer concepts.

## Key Takeaways

1. **Pointers store memory addresses**, not values
2. **`&` gets the address** of a variable
3. **`*` has two meanings:**
   - In type: declares pointer type (`*int`)
   - With variable: dereferences (gets value)
4. **Pass pointers to functions** when you need to modify original data
5. **Always check for `nil`** before dereferencing
6. **Slices and maps don't usually need pointers** (already reference types)
7. Go pointers are **safe** - no arithmetic, automatic memory management

## What's Next?

Now that you understand pointers, you're ready to learn about **custom types and methods** where pointer receivers are commonly used!
