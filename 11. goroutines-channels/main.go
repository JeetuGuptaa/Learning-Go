package main

import (
	"fmt"
	"time"
)

// Basic goroutine example
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Number: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	letters := []string{"A", "B", "C", "D", "E"}
	for _, letter := range letters {
		fmt.Printf("Letter: %s\n", letter)
		time.Sleep(150 * time.Millisecond)
	}
}

// Channel examples
func sendData(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // Send data to channel
		fmt.Printf("Sent: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Close channel when done
}

func receiveData(ch chan int) {
	for num := range ch { // Receive until channel is closed
		fmt.Printf("Received: %d\n", num)
	}
}

// Buffered channel example
func bufferedChannelExample() {
	fmt.Println("\n--- Buffered Channel Example ---")
	ch := make(chan string, 2) // Buffered channel with capacity of 2

	ch <- "first"
	ch <- "second"
	// ch <- "third" // This would block since buffer is full

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// Worker pool pattern
func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(100 * time.Millisecond)
		results <- job * 2
	}
}

func workerPoolExample() {
	fmt.Println("\n--- Worker Pool Example ---")
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}

// Select statement example
func selectExample() {
	fmt.Println("\n--- Select Statement Example ---")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}

func main() {
	fmt.Println("=== Goroutines and Channels ===")
	fmt.Println()

	// Example 1: Basic goroutines
	fmt.Println("--- Basic Goroutines ---")
	go printNumbers()
	go printLetters()
	time.Sleep(1 * time.Second) // Wait for goroutines to finish

	// Example 2: Channels
	fmt.Println("\n--- Channels ---")
	ch := make(chan int)
	go sendData(ch)
	receiveData(ch)

	// Example 3: Buffered channels
	bufferedChannelExample()

	// Example 4: Worker pool
	workerPoolExample()

	// Example 5: Select statement
	selectExample()

	fmt.Println("\nAll examples completed!")
}
