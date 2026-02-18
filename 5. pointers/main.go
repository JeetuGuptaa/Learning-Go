package main

import "fmt"

func main() {
	fmt.Println("=== Understanding Pointers ===")
	fmt.Println()

	// Example 1: Basic pointer concepts
	basicPointers()

	// Example 2: The & operator (address-of)
	addressOperator()

	// Example 3: The * operator (dereference)
	dereferenceOperator()

	// Example 4: Pointers vs Values
	pointersVsValues()

	// Example 5: Pointers with functions
	pointersWithFunctions()

	// Example 6: Nil pointers
	nilPointers()

	// Example 7: Pointers with slices
	pointersWithSlices()

	// Example 8: Common use cases
	commonUseCases()
}

// Example 1: Basic pointer concepts
func basicPointers() {
	fmt.Println("1. Basic Pointer Concepts:")

	// Regular variable
	age := 25
	fmt.Printf("age value: %d\n", age)
	fmt.Printf("age memory address: %p\n", &age)

	// Pointer variable
	var agePointer *int
	agePointer = &age
	fmt.Printf("agePointer stores address: %p\n", agePointer)
	fmt.Printf("agePointer points to value: %d\n", *agePointer)
	fmt.Println()
}

// Example 2: The & operator (address-of)
func addressOperator() {
	fmt.Println("2. The & Operator (Get Address):")

	name := "Alice"
	count := 42
	price := 19.99

	fmt.Printf("name: %s, address: %p\n", name, &name)
	fmt.Printf("count: %d, address: %p\n", count, &count)
	fmt.Printf("price: %.2f, address: %p\n", price, &price)
	fmt.Println()
}

// Example 3: The * operator (dereference)
func dereferenceOperator() {
	fmt.Println("3. The * Operator (Dereference/Access Value):")

	x := 100
	ptr := &x // ptr is a pointer to x

	fmt.Printf("x = %d\n", x)
	fmt.Printf("ptr = %p (address of x)\n", ptr)
	fmt.Printf("*ptr = %d (value at that address)\n", *ptr)

	// Modify through pointer
	*ptr = 200
	fmt.Printf("After *ptr = 200:\n")
	fmt.Printf("x = %d (x changed!)\n", x)
	fmt.Printf("*ptr = %d\n", *ptr)
	fmt.Println()
}

// Example 4: Pointers vs Values
func pointersVsValues() {
	fmt.Println("4. Pointers vs Values:")

	original := 10
	fmt.Printf("Original value: %d\n", original)

	// Pass by value (copy)
	copyValue := original
	copyValue = 20
	fmt.Printf("After modifying copy: original=%d, copy=%d\n", original, copyValue)

	// Pass by reference (pointer)
	original = 10 // Reset
	pointerToOriginal := &original
	*pointerToOriginal = 30
	fmt.Printf("After modifying via pointer: original=%d\n", original)
	fmt.Println()
}

// Example 5: Pointers with functions
func pointersWithFunctions() {
	fmt.Println("5. Pointers with Functions:")

	balance := 100.0
	fmt.Printf("Initial balance: $%.2f\n", balance)

	// Try to modify without pointer (won't work)
	addMoneyByValue(balance, 50.0)
	fmt.Printf("After addMoneyByValue: $%.2f (unchanged!)\n", balance)

	// Modify with pointer (works!)
	addMoneyByPointer(&balance, 50.0)
	fmt.Printf("After addMoneyByPointer: $%.2f (changed!)\n", balance)
	fmt.Println()
}

func addMoneyByValue(balance float64, amount float64) {
	balance += amount // Only modifies the copy
}

func addMoneyByPointer(balance *float64, amount float64) {
	*balance += amount // Modifies the original
}

// Example 6: Nil pointers
func nilPointers() {
	fmt.Println("6. Nil Pointers:")

	var ptr *int // Declared but not initialized = nil
	fmt.Printf("ptr is nil: %v\n", ptr == nil)

	if ptr == nil {
		fmt.Println("Cannot dereference nil pointer (would crash!)")
	}

	// Initialize it
	value := 42
	ptr = &value
	fmt.Printf("After initialization: ptr=%p, *ptr=%d\n", ptr, *ptr)
	fmt.Println()
}

// Example 7: Pointers with slices
func pointersWithSlices() {
	fmt.Println("7. Pointers with Slices:")

	// Slices are already reference types!
	numbers := []int{1, 2, 3}
	fmt.Printf("Original slice: %v\n", numbers)

	modifySlice(numbers)
	fmt.Printf("After modifySlice: %v (changed!)\n", numbers)

	// Why? Slices internally contain a pointer to the array
	// So you don't need to pass a pointer to a slice
	fmt.Println()
}

func modifySlice(s []int) {
	s[0] = 999 // Modifies original because slices are reference types
}

// Example 8: Common use cases
func commonUseCases() {
	fmt.Println("8. Common Use Cases for Pointers:")

	// Use case 1: Avoid copying large structs
	type LargeStruct struct {
		data [1000]int
	}

	// Helper function for processing pointer
	processByPointer := func(ls *LargeStruct) {
		// Do something with the large struct
		ls.data[0] = 1
	}

	large := LargeStruct{}
	fmt.Printf("Size of LargeStruct: %d bytes\n", 1000*8) // Rough estimate

	// Passing pointer is more efficient than copying
	processByPointer(&large) // Only passes 8 bytes (pointer size)
	fmt.Println("✓ Passed pointer (efficient)")

	// Use case 2: Shared state
	counter := 0
	increment := func() {
		counter++ // Closure can modify outer variable
	}
	increment()
	increment()
	fmt.Printf("Counter after 2 increments: %d\n", counter)

	// Use case 3: Optional values
	var optionalValue *string
	if optionalValue == nil {
		fmt.Println("✓ No value provided (nil pointer pattern)")
	}

	fmt.Println()
}
