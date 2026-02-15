package main

import (
	"fmt"
	"math"
)

type Person struct {
	firstName string
	lastName  string
	age       int
	email     string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p Person) greet() {
	fmt.Printf("Hello, I'm %s and I'm %d years old.\n", p.fullName(), p.age)
}

func (p *Person) haveBirthday() {
	p.age++
	fmt.Printf("%s is now %d years old!\n", p.fullName(), p.age)
}

func (p *Person) updateEmail(newEmail string) {
	p.email = newEmail
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}

type BankAccount struct {
	owner   string
	balance float64
}

func NewBankAccount(owner string, initialBalance float64) *BankAccount {
	return &BankAccount{
		owner:   owner,
		balance: initialBalance,
	}
}

func (ba *BankAccount) deposit(amount float64) {
	if amount > 0 {
		ba.balance += amount
		fmt.Printf("Deposited $%.2f. New balance: $%.2f\n", amount, ba.balance)
	}
}

func (ba *BankAccount) withdraw(amount float64) bool {
	if amount > 0 && amount <= ba.balance {
		ba.balance -= amount
		fmt.Printf("Withdrew $%.2f. New balance: $%.2f\n", amount, ba.balance)
		return true
	}
	fmt.Println("Insufficient funds or invalid amount")
	return false
}

func (ba BankAccount) getBalance() float64 {
	return ba.balance
}

func (ba BankAccount) displayInfo() {
	fmt.Printf("Account Owner: %s, Balance: $%.2f\n", ba.owner, ba.balance)
}

type Address struct {
	street  string
	city    string
	country string
}

type Employee struct {
	name    string
	age     int
	salary  float64
	address Address
}

func (e Employee) displayDetails() {
	fmt.Printf("Employee: %s, Age: %d, Salary: $%.2f\n", e.name, e.age, e.salary)
	fmt.Printf("Address: %s, %s, %s\n", e.address.street, e.address.city, e.address.country)
}

type Calculator struct {
	result float64
}

func (c *Calculator) add(value float64) *Calculator {
	c.result += value
	return c
}

func (c *Calculator) subtract(value float64) *Calculator {
	c.result -= value
	return c
}

func (c *Calculator) multiply(value float64) *Calculator {
	c.result *= value
	return c
}

func (c *Calculator) divide(value float64) *Calculator {
	if value != 0 {
		c.result /= value
	}
	return c
}

func (c Calculator) getResult() float64 {
	return c.result
}

func main() {
	fmt.Println("=== Custom Types and Receiver Functions ===\n")

	fmt.Println("1. CREATING STRUCTS:")
	var person1 Person
	person1.firstName = "John"
	person1.lastName = "Doe"
	person1.age = 30
	person1.email = "john@example.com"
	fmt.Printf("Person 1: %+v\n", person1)

	person2 := Person{"Jane", "Smith", 25, "jane@example.com"}
	fmt.Printf("Person 2: %+v\n", person2)

	person3 := Person{
		firstName: "Bob",
		lastName:  "Johnson",
		age:       35,
		email:     "bob@example.com",
	}
	fmt.Printf("Person 3: %+v\n", person3)

	fmt.Println("\n2. CALLING METHODS (Value Receiver):")
	fmt.Println("Full name:", person1.fullName())
	person1.greet()
	person2.greet()

	fmt.Println("\n3. POINTER RECEIVER (Modifies Original):")
	fmt.Printf("Before birthday: %s is %d\n", person1.fullName(), person1.age)
	person1.haveBirthday()
	fmt.Printf("After birthday: %s is %d\n", person1.fullName(), person1.age)

	fmt.Println("\n4. UPDATING EMAIL:")
	fmt.Println("Old email:", person1.email)
	person1.updateEmail("john.doe@newmail.com")
	fmt.Println("New email:", person1.email)

	fmt.Println("\n5. RECTANGLE WITH METHODS:")
	rect := Rectangle{width: 10, height: 5}
	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.area())
	fmt.Printf("Perimeter: %.2f\n", rect.perimeter())

	fmt.Println("\n6. CIRCLE WITH METHODS:")
	circle := Circle{radius: 7}
	fmt.Printf("Circle radius: %.2f\n", circle.radius)
	fmt.Printf("Area: %.2f\n", circle.area())
	fmt.Printf("Circumference: %.2f\n", circle.circumference())

	fmt.Println("\n7. USING CONSTRUCTOR FUNCTION:")
	account := NewBankAccount("Alice Brown", 1000.00)
	account.displayInfo()

	fmt.Println("\n8. BANK ACCOUNT OPERATIONS:")
	account.deposit(500.00)
	account.withdraw(200.00)
	account.withdraw(2000.00)
	fmt.Printf("Final balance: $%.2f\n", account.getBalance())

	fmt.Println("\n9. EMBEDDED STRUCTS:")
	emp := Employee{
		name:   "David Wilson",
		age:    28,
		salary: 75000.00,
		address: Address{
			street:  "123 Main St",
			city:    "New York",
			country: "USA",
		},
	}
	emp.displayDetails()

	fmt.Println("\n10. ACCESSING NESTED FIELDS:")
	fmt.Println("Employee city:", emp.address.city)
	emp.address.city = "Los Angeles"
	fmt.Println("Updated city:", emp.address.city)

	fmt.Println("\n11. METHOD CHAINING:")
	calc := Calculator{}
	result := calc.add(10).multiply(2).subtract(5).divide(3).getResult()
	fmt.Printf("Calculation result: %.2f\n", result)

	fmt.Println("\n12. SLICE OF STRUCTS:")
	people := []Person{
		{firstName: "Alice", lastName: "Wonder", age: 25, email: "alice@example.com"},
		{firstName: "Bob", lastName: "Builder", age: 30, email: "bob@example.com"},
		{firstName: "Charlie", lastName: "Chaplin", age: 35, email: "charlie@example.com"},
	}

	fmt.Println("All people:")
	for i, person := range people {
		fmt.Printf("%d. %s (%d years old)\n", i+1, person.fullName(), person.age)
	}

	fmt.Println("\n=== Program Complete ===")
}
