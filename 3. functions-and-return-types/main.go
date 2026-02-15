package main

import "fmt"

func main() {
	fmt.Println("=== Go Functions and Return Types ===\n")

	// 1. BASIC FUNCTION CALL
	fmt.Println("1. BASIC FUNCTION:")
	greet()

	// 2. FUNCTION WITH PARAMETERS
	fmt.Println("\n2. FUNCTION WITH PARAMETERS:")
	greetPerson("Alice")
	greetPerson("Bob")

	// 3. FUNCTION WITH RETURN VALUE
	fmt.Println("\n3. FUNCTION WITH RETURN VALUE:")
	sum := add(10, 20)
	fmt.Println("10 + 20 =", sum)
	fmt.Println("5 + 7 =", add(5, 7))

	// 4. MULTIPLE PARAMETERS OF SAME TYPE
	fmt.Println("\n4. MULTIPLE PARAMETERS:")
	result := multiply(4, 5)
	fmt.Println("4 * 5 =", result)

	// 5. MULTIPLE RETURN VALUES
	fmt.Println("\n5. MULTIPLE RETURN VALUES:")
	quotient, remainder := divide(17, 5)
	fmt.Printf("17 / 5 = %d remainder %d\n", quotient, remainder)

	// 6. NAMED RETURN VALUES
	fmt.Println("\n6. NAMED RETURN VALUES:")
	area, perimeter := rectangleStats(5, 3)
	fmt.Printf("Rectangle (5x3): Area = %d, Perimeter = %d\n", area, perimeter)

	// 7. IGNORING RETURN VALUES
	fmt.Println("\n7. IGNORING RETURN VALUES:")
	value, _ := divide(20, 3) // Ignore remainder with _
	fmt.Println("Just the quotient:", value)

	// 8. FUNCTION AS VARIABLE
	fmt.Println("\n8. FUNCTION AS VARIABLE:")
	mathFunc := add
	fmt.Println("Using function variable:", mathFunc(15, 25))

	// 9. VARIADIC FUNCTIONS (variable number of arguments)
	fmt.Println("\n9. VARIADIC FUNCTIONS:")
	total1 := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Sum of 1,2,3,4,5:", total1)
	total2 := sumAll(10, 20)
	fmt.Println("Sum of 10,20:", total2)

	// 10. DEFER STATEMENT
	fmt.Println("\n10. DEFER STATEMENT:")
	demoDefer()

	// 11. RECURSION
	fmt.Println("\n11. RECURSION:")
	fmt.Println("Factorial of 5:", factorial(5))
	fmt.Println("Factorial of 6:", factorial(6))

	// 12. ANONYMOUS FUNCTIONS
	fmt.Println("\n12. ANONYMOUS FUNCTIONS:")
	square := func(x int) int {
		return x * x
	}
	fmt.Println("Square of 7:", square(7))

	// Inline anonymous function
	result = func(a, b int) int {
		return a - b
	}(50, 30)
	fmt.Println("50 - 30 =", result)

	fmt.Println("\n=== Program Complete ===")
}

// 1. Basic function with no parameters and no return value
func greet() {
	fmt.Println("Hello, World!")
}

// 2. Function with one parameter
func greetPerson(name string) {
	fmt.Println("Hello,", name + "!")
}

// 3. Function with parameters and return value
func add (a int, b int) int {
	return a + b;
}

// 4. Multiple parameters of the same type (shorthand)
func multiply(x, y int) int {
	return x * y
}

// 5. Multiple return values
func divide(dividend, divisor int) (int, int) {
	quotient := dividend / divisor
	remainder := dividend % divisor
	return quotient, remainder
}

// 6. Named return values
func rectangleStats(length, width int) (area int, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // Naked return - returns named values
}

// 7. Function returning multiple values of different types
func getPersonInfo() (string, int, bool) {
	name := "John Doe"
	age := 30
	isStudent := false
	return name, age, isStudent
}

// 8. Variadic function (accepts variable number of arguments)
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 9. Function demonstrating defer
func demoDefer() {
	defer fmt.Println("This prints last (deferred)")
	fmt.Println("This prints first")
	fmt.Println("This prints second")
	// Deferred function executes just before demoDefer returns
}

// 10. Recursive function
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 11. Function with multiple different parameter types
func formatMessage(name string, age int, score float64) string {
	return fmt.Sprintf("%s is %d years old with score %.2f", name, age, score)
}

// 12. Helper function for string operations
func repeat(text string, times int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += text
	}
	return result
}
