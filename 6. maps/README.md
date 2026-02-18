# Maps in Go

This lesson covers maps - Go's built-in hash table/dictionary data structure for storing key-value pairs.

## What is a Map?

A **map** is an unordered collection of key-value pairs. It's like a dictionary where you look up values by their keys.

**Think of it as:**
- Phone book: name (key) → phone number (value)
- Dictionary: word (key) → definition (value)
- Configuration: setting name (key) → setting value (value)

## Map Basics

### Syntax

```go
map[KeyType]ValueType
```

- `KeyType` - Must be a **comparable** type (can use `==` or `!=`). This includes: int, string, bool, pointer, array, channel, interface types, and structs (if all fields are comparable)
- `ValueType` - Can be any type

**Note:** Slices, maps, and functions cannot be used as keys because they are not comparable.

### Creating Maps

#### Method 1: Using `make()` (Most Common)

```go
ages := make(map[string]int)
ages["Alice"] = 25
ages["Bob"] = 30
```

#### Method 2: Map Literal (Empty)

```go
scores := map[string]int{}
scores["Math"] = 95
```

#### Method 3: Map Literal (With Initial Values)

```go
colors := map[string]string{
    "red":   "#FF0000",
    "green": "#00FF00",
    "blue":  "#0000FF",
}
```

**Note the comma after the last element!** (Required when closing brace is on new line)

#### Nil Maps

```go
var m map[string]int  // Declared but not initialized = nil

// Cannot add to nil map (will panic!)
// m["key"] = 1  // PANIC!

// Must initialize first:
m = make(map[string]int)
m["key"] = 1  // ✓ OK
```

## Basic Operations

### Adding/Updating Elements

```go
users := make(map[int]string)

users[1] = "Alice"     // Add
users[2] = "Bob"       // Add
users[1] = "Alicia"    // Update (same key)
```

### Accessing Elements

```go
name := users[1]       // Get value
fmt.Println(name)      // "Alicia"

// Non-existent key returns zero value
missing := users[999]  // Returns "" (empty string)
```

### Deleting Elements

```go
delete(users, 1)       // Remove key 1

// Deleting non-existent key is safe (no error)
delete(users, 999)     // Does nothing
```

## The Comma-Ok Idiom (Checking if Key Exists)

Accessing a non-existent key returns the zero value, which might be a valid value. Use the comma-ok idiom to check existence:

```go
capitals := map[string]string{
    "France": "Paris",
    "Japan":  "Tokyo",
}

// Two return values: value and bool (exists or not)
capital, exists := capitals["France"]
fmt.Println(capital, exists)  // "Paris" true

capital, exists = capitals["Germany"]
fmt.Println(capital, exists)  // "" false
```

### Common Pattern

```go
if capital, ok := capitals["Japan"]; ok {
    fmt.Println("Found:", capital)
} else {
    fmt.Println("Not found")
}
```

### Ignore the Value

```go
if _, ok := myMap["key"]; ok {
    fmt.Println("Key exists")
}
```

## Iterating Over Maps

Use `for...range` to iterate:

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 28,
}

// Both key and value
for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}

// Keys only
for name := range ages {
    fmt.Println(name)
}

// Values only (use _ for key)
for _, age := range ages {
    fmt.Println(age)
}
```

**Important:** Map iteration order is **random**! Don't rely on any specific order.

## Valid Key Types

Keys must be **comparable** (can use `==` or `!=`):

✅ **Valid key types:**
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`
- `float32`, `float64`
- `string`
- `bool`
- `pointer`
- `struct` (if all fields are comparable)
- `array` (if element type is comparable)

❌ **Invalid key types:**
- `slice`
- `map`
- `function`

```go
// ✅ Valid
ages := map[string]int{}
counts := map[int]bool{}
points := map[[2]int]string{}  // Array as key

// ❌ Invalid
// bad := map[[]int]string{}  // Slice as key - ERROR!
```

## Maps with Different Value Types

### Maps with Struct Values

```go
type Person struct {
    age  int
    city string
}

people := map[string]Person{
    "Alice": {age: 25, city: "NYC"},
    "Bob":   {age: 30, city: "LA"},
}

fmt.Println(people["Alice"].age)  // 25
```

### Maps with Slice Values

```go
hobbies := map[string][]string{
    "Alice": {"reading", "hiking"},
    "Bob":   {"gaming", "cooking"},
}

fmt.Println(hobbies["Alice"][0])  // "reading"
```

### Nested Maps

```go
settings := map[string]map[string]bool{
    "user1": {"darkMode": true, "notifications": false},
    "user2": {"darkMode": false, "notifications": true},
}

fmt.Println(settings["user1"]["darkMode"])  // true
```

## Maps are Reference Types

**Important:** Maps are reference types (like slices). When you assign or pass a map, you're passing a reference, not a copy.

```go
original := map[string]int{"a": 1}
reference := original

reference["b"] = 2

fmt.Println(original)   // map[a:1 b:2]
fmt.Println(reference)  // map[a:1 b:2]
// Both changed!
```

### Maps in Functions

```go
func modify(m map[string]int) {
    m["new"] = 100  // Modifies the original map
}

func main() {
    data := map[string]int{"a": 1}
    modify(data)
    fmt.Println(data)  // map[a:1 new:100]
}
```

**You don't need a pointer to modify a map in a function!**

## Getting Map Length

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
}

fmt.Println(len(ages))  // 2

ages["Carol"] = 28
fmt.Println(len(ages))  // 3
```

## Practical Examples

### 1. Word Counter

```go
text := "hello world hello go"
words := strings.Fields(text)

wordCount := make(map[string]int)
for _, word := range words {
    wordCount[word]++
}

fmt.Println(wordCount)  // map[go:1 hello:2 world:1]
```

### 2. Character Frequency

```go
str := "hello"
charFreq := make(map[rune]int)

for _, char := range str {
    charFreq[char]++
}

fmt.Println(charFreq)  // map[e:1 h:1 l:2 o:1]
```

### 3. Grouping Data

```go
students := []struct {
    name  string
    grade string
}{
    {"Alice", "A"},
    {"Bob", "B"},
    {"Charlie", "A"},
}

byGrade := make(map[string][]string)
for _, student := range students {
    byGrade[student.grade] = append(byGrade[student.grade], student.name)
}

// map[A:[Alice Charlie] B:[Bob]]
```

### 4. Set Implementation

Maps can implement sets (unique collections):

```go
// Set of unique numbers
numbers := map[int]bool{
    1: true,
    2: true,
    3: true,
}

// Add to set
numbers[4] = true

// Check membership
if numbers[2] {
    fmt.Println("2 is in the set")
}

// Remove from set
delete(numbers, 3)
```

### 5. Cache/Memoization

```go
cache := make(map[int]int)

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    
    // Check cache
    if val, ok := cache[n]; ok {
        return val
    }
    
    // Compute and cache
    result := fibonacci(n-1) + fibonacci(n-2)
    cache[n] = result
    return result
}
```

## Common Patterns

### Safe Map Access

```go
// Check before access
if value, ok := myMap[key]; ok {
    // Use value
} else {
    // Handle missing key
}
```

### Initialize Nested Map

```go
data := make(map[string]map[string]int)

// Must initialize inner map before using
if data["user1"] == nil {
    data["user1"] = make(map[string]int)
}
data["user1"]["score"] = 100
```

### Copy a Map

Maps don't have a built-in copy. You must do it manually:

```go
original := map[string]int{"a": 1, "b": 2}
copy := make(map[string]int)

for key, value := range original {
    copy[key] = value
}
```

## Map vs Slice vs Array

| Feature | Array | Slice | Map |
|---------|-------|-------|-----|
| **Size** | Fixed | Dynamic | Dynamic |
| **Access** | By index (0-based) | By index (0-based) | By key |
| **Order** | Ordered | Ordered | Unordered |
| **Type** | Value type | Reference type | Reference type |
| **Use When** | Fixed-size data | Lists, sequences | Key-value pairs, lookups |

## Common Mistakes

### Mistake 1: Using Nil Map

```go
// ❌ BAD
var m map[string]int
m["key"] = 1  // PANIC!

// ✅ GOOD
m := make(map[string]int)
m["key"] = 1
```

### Mistake 2: Not Checking if Key Exists

```go
// ❌ BAD - Can't tell if value is missing or actually 0
count := wordCount["word"]
if count == 0 {
    // Is it missing or is the count actually 0?
}

// ✅ GOOD
if count, ok := wordCount["word"]; ok {
    // Key exists, use count
} else {
    // Key doesn't exist
}
```

### Mistake 3: Modifying Map During Iteration

```go
// ⚠️ CAREFUL - Adding keys during iteration is allowed but unpredictable
m := map[string]int{"a": 1}
for k, v := range m {
    m[k+"new"] = v  // Might or might not iterate over new keys
}

// ❌ BAD - Deleting other keys while iterating
for k := range m {
    delete(m, "other_key")  // Dangerous!
}

// ✅ GOOD - Delete current key is safe
for k := range m {
    delete(m, k)  // Safe to delete current key
}
```

### Mistake 4: Assuming Map Order

```go
// ❌ BAD
m := map[int]string{1: "a", 2: "b", 3: "c"}
for k, v := range m {
    // Don't assume order!
}

// ✅ GOOD - Sort keys if order matters
keys := make([]int, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Ints(keys)

for _, k := range keys {
    fmt.Println(m[k])
}
```

## Performance Characteristics

- **Lookup:** O(1) average case
- **Insert:** O(1) average case
- **Delete:** O(1) average case
- **Space:** O(n)

Maps are very efficient for lookups by key!

## Running the Examples

```bash
cd "6. maps"
go run main.go
```

This will run all 8 examples demonstrating map operations and patterns.

## Key Takeaways

1. **Maps store key-value pairs** with fast lookups
2. **Create with `make()`** or map literals
3. **Keys must be comparable** (int, string, bool, etc.)
4. **Use comma-ok idiom** to check if key exists
5. **Maps are reference types** - no need for pointers
6. **Iteration order is random** - don't rely on it
7. **Nil maps panic when written to** - must initialize first
8. **Use maps for:** lookups, counters, grouping, caches, sets

## What's Next?

Now that you understand maps, you have all the fundamental data structures in Go! Next, you'll learn about **custom types and methods** to build more complex data structures.
