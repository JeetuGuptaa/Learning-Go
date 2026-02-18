package main

import "fmt"

func main() {
	fmt.Println("=== Understanding Maps in Go ===")
	fmt.Println()

	// Example 1: Creating and initializing maps
	creatingMaps()

	// Example 2: Adding and accessing elements
	addingAndAccessing()

	// Example 3: Updating and deleting elements
	updatingAndDeleting()

	// Example 4: Checking if key exists
	checkingKeys()

	// Example 5: Iterating over maps
	iteratingMaps()

	// Example 6: Maps with different types
	differentTypes()

	// Example 7: Maps are reference types
	referenceTypes()

	// Example 8: Practical examples
	practicalExamples()
}

// Example 1: Creating and initializing maps
func creatingMaps() {
	fmt.Println("1. Creating and Initializing Maps:")

	// Method 1: Using make
	ages := make(map[string]int)
	ages["Alice"] = 25
	ages["Bob"] = 30
	fmt.Printf("Using make: %v\n", ages)

	// Method 2: Map literal (empty)
	scores := map[string]int{}
	scores["Math"] = 95
	fmt.Printf("Empty literal: %v\n", scores)

	// Method 3: Map literal (with values)
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	fmt.Printf("With values: %v\n", colors)

	// Nil map (not initialized)
	var nilMap map[string]int
	fmt.Printf("Nil map: %v, is nil: %v\n", nilMap, nilMap == nil)

	fmt.Println()
}

// Example 2: Adding and accessing elements
func addingAndAccessing() {
	fmt.Println("2. Adding and Accessing Elements:")

	users := make(map[int]string)

	// Adding elements
	users[1] = "Alice"
	users[2] = "Bob"
	users[3] = "Charlie"

	// Accessing elements
	fmt.Printf("User 1: %s\n", users[1])
	fmt.Printf("User 2: %s\n", users[2])

	// Accessing non-existent key returns zero value
	fmt.Printf("User 999: '%s' (empty string - zero value)\n", users[999])

	fmt.Println()
}

// Example 3: Updating and deleting elements
func updatingAndDeleting() {
	fmt.Println("3. Updating and Deleting Elements:")

	inventory := map[string]int{
		"apples":  10,
		"bananas": 5,
		"oranges": 8,
	}

	fmt.Printf("Original: %v\n", inventory)

	// Updating
	inventory["apples"] = 15
	fmt.Printf("After update: %v\n", inventory)

	// Deleting with delete()
	delete(inventory, "bananas")
	fmt.Printf("After delete: %v\n", inventory)

	// Deleting non-existent key is safe (no error)
	delete(inventory, "grapes")
	fmt.Printf("After deleting non-existent: %v\n", inventory)

	fmt.Println()
}

// Example 4: Checking if key exists
func checkingKeys() {
	fmt.Println("4. Checking if Key Exists:")

	capitals := map[string]string{
		"France": "Paris",
		"Japan":  "Tokyo",
		"Egypt":  "Cairo",
	}

	// The comma-ok idiom
	capital, exists := capitals["France"]
	fmt.Printf("France: %s, exists: %v\n", capital, exists)

	capital, exists = capitals["Germany"]
	fmt.Printf("Germany: %s, exists: %v\n", capital, exists)

	// Common pattern
	if capital, ok := capitals["Japan"]; ok {
		fmt.Printf("✓ Found: Japan's capital is %s\n", capital)
	}

	if _, ok := capitals["Mars"]; !ok {
		fmt.Println("✓ Mars not found in map")
	}

	fmt.Println()
}

// Example 5: Iterating over maps
func iteratingMaps() {
	fmt.Println("5. Iterating Over Maps:")

	grades := map[string]int{
		"Alice":   95,
		"Bob":     87,
		"Charlie": 92,
		"Diana":   88,
	}

	// Iterate over keys and values
	fmt.Println("All grades:")
	for name, grade := range grades {
		fmt.Printf("  %s: %d\n", name, grade)
	}

	// Iterate over keys only
	fmt.Println("\nNames only:")
	for name := range grades {
		fmt.Printf("  %s\n", name)
	}

	// Note: Map iteration order is random!
	fmt.Println("\nIteration order is random - run again to see different order")

	fmt.Println()
}

// Example 6: Maps with different types
func differentTypes() {
	fmt.Println("6. Maps with Different Types:")

	// Map with struct values
	type Person struct {
		age  int
		city string
	}

	people := map[string]Person{
		"Alice": {age: 25, city: "NYC"},
		"Bob":   {age: 30, city: "LA"},
	}
	fmt.Printf("Map with struct values: %v\n", people)

	// Map with slice values
	hobbies := map[string][]string{
		"Alice": {"reading", "hiking"},
		"Bob":   {"gaming", "cooking", "cycling"},
	}
	fmt.Printf("Map with slice values: %v\n", hobbies)

	// Map with map values (nested)
	settings := map[string]map[string]bool{
		"user1": {"darkMode": true, "notifications": false},
		"user2": {"darkMode": false, "notifications": true},
	}
	fmt.Printf("Nested maps: %v\n", settings)

	fmt.Println()
}

// Example 7: Maps are reference types
func referenceTypes() {
	fmt.Println("7. Maps are Reference Types:")

	original := map[string]int{
		"a": 1,
		"b": 2,
	}

	// Assignment creates reference, not copy
	reference := original
	reference["c"] = 3

	fmt.Printf("Original: %v\n", original)
	fmt.Printf("Reference: %v\n", reference)
	fmt.Println("Both changed! Maps are reference types")

	// Passing to function
	modifyMap(original)
	fmt.Printf("After function call: %v\n", original)
	fmt.Println("Function modified original! (no pointer needed)")

	fmt.Println()
}

func modifyMap(m map[string]int) {
	m["d"] = 4
}

// Example 8: Practical examples
func practicalExamples() {
	fmt.Println("8. Practical Examples:")

	// Word counter
	words := []string{"hello", "world", "hello", "go", "world"}
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}
	fmt.Printf("Word count: %v\n", wordCount)

	// Character frequency
	str := "hello"
	charFreq := make(map[rune]int)
	for _, char := range str {
		charFreq[char]++
	}
	fmt.Printf("Character frequency: %v\n", charFreq)

	// Grouping data
	students := []struct {
		name  string
		grade string
	}{
		{"Alice", "A"},
		{"Bob", "B"},
		{"Charlie", "A"},
		{"Diana", "B"},
	}

	byGrade := make(map[string][]string)
	for _, student := range students {
		byGrade[student.grade] = append(byGrade[student.grade], student.name)
	}
	fmt.Printf("Students by grade: %v\n", byGrade)

	// Set implementation (map with bool values)
	uniqueNumbers := map[int]bool{
		1: true,
		2: true,
		3: true,
	}
	fmt.Printf("Set of numbers: %v\n", uniqueNumbers)

	// Check membership
	if uniqueNumbers[2] {
		fmt.Println("✓ 2 is in the set")
	}

	fmt.Println()
}
