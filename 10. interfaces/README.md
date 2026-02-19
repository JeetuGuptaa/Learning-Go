# Go Interfaces - Detailed Explanation

## What Are Interfaces?

An **interface** in Go is a type that specifies a set of method signatures (behaviors). Interfaces define *what* an object can do, not *how* it does it. They are one of Go's most powerful features for writing flexible, maintainable, and testable code.

## Key Concepts

### 1. **Interface Definition**

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

- An interface is a collection of method signatures
- Any type that implements all the methods satisfies the interface
- Interfaces are named using nouns (Shape, Writer, Reader) or adjectives ending in "-er" (Stringer, Closer)

### 2. **Implicit Implementation**

Unlike other languages (Java, C#), Go uses **implicit interface satisfaction**:

```go
type Circle struct {
    Radius float64
}

// Circle implements Shape by having Area() and Perimeter() methods
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}
```

**No explicit declaration needed!** If a type has the methods, it implements the interface.

### 3. **Why Use Interfaces?**

#### **Polymorphism**
```go
func printShapeInfo(s Shape) {
    fmt.Printf("Area: %.2f\n", s.Area())
}

// Works with ANY type that implements Shape
printShapeInfo(Circle{Radius: 5})
printShapeInfo(Rectangle{Width: 4, Height: 6})
```

#### **Flexibility**
- Write functions that work with multiple types
- Change implementations without changing interface users
- Easy to extend with new types

#### **Testability**
- Mock interfaces for testing
- Swap real implementations with test doubles

### 4. **Interface Composition**

Combine multiple interfaces:

```go
type Geometry interface {
    Shape      // Embeds Shape interface
    Describer  // Embeds Describer interface
}
```

Any type implementing both Shape and Describer satisfies Geometry.

### 5. **Empty Interface**

`interface{}` (or `any` in Go 1.18+) can hold values of any type:

```go
func printAnything(i interface{}) {
    fmt.Println(i)
}

printAnything(42)
printAnything("hello")
printAnything([]int{1, 2, 3})
```

**Use cases:**
- Generic data structures (before Go 1.18 generics)
- JSON unmarshaling
- Handling unknown types

### 6. **Type Assertions**

Extract the concrete type from an interface:

```go
// Safe type assertion with ok idiom
if circle, ok := shape.(Circle); ok {
    fmt.Println("It's a circle with radius:", circle.Radius)
} else {
    fmt.Println("Not a circle")
}

// Unsafe assertion (panics if wrong type)
circle := shape.(Circle)  // Panics if shape is not Circle
```

### 7. **Type Switches**

Handle different types elegantly:

```go
switch v := shape.(type) {
case Circle:
    fmt.Println("Circle with radius:", v.Radius)
case Rectangle:
    fmt.Println("Rectangle:", v.Width, "x", v.Height)
default:
    fmt.Println("Unknown shape")
}
```

### 8. **Pointer vs Value Receivers**

```go
// Value receiver - works with both values and pointers
func (c Circle) Area() float64 { 
    return math.Pi * c.Radius * c.Radius 
}

// Pointer receiver - requires pointer to satisfy interface
func (ic *IntCounter) Increment() { 
    ic.count++ 
}

// Usage
var counter Counter = &IntCounter{}  // Must use pointer if method has pointer receiver
```

**Rule:** If ANY method has a pointer receiver, you must use a pointer to satisfy the interface.

## Common Go Interfaces

### Standard Library Interfaces

```go
// io.Reader - read data
type Reader interface {
    Read(p []byte) (n int, err error)
}

// io.Writer - write data
type Writer interface {
    Write(p []byte) (n int, err error)
}

// fmt.Stringer - string representation
type Stringer interface {
    String() string
}

// error - error handling
type error interface {
    Error() string
}
```

## Best Practices

### 1. **Keep Interfaces Small**
> "The bigger the interface, the weaker the abstraction." - Rob Pike

```go
// Good - small, focused interface
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Less ideal - large interface
type FileManager interface {
    Read() []byte
    Write([]byte)
    Delete()
    Rename(string)
    Copy(string)
    // ... many more methods
}
```

### 2. **Accept Interfaces, Return Structs**

```go
// Good - accepts interface (flexible)
func ProcessData(r io.Reader) error {
    // ...
}

// Good - returns concrete type (clear contract)
func NewReader() *FileReader {
    return &FileReader{}
}
```

### 3. **Define Interfaces Where Used**

```go
// In consumer package (user of the interface)
package analyzer

type DataFetcher interface {
    Fetch() ([]byte, error)
}

func Analyze(df DataFetcher) {
    data, _ := df.Fetch()
    // analyze data
}
```

### 4. **Use Descriptive Names**

- End with "-er": Reader, Writer, Closer, Stringer
- Describe behavior: Shape, Geometry, PaymentMethod

## Running the Example

```bash
cd "10. interfaces"
go run main.go
```

## Common Pitfalls

### 1. **Nil Interfaces vs Nil Values**

```go
var s Shape              // nil interface
var c *Circle = nil      
s = c                    // non-nil interface holding nil value!

if s == nil {            // false! Interface is not nil
    fmt.Println("nil")
}
```

### 2. **Interface Values Are Not Comparable**

```go
// Compile error if underlying type is not comparable
shapes := []Shape{Circle{}, Rectangle{}}
if shapes[0] == shapes[1] {  // Only works if concrete types are comparable
    // ...
}
```

### 3. **Forgetting Pointer Receivers**

```go
type Counter interface {
    Increment()
}

type MyCounter struct { count int }

func (m *MyCounter) Increment() { m.count++ }

// Error! MyCounter (value) doesn't implement Counter
var c Counter = MyCounter{}      // Won't compile

// Correct! *MyCounter implements Counter
var c Counter = &MyCounter{}     // Works
```

## Real-World Use Cases

1. **Database Abstractions**: Different database drivers (MySQL, PostgreSQL) implement the same interface
2. **HTTP Handlers**: Different handlers implement `http.Handler` interface
3. **Logging**: Different loggers (file, console, remote) implement same logging interface
4. **Payment Processing**: Multiple payment gateways behind unified interface
5. **Testing**: Mock implementations for unit tests

## Summary

- Interfaces define behavior (method signatures)
- Implementation is implicit (no `implements` keyword)
- Use interfaces for flexibility and polymorphism
- Keep interfaces small and focused
- Accept interfaces, return concrete types
- Powerful tool for decoupling and testing

## Next Steps

After mastering interfaces, explore:
- Error handling patterns using error interface
- Context package (context.Context interface)
- Goroutines and channels with interface types
- Generics (Go 1.18+) and when to use them vs interfaces
