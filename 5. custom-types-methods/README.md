# Custom Types and Receiver Functions

This project demonstrates how to create custom types (structs) and attach methods to them using receiver functions in Go.

## What is a Struct?

A struct is a **custom type** that groups together variables (fields) of different types. It's similar to a class in other languages, but simpler.

Think of a struct as a blueprint for creating objects with related data.

## Defining a Struct

```go
type Person struct {
    firstName string
    lastName  string
    age       int
    email     string
}
```

**Components:**
- `type` - Keyword to define a new type
- `Person` - Name of the custom type (use PascalCase)
- `struct` - Keyword indicating this is a struct
- Fields inside `{}` - Variables that belong to this type

## Creating Struct Instances

### Method 1: Declare then Assign
```go
var person Person
person.firstName = "John"
person.lastName = "Doe"
person.age = 30
```

### Method 2: Struct Literal (Ordered)
```go
person := Person{"John", "Doe", 30, "john@example.com"}
```
⚠️ **Not recommended** - Order-dependent, error-prone

### Method 3: Struct Literal (Named Fields)
```go
person := Person{
    firstName: "John",
    lastName:  "Doe",
    age:       30,
    email:     "john@example.com",
}
```
✅ **Recommended** - Clear, order-independent, can omit fields

### Method 4: Pointer to Struct
```go
person := &Person{
    firstName: "John",
    lastName:  "Doe",
    age:       30,
}
```

## Accessing Struct Fields

```go
person := Person{firstName: "John", lastName: "Doe", age: 30}

// Read fields
fmt.Println(person.firstName) // "John"
fmt.Println(person.age)       // 30

// Update fields
person.age = 31
person.email = "john@example.com"
```

## Receiver Functions (Methods)

Methods are functions that are attached to a type. They're defined using a **receiver**.

### Syntax

```go
func (receiver Type) methodName(parameters) returnType {
    // method body
}
```

### Value Receiver

The method receives a **copy** of the struct (cannot modify original):

```go
func (p Person) fullName() string {
    return p.firstName + " " + p.lastName
}

// Call it
person := Person{firstName: "John", lastName: "Doe"}
name := person.fullName() // "John Doe"
```

**When to use:**
- Method only needs to read data
- Struct is small (cheap to copy)
- You want to prevent modifications

### Pointer Receiver

The method receives a **pointer** to the struct (can modify original):

```go
func (p *Person) haveBirthday() {
    p.age++ // Modifies the original struct
}

// Call it
person := Person{firstName: "John", age: 30}
person.haveBirthday() // person.age is now 31
```

**When to use:**
- Method needs to modify the struct
- Struct is large (avoid copying)
- Most common choice for consistency

### Go Automatically Handles Pointers

```go
// Calling pointer receiver method on a value:
person := Person{firstName: "John", age: 30}
person.haveBirthday()  // Go automatically uses &person
(&person).haveBirthday() // Explicit pointer - same result

// Calling value receiver method on a pointer:
personPtr := &Person{firstName: "John", lastName: "Doe"}
personPtr.fullName()  // Go automatically dereferences
(*personPtr).fullName() // Explicit dereference - same result
```

## Value Receiver vs Pointer Receiver

| Aspect | Value Receiver | Pointer Receiver |
|--------|---------------|------------------|
| **Modifies struct?** | No (works on copy) | Yes (works on original) |
| **Memory efficiency** | Less (copies data) | More (passes pointer) |
| **Use when** | Read-only operations | Need to modify struct |
| **Convention** | Small, immutable types | Most methods |

**Best Practice:** Use pointer receivers for all methods on a type for consistency, unless you specifically need value semantics.

## Constructor Functions

Go doesn't have constructors, but by convention, we create functions named `NewTypeName`:

```go
func NewPerson(firstName, lastName string, age int) *Person {
    return &Person{
        firstName: firstName,
        lastName:  lastName,
        age:       age,
        email:     "", // Default value
    }
}

// Usage
person := NewPerson("John", "Doe", 30)
```

**Benefits:**
- Enforces required fields
- Sets default values
- Validates input
- Returns pointer (common pattern)

## Exported vs Unexported Fields

```go
type Person struct {
    Name  string // Exported (public) - starts with uppercase
    age   int    // Unexported (private) - starts with lowercase
}
```

- **Exported** (uppercase) - Can be accessed from other packages
- **Unexported** (lowercase) - Only accessible within same package

## Embedded Structs (Composition)

Go doesn't have inheritance, but you can embed structs:

```go
type Address struct {
    street  string
    city    string
    country string
}

type Person struct {
    name    string
    age     int
    address Address // Embedded struct
}

// Usage
person := Person{
    name: "John",
    age:  30,
    address: Address{
        street:  "123 Main St",
        city:    "New York",
        country: "USA",
    },
}

fmt.Println(person.address.city) // "New York"
```

## Anonymous Structs

Structs without a type name, useful for one-time use:

```go
car := struct {
    brand string
    model string
    year  int
}{
    brand: "Toyota",
    model: "Camry",
    year:  2024,
}

fmt.Println(car.brand) // "Toyota"
```

**Use cases:**
- Temporary data structures
- Configuration objects
- Test data
- JSON unmarshaling when structure is simple

## Method Chaining

Return `*Type` from methods to enable chaining:

```go
type Calculator struct {
    result float64
}

func (c *Calculator) add(value float64) *Calculator {
    c.result += value
    return c // Return self
}

func (c *Calculator) multiply(value float64) *Calculator {
    c.result *= value
    return c
}

func (c *Calculator) getResult() float64 {
    return c.result
}

// Usage
calc := Calculator{}
result := calc.add(10).multiply(2).getResult() // 20
```

## Comparing Structs

Structs are comparable if all fields are comparable:

```go
p1 := Person{firstName: "John", lastName: "Doe", age: 30}
p2 := Person{firstName: "John", lastName: "Doe", age: 30}

p1 == p2 // true (all fields match)
```

**Cannot compare if struct contains:**
- Slices
- Maps
- Functions

## Practical Example: Bank Account

```go
type BankAccount struct {
    owner   string
    balance float64
}

func NewBankAccount(owner string, initial float64) *BankAccount {
    return &BankAccount{
        owner:   owner,
        balance: initial,
    }
}

func (ba *BankAccount) deposit(amount float64) {
    ba.balance += amount
}

func (ba *BankAccount) withdraw(amount float64) bool {
    if amount <= ba.balance {
        ba.balance -= amount
        return true
    }
    return false
}

func (ba BankAccount) getBalance() float64 {
    return ba.balance
}

// Usage
account := NewBankAccount("Alice", 1000)
account.deposit(500)
account.withdraw(200)
fmt.Println(account.getBalance()) // 1300
```

## Common Patterns

### 1. Builder Pattern
```go
type User struct {
    name  string
    email string
    age   int
}

func NewUser(name string) *User {
    return &User{name: name}
}

func (u *User) SetEmail(email string) *User {
    u.email = email
    return u
}

func (u *User) SetAge(age int) *User {
    u.age = age
    return u
}

// Usage
user := NewUser("John").SetEmail("john@example.com").SetAge(30)
```

### 2. Getter/Setter Pattern
```go
type Person struct {
    name string
    age  int
}

// Getter
func (p Person) Name() string {
    return p.name
}

// Setter
func (p *Person) SetName(name string) {
    p.name = name
}
```

### 3. Validation in Constructor
```go
func NewBankAccount(owner string, initial float64) (*BankAccount, error) {
    if initial < 0 {
        return nil, errors.New("initial balance cannot be negative")
    }
    return &BankAccount{owner: owner, balance: initial}, nil
}
```

## Slice of Structs

```go
people := []Person{
    {firstName: "Alice", lastName: "Wonder", age: 25},
    {firstName: "Bob", lastName: "Builder", age: 30},
}

for _, person := range people {
    fmt.Println(person.fullName())
}
```

## Zero Value

Uninitialized struct fields get their type's zero value:

```go
var person Person
// person.firstName = "" (empty string)
// person.age = 0
// person.email = ""
```

## Best Practices

1. **Use pointer receivers** for all methods on a type (consistency)
2. **Constructor functions** for complex initialization
3. **Name methods clearly** - `GetBalance()` not `GB()`
4. **Keep structs small** - Single responsibility principle
5. **Use embedded structs** instead of inheritance
6. **Export fields only when necessary** - Encapsulation
7. **Document exported types** with comments

## Common Mistakes

### ❌ Wrong: Modifying with value receiver
```go
func (p Person) haveBirthday() {
    p.age++ // This modifies the copy, not the original!
}
```

### ✅ Correct: Use pointer receiver
```go
func (p *Person) haveBirthday() {
    p.age++ // Modifies the original
}
```

### ❌ Wrong: Not returning error from constructor
```go
func NewUser(age int) *User {
    return &User{age: age} // No validation
}
```

### ✅ Correct: Return error
```go
func NewUser(age int) (*User, error) {
    if age < 0 {
        return nil, errors.New("age cannot be negative")
    }
    return &User{age: age}, nil
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

1. Create a `Book` struct with title, author, pages, and price
2. Add methods to `Book`: `isLong()` (>300 pages), `discount(percent)`
3. Create a `Library` struct that contains a slice of books
4. Add methods to `Library`: `addBook()`, `findByAuthor()`, `averagePrice()`
5. Create a `Student` struct with embedded `Person`
6. Implement a `Rectangle` with methods for area, perimeter, and `resize()`
7. Create a `TodoList` struct with methods to add, remove, and complete tasks

## Key Takeaways

- **Structs** group related data into custom types
- **Methods** attach behavior to types using receivers
- **Pointer receivers** can modify the struct, value receivers cannot
- **Constructor functions** follow the `NewTypeName` pattern
- **Go doesn't have classes** - use structs and methods instead
- **Composition over inheritance** - use embedded structs
- **Exported fields** start with uppercase, unexported with lowercase
- Methods can be called on both values and pointers (Go handles conversion)

## Next Steps

After mastering custom types and methods, learn about:
- Interfaces (defining behavior)
- Error handling patterns
- Packages and modules
- Pointers in depth
- JSON encoding/decoding with structs

Happy coding! 🚀
