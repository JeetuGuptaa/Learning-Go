package main

import "fmt"

func main() {
	fmt.Println("=== Arrays, Slices, and Loops in Go ===\n")

	// 1. ARRAYS - Fixed Length
	fmt.Println("1. ARRAYS (Fixed Length):")
	var numbers [5]int // Array of 5 integers
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println("Array:", numbers)
	fmt.Println("First element:", numbers[0])
	fmt.Println("Array length:", len(numbers))

	// Array with initialization
	colors := [3]string{"Red", "Green", "Blue"}
	fmt.Println("Colors array:", colors)

	// Array with ... (compiler counts)
	fruits := [...]string{"Apple", "Banana", "Orange", "Mango"}
	fmt.Println("Fruits array:", fruits)
	fmt.Println("Number of fruits:", len(fruits))

	// 2. SLICES - Dynamic Arrays
	fmt.Println("\n2. SLICES (Dynamic Length):")
	
	// Create slice using make
	scores := make([]int, 3) // Slice with length 3
	scores[0] = 85
	scores[1] = 90
	scores[2] = 95
	fmt.Println("Scores slice:", scores)
	fmt.Println("Length:", len(scores))
	fmt.Println("Capacity:", cap(scores))

	// Slice literal
	cities := []string{"New York", "London", "Tokyo"}
	fmt.Println("Cities:", cities)

	// 3. APPENDING TO SLICES
	fmt.Println("\n3. APPENDING TO SLICES:")
	var names []string // nil slice
	fmt.Println("Empty slice:", names)
	
	names = append(names, "Alice")
	names = append(names, "Bob")
	names = append(names, "Charlie")
	fmt.Println("After appending:", names)
	
	// Append multiple values
	names = append(names, "David", "Eve", "Frank")
	fmt.Println("After appending more:", names)

	// 4. SLICING OPERATIONS
	fmt.Println("\n4. SLICING OPERATIONS:")
	nums := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Original:", nums)
	fmt.Println("nums[2:5]:", nums[2:5])   // Elements 2,3,4
	fmt.Println("nums[:4]:", nums[:4])     // First 4 elements
	fmt.Println("nums[5:]:", nums[5:])     // From index 5 to end
	fmt.Println("nums[:]:", nums[:])       // All elements

	// 5. FOR LOOP - Traditional Style
	fmt.Println("\n5. FOR LOOP - Traditional Style:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	// 6. FOR LOOP - While Style
	fmt.Println("\n6. FOR LOOP - While Style:")
	count := 0
	for count < 3 {
		fmt.Printf("Count: %d\n", count)
		count++
	}

	// 7. FOR LOOP - Infinite Loop (with break)
	fmt.Println("\n7. FOR LOOP - Infinite Loop with Break:")
	counter := 0
	for {
		if counter >= 3 {
			break
		}
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}

	// 8. FOR RANGE - Iterate Over Slice
	fmt.Println("\n8. FOR RANGE - Iterate Over Slice:")
	languages := []string{"Go", "Python", "JavaScript", "Rust"}
	
	// With index and value
	for index, language := range languages {
		fmt.Printf("%d: %s\n", index, language)
	}

	// Only value (ignore index)
	fmt.Println("\nOnly values:")
	for _, language := range languages {
		fmt.Printf("Language: %s\n", language)
	}

	// Only index (ignore value)
	fmt.Println("\nOnly indices:")
	for index := range languages {
		fmt.Printf("Index: %d\n", index)
	}

	// 9. FOR RANGE - Iterate Over Array
	fmt.Println("\n9. FOR RANGE - Iterate Over Array:")
	days := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	for i, day := range days {
		fmt.Printf("Day %d: %s\n", i+1, day)
	}

	// 10. CONTINUE STATEMENT
	fmt.Println("\n10. CONTINUE Statement (Skip Even Numbers):")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 11. NESTED LOOPS
	fmt.Println("\n11. NESTED LOOPS (Multiplication Table):")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d x %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}

	// 12. COPYING SLICES
	fmt.Println("\n12. COPYING SLICES:")
	original := []int{1, 2, 3, 4, 5}
	copied := make([]int, len(original))
	copy(copied, original)
	
	fmt.Println("Original:", original)
	fmt.Println("Copied:", copied)
	
	copied[0] = 999
	fmt.Println("After modifying copy:")
	fmt.Println("Original:", original)
	fmt.Println("Copied:", copied)

	// 13. 2D SLICES (Slice of Slices)
	fmt.Println("\n13. 2D SLICES:")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	
	fmt.Println("Matrix:")
	for i, row := range matrix {
		fmt.Printf("Row %d: %v\n", i, row)
	}

	// 14. REMOVING ELEMENTS FROM SLICE
	fmt.Println("\n14. REMOVING ELEMENTS FROM SLICE:")
	numbers2 := []int{10, 20, 30, 40, 50}
	fmt.Println("Original:", numbers2)
	
	// Remove element at index 2 (value 30)
	indexToRemove := 2
	numbers2 = append(numbers2[:indexToRemove], numbers2[indexToRemove+1:]...)
	fmt.Println("After removing index 2:", numbers2)

	// 15. PRACTICAL EXAMPLE - Sum and Average
	fmt.Println("\n15. PRACTICAL EXAMPLE - Sum and Average:")
	grades := []float64{85.5, 92.0, 78.5, 90.0, 88.5}
	
	sum := 0.0
	for _, grade := range grades {
		sum += grade
	}
	average := sum / float64(len(grades))
	
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Sum: %.2f\n", sum)
	fmt.Printf("Average: %.2f\n", average)

	// 16. PRACTICAL EXAMPLE - Finding Max Value
	fmt.Println("\n16. PRACTICAL EXAMPLE - Finding Maximum:")
	values := []int{23, 67, 12, 89, 45, 34}
	
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	
	fmt.Printf("Values: %v\n", values)
	fmt.Printf("Maximum value: %d\n", max)

	// 17. PRACTICAL EXAMPLE - Filtering
	fmt.Println("\n17. PRACTICAL EXAMPLE - Filtering Even Numbers:")
	allNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evenNumbers []int
	
	for _, num := range allNumbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	
	fmt.Printf("All numbers: %v\n", allNumbers)
	fmt.Printf("Even numbers: %v\n", evenNumbers)

	fmt.Println("\n=== Program Complete ===")
}
