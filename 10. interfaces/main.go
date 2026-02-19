package main

import (
	"fmt"
	"math"
)

// ============================================
// 1. BASIC INTERFACE DEFINITION
// ============================================

// Shape interface defines a contract for any type that can calculate area
// Interfaces in Go are named collections of method signatures
type Shape interface {
	Area() float64
	Perimeter() float64
}

// ============================================
// 2. IMPLEMENTING INTERFACES (IMPLICIT)
// ============================================

// Circle type - implements Shape interface implicitly
type Circle struct {
	Radius float64
}

// Area method for Circle - implementing Shape interface
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter method for Circle - implementing Shape interface
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle type - also implements Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// ============================================
// 3. USING INTERFACES AS PARAMETERS
// ============================================

// printShapeInfo accepts any type that implements the Shape interface
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// ============================================
// 4. INTERFACE WITH MULTIPLE BEHAVIORS
// ============================================

// Describer interface for types that can describe themselves
type Describer interface {
	Describe() string
}

// Making Circle implement Describer interface as well
func (c Circle) Describe() string {
	return fmt.Sprintf("Circle with radius %.2f", c.Radius)
}

// Making Rectangle implement Describer interface
func (r Rectangle) Describe() string {
	return fmt.Sprintf("Rectangle with width %.2f and height %.2f", r.Width, r.Height)
}

// ============================================
// 5. INTERFACE COMPOSITION
// ============================================

// Geometry interface combines multiple interfaces
type Geometry interface {
	Shape
	Describer
}

// printGeometry works with any type implementing both Shape and Describer
func printGeometry(g Geometry) {
	fmt.Println(g.Describe())
	fmt.Printf("  Area: %.2f\n", g.Area())
	fmt.Printf("  Perimeter: %.2f\n", g.Perimeter())
}

// ============================================
// 6. EMPTY INTERFACE (interface{})
// ============================================

// printAnything accepts any type using empty interface
func printAnything(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

// ============================================
// 7. TYPE ASSERTIONS
// ============================================

// getCircleRadius demonstrates type assertion
func getCircleRadius(s Shape) {
	// Type assertion: checks if s is actually a Circle
	if circle, ok := s.(Circle); ok {
		fmt.Printf("This is a circle with radius: %.2f\n", circle.Radius)
	} else {
		fmt.Println("This is not a circle")
	}
}

// ============================================
// 8. TYPE SWITCHES
// ============================================

// describeShape uses type switch to handle different types
func describeShape(s Shape) {
	switch v := s.(type) {
	case Circle:
		fmt.Printf("Type switch found: Circle with radius %.2f\n", v.Radius)
	case Rectangle:
		fmt.Printf("Type switch found: Rectangle (%.2f x %.2f)\n", v.Width, v.Height)
	default:
		fmt.Println("Unknown shape type")
	}
}

// ============================================
// 9. PRACTICAL EXAMPLE: PAYMENT PROCESSING
// ============================================

// PaymentMethod interface for different payment types
type PaymentMethod interface {
	Pay(amount float64) string
}

// CreditCard type
type CreditCard struct {
	CardNumber string
	CardHolder string
}

func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Credit Card ending in %s",
		amount, cc.CardNumber[len(cc.CardNumber)-4:])
}

// PayPal type
type PayPal struct {
	Email string
}

func (pp PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using PayPal account %s", amount, pp.Email)
}

// Cash type
type Cash struct{}

func (c Cash) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f in cash", amount)
}

// processPayment works with any payment method
func processPayment(pm PaymentMethod, amount float64) {
	fmt.Println(pm.Pay(amount))
}

// ============================================
// 10. INTERFACES WITH POINTER RECEIVERS
// ============================================

// Counter interface with mutating method
type Counter interface {
	Increment()
	Value() int
}

// IntCounter with pointer receiver (modifies state)
type IntCounter struct {
	count int
}

// Increment uses pointer receiver because it modifies the counter
func (ic *IntCounter) Increment() {
	ic.count++
}

// Value can use value receiver (just reads)
func (ic *IntCounter) Value() int {
	return ic.count
}

// ============================================
// MAIN FUNCTION - DEMONSTRATING ALL CONCEPTS
// ============================================

func main() {
	fmt.Println("=== Go Interfaces Tutorial ===\n")

	// 1. Basic interface usage
	fmt.Println("1. BASIC INTERFACE USAGE:")
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	printShapeInfo(circle)
	printShapeInfo(rectangle)
	fmt.Println()

	// 2. Storing different types in a slice using interface
	fmt.Println("2. SLICE OF INTERFACES:")
	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 5, Height: 2},
		Circle{Radius: 7},
	}

	totalArea := 0.0
	for i, shape := range shapes {
		fmt.Printf("Shape %d: ", i+1)
		printShapeInfo(shape)
		totalArea += shape.Area()
	}
	fmt.Printf("Total area of all shapes: %.2f\n\n", totalArea)

	// 3. Interface composition
	fmt.Println("3. INTERFACE COMPOSITION:")
	printGeometry(circle)
	printGeometry(rectangle)
	fmt.Println()

	// 4. Empty interface
	fmt.Println("4. EMPTY INTERFACE (interface{}):")
	printAnything(42)
	printAnything("Hello, Go!")
	printAnything(circle)
	printAnything([]int{1, 2, 3})
	fmt.Println()

	// 5. Type assertions
	fmt.Println("5. TYPE ASSERTIONS:")
	getCircleRadius(circle)
	getCircleRadius(rectangle)
	fmt.Println()

	// 6. Type switches
	fmt.Println("6. TYPE SWITCHES:")
	describeShape(circle)
	describeShape(rectangle)
	fmt.Println()

	// 7. Practical example - payment processing
	fmt.Println("7. PRACTICAL EXAMPLE - PAYMENT PROCESSING:")
	creditCard := CreditCard{CardNumber: "1234567890123456", CardHolder: "John Doe"}
	paypal := PayPal{Email: "john@example.com"}
	cash := Cash{}

	processPayment(creditCard, 99.99)
	processPayment(paypal, 49.50)
	processPayment(cash, 25.00)
	fmt.Println()

	// 8. Interfaces with pointer receivers
	fmt.Println("8. POINTER RECEIVERS:")
	counter := &IntCounter{count: 0}
	fmt.Printf("Initial value: %d\n", counter.Value())
	counter.Increment()
	counter.Increment()
	counter.Increment()
	fmt.Printf("After 3 increments: %d\n", counter.Value())
	fmt.Println()

	// 9. Nil interface check
	fmt.Println("9. NIL INTERFACE:")
	var s Shape
	if s == nil {
		fmt.Println("Interface s is nil")
	}
	s = circle
	if s != nil {
		fmt.Printf("Interface s now holds a value: %v\n", s)
	}
}
