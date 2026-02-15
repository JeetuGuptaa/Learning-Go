package main

import "fmt"

func main() {
	fmt.Println("=== Go Basic Types and Variables ===\n")

	// 1. STRING TYPE
	fmt.Println("1. STRING TYPE:")
	var name string = "John Doe"
	var city string // Declares variable, default value is "" (empty string)
	city = "New York"
	country := "USA" // Short declaration (type inferred)
	
	fmt.Println("Name:", name)
	fmt.Println("City:", city)
	fmt.Println("Country:", country)

	// 2. INTEGER TYPES
	fmt.Println("\n2. INTEGER TYPES:")
	var age int = 25
	var score int // Default value is 0
	score = 95
	count := 100 // Short declaration
	
	fmt.Println("Age:", age)
	fmt.Println("Score:", score)
	fmt.Println("Count:", count)

	// Different integer sizes
	var smallNum int8 = 127    // -128 to 127
	var bigNum int64 = 9999999 // Much larger range
	fmt.Println("Small number (int8):", smallNum)
	fmt.Println("Big number (int64):", bigNum)

	// 3. FLOATING POINT TYPES
	fmt.Println("\n3. FLOATING POINT TYPES:")
	var price float64 = 19.99
	var temperature float32 = 23.5
	pi := 3.14159 // Default is float64
	
	fmt.Println("Price:", price)
	fmt.Println("Temperature:", temperature)
	fmt.Println("Pi:", pi)

	// 4. BOOLEAN TYPE
	fmt.Println("\n4. BOOLEAN TYPE:")
	var isStudent bool = true
	var hasLicense bool // Default value is false
	isAdult := false
	
	fmt.Println("Is Student:", isStudent)
	fmt.Println("Has License:", hasLicense)
	fmt.Println("Is Adult:", isAdult)

	// 5. MULTIPLE VARIABLE DECLARATION
	fmt.Println("\n5. MULTIPLE DECLARATIONS:")
	var (
		firstName string = "Jane"
		lastName  string = "Smith"
		userAge   int    = 30
	)
	fmt.Println("Full Name:", firstName, lastName)
	fmt.Println("Age:", userAge)

	// Multiple assignment
	x, y, z := 1, 2, 3
	fmt.Println("x, y, z:", x, y, z)

	// 6. CONSTANTS
	fmt.Println("\n6. CONSTANTS:")
	const PI = 3.14159
	const CompanyName = "Tech Corp"
	const MaxUsers = 1000
	
	fmt.Println("PI:", PI)
	fmt.Println("Company:", CompanyName)
	fmt.Println("Max Users:", MaxUsers)

	// 7. TYPE CONVERSION
	fmt.Println("\n7. TYPE CONVERSION:")
	var intValue int = 42
	var floatValue float64 = float64(intValue) // Convert int to float64
	var floatNum float64 = 23.99
	var anotherInt int = int(floatNum)         // Convert float to int (truncates)
	
	fmt.Println("Int Value:", intValue)
	fmt.Println("Converted to Float:", floatValue)
	fmt.Println("23.99 converted to int:", anotherInt)

	// 8. ZERO VALUES (Default values)
	fmt.Println("\n8. ZERO VALUES:")
	var defaultString string
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	
	fmt.Printf("Default string: '%s' (empty)\n", defaultString)
	fmt.Printf("Default int: %d\n", defaultInt)
	fmt.Printf("Default float: %f\n", defaultFloat)
	fmt.Printf("Default bool: %t\n", defaultBool)

	fmt.Println("\n=== Program Complete ===")
}