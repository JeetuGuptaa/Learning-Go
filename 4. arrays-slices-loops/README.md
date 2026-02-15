# Arrays, Slices, and Loops in Go

This project covers three fundamental concepts in Go: arrays (fixed-size collections), slices (dynamic collections), and loops (iteration).

## Arrays

Arrays are **fixed-length** sequences of elements of the same type.

### Declaring Arrays

```go
// Method 1: Declare and assign later
var numbers [5]int
numbers[0] = 10
numbers[1] = 20

// Method 2: Declare with values
colors := [3]string{"Red", "Green", "Blue"}

// Method 3: Let compiler count (using ...)
fruits := [...]string{"Apple", "Banana", "Orange"}
```

### Key Points About Arrays

- **Fixed size** - Length is part of the type (`[3]int` ≠ `[5]int`)
- **Zero-indexed** - First element is at index 0
- **Zero values** - Unassigned elements get zero value (0 for int, "" for string)
- **Length** - Use `len(array)` to get length
- **Cannot grow** - Size cannot change after declaration

### Accessing Array Elements

```go
fruits := [3]string{"Apple", "Banana", "Orange"}

fmt.Println(fruits[0])  // "Apple" (first element)
fmt.Println(fruits[2])  // "Orange" (last element)
fmt.Println(len(fruits)) // 3 (length)

fruits[1] = "Mango" // Modify element
```

### Arrays Are Rarely Used in Go

Arrays have fixed sizes, which is limiting. **Slices** are almost always preferred!

## Slices

Slices are **dynamic-length** sequences built on top of arrays. They're the most common way to work with collections in Go.

### Creating Slices

```go
// Method 1: Slice literal
numbers := []int{1, 2, 3, 4, 5}

// Method 2: Using make (creates slice with length and capacity)
scores := make([]int, 5)    // length 5, capacity 5
names := make([]string, 0, 10) // length 0, capacity 10

// Method 3: nil slice (uninitialized)
var items []string // nil slice, length and capacity are 0
```

### Slice vs Array Syntax

```go
array := [3]int{1, 2, 3}  // Array (fixed size, has number in brackets)
slice := []int{1, 2, 3}   // Slice (dynamic, no number in brackets)
```

### Length vs Capacity

- **Length** - Number of elements in the slice: `len(slice)`
- **Capacity** - Number of elements in underlying array: `cap(slice)`

```go
s := make([]int, 3, 5)
fmt.Println(len(s)) // 3 (current number of elements)
fmt.Println(cap(s)) // 5 (space available)
```

### Appending to Slices

```go
var numbers []int
numbers = append(numbers, 10)        // Add one element
numbers = append(numbers, 20, 30, 40) // Add multiple elements

// Append another slice
more := []int{50, 60}
numbers = append(numbers, more...) // Use ... to unpack slice
```

**Important:** `append` may create a new underlying array if capacity is exceeded!

### Slicing Operations

Extract portions of slices or arrays:

```go
nums := []int{10, 20, 30, 40, 50, 60}

nums[2:5]  // [30 40 50] - indices 2,3,4
nums[:3]   // [10 20 30] - first 3 elements  
nums[3:]   // [40 50 60] - from index 3 to end
nums[:]    // [10 20 30 40 50 60] - all elements
```

**Format:** `slice[start:end]` (includes start, excludes end)

### Modifying Slices

```go
numbers := []int{1, 2, 3, 4, 5}

// Modify element
numbers[2] = 999

// Add element
numbers = append(numbers, 6)

// Remove element at index 2
numbers = append(numbers[:2], numbers[3:]...)
```

### Copying Slices

```go
original := []int{1, 2, 3, 4, 5}
copied := make([]int, len(original))
copy(copied, original)

// Now modifications to copied won't affect original
copied[0] = 999
```

**Important:** Simple assignment creates a reference, not a copy!

```go
slice1 := []int{1, 2, 3}
slice2 := slice1 // slice2 references same underlying array
slice2[0] = 999  // This also changes slice1[0]!
```

### Multi-Dimensional Slices

```go
// 2D slice (slice of slices)
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

fmt.Println(matrix[0][1]) // Access row 0, column 1 (value: 2)
```

## For Loops

Go has only one loop keyword: `for`. But it's very flexible!

### 1. Traditional For Loop

```go
for initialization; condition; post {
    // loop body
}

// Example
for i := 0; i < 5; i++ {
    fmt.Println(i) // Prints 0, 1, 2, 3, 4
}
```

**Components:**
- **Initialization** - Runs once before loop (`i := 0`)
- **Condition** - Checked before each iteration (`i < 5`)
- **Post** - Runs after each iteration (`i++`)

### 2. While-Style Loop

```go
count := 0
for count < 5 {
    fmt.Println(count)
    count++
}
```

Just omit initialization and post statement!

### 3. Infinite Loop

```go
for {
    // Runs forever (until break)
    if someCondition {
        break
    }
}
```

### 4. For-Range Loop

The most common way to iterate over slices and arrays:

```go
names := []string{"Alice", "Bob", "Charlie"}

// With index and value
for index, name := range names {
    fmt.Printf("%d: %s\n", index, name)
}

// Only value (ignore index)
for _, name := range names {
    fmt.Println(name)
}

// Only index (ignore value)
for index := range names {
    fmt.Println(index)
}
```

**Key Points:**
- `range` returns two values: index and value
- Use `_` to ignore values you don't need
- Works with arrays, slices, strings, maps, and channels

### Break Statement

Exits the loop immediately:

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // Exit loop when i equals 5
    }
    fmt.Println(i) // Prints 0, 1, 2, 3, 4
}
```

### Continue Statement

Skips the current iteration and moves to the next:

```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue // Skip when i equals 3
    }
    fmt.Println(i) // Prints 1, 2, 4, 5
}
```

### Nested Loops

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        fmt.Printf("(%d,%d) ", i, j)
    }
    fmt.Println()
}
// Output:
// (1,1) (1,2) (1,3) 
// (2,1) (2,2) (2,3) 
// (3,1) (3,2) (3,3)
```

## Common Patterns

### 1. Sum of Elements

```go
numbers := []int{10, 20, 30, 40, 50}
sum := 0
for _, num := range numbers {
    sum += num
}
fmt.Println("Sum:", sum) // 150
```

### 2. Find Maximum

```go
numbers := []int{23, 67, 12, 89, 45}
max := numbers[0]
for _, num := range numbers {
    if num > max {
        max = num
    }
}
fmt.Println("Max:", max) // 89
```

### 3. Filter Elements

```go
numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
var evens []int
for _, num := range numbers {
    if num%2 == 0 {
        evens = append(evens, num)
    }
}
fmt.Println(evens) // [2 4 6 8 10]
```

### 4. Reverse a Slice

```go
numbers := []int{1, 2, 3, 4, 5}
for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
    numbers[i], numbers[j] = numbers[j], numbers[i]
}
fmt.Println(numbers) // [5 4 3 2 1]
```

### 5. Check if Element Exists

```go
names := []string{"Alice", "Bob", "Charlie"}
searchName := "Bob"
found := false

for _, name := range names {
    if name == searchName {
        found = true
        break
    }
}
fmt.Println("Found:", found) // true
```

## Array vs Slice Quick Reference

| Feature | Array | Slice |
|---------|-------|-------|
| **Size** | Fixed | Dynamic |
| **Declaration** | `[5]int` | `[]int` |
| **Can grow** | No | Yes (with `append`) |
| **Passed to functions** | By value (copied) | By reference |
| **Use case** | Rare | Very common |

## Common Mistakes

### ❌ Wrong: Index out of bounds
```go
numbers := []int{1, 2, 3}
fmt.Println(numbers[3]) // Error! Valid indices: 0,1,2
```

### ✅ Correct: Check length first
```go
numbers := []int{1, 2, 3}
if len(numbers) > 3 {
    fmt.Println(numbers[3])
}
```

### ❌ Wrong: Modifying slice during iteration
```go
numbers := []int{1, 2, 3, 4, 5}
for i := range numbers {
    numbers = append(numbers, i) // Dangerous!
}
```

### ✅ Correct: Iterate over a copy or use traditional loop
```go
numbers := []int{1, 2, 3, 4, 5}
length := len(numbers)
for i := 0; i < length; i++ {
    numbers = append(numbers, i)
}
```

### ❌ Wrong: Not updating slice after append
```go
numbers := []int{1, 2, 3}
append(numbers, 4) // This doesn't work!
fmt.Println(numbers) // Still [1 2 3]
```

### ✅ Correct: Assign result of append
```go
numbers := []int{1, 2, 3}
numbers = append(numbers, 4) // Correct!
fmt.Println(numbers) // [1 2 3 4]
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

1. Create a slice of your favorite foods and print each one
2. Write a function that returns the average of a slice of numbers
3. Create a function to find the minimum value in a slice
4. Reverse a slice of strings
5. Remove all duplicate values from a slice
6. Create a 2D slice representing a tic-tac-toe board
7. Count how many times a specific value appears in a slice
8. Merge two sorted slices into one sorted slice

## Key Takeaways

- **Arrays** have fixed length, rarely used in Go
- **Slices** are dynamic, built on arrays, very commonly used
- **Always assign** the result of `append()` back to the slice
- **`for range`** is the idiomatic way to loop over collections
- **`len()`** returns the number of elements
- **`cap()`** returns the capacity of underlying array
- **Slices** share underlying arrays - use `copy()` for independent copies
- **Break** exits the loop, **continue** skips to next iteration

## Next Steps

After mastering arrays, slices, and loops, learn about:
- Maps (key-value pairs)
- Structs (custom data types)
- Pointers
- Error handling
- Packages and modules

Happy coding! 🚀
