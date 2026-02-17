package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("=== File I/O Examples ===")
	fmt.Println()

	// Example 1: Writing to a file (simple)
	writeSimpleFile()

	// Example 2: Reading from a file (simple)
	readSimpleFile()

	// Example 3: Writing with buffered writer
	writeBufferedFile()

	// Example 4: Reading with buffered reader
	readBufferedFile()

	// Example 5: Appending to a file
	appendToFile()

	// Example 6: Reading file line by line
	readLineByLine()

	// Example 7: Copying files
	copyFile()

	// Example 8: Checking if file exists
	checkFileExists()
}

// Example 1: Writing to a file (simple)
func writeSimpleFile() {
	fmt.Println("1. Writing to a file (simple):")
	data := []byte("Hello, File I/O!\nThis is a test file.\n")
	err := os.WriteFile("output.txt", data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	fmt.Println("✓ Successfully wrote to output.txt")
	fmt.Println()
}

// Example 2: Reading from a file (simple)
func readSimpleFile() {
	fmt.Println("2. Reading from a file (simple):")
	data, err := os.ReadFile("output.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File contents:")
	fmt.Println(string(data))
}

// Example 3: Writing with buffered writer
func writeBufferedFile() {
	fmt.Println("3. Writing with buffered writer:")
	file, err := os.Create("buffered.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("Line 1: Using buffered writer\n")
	writer.WriteString("Line 2: More efficient for multiple writes\n")
	writer.WriteString("Line 3: Don't forget to flush!\n")
	writer.Flush()
	fmt.Println("✓ Successfully wrote buffered.txt")
	fmt.Println()
}

// Example 4: Reading with buffered reader
func readBufferedFile() {
	fmt.Println("4. Reading with buffered reader:")
	file, err := os.Open("buffered.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	content, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File contents:")
	fmt.Println(string(content))
}

// Example 5: Appending to a file
func appendToFile() {
	fmt.Println("5. Appending to a file:")
	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file for append:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("This line was appended!\n")
	if err != nil {
		fmt.Println("Error appending to file:", err)
		return
	}
	fmt.Println("✓ Successfully appended to output.txt")
	fmt.Println()
}

// Example 6: Reading file line by line
func readLineByLine() {
	fmt.Println("6. Reading file line by line:")
	file, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("Line %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	fmt.Println()
}

// Example 7: Copying files
func copyFile() {
	fmt.Println("7. Copying files:")
	sourceFile, err := os.Open("output.txt")
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	destFile, err := os.Create("output_copy.txt")
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destFile.Close()

	bytesWritten, err := io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}
	fmt.Printf("✓ Copied %d bytes to output_copy.txt\n\n", bytesWritten)
}

// Example 8: Checking if file exists
func checkFileExists() {
	fmt.Println("8. Checking if file exists:")
	files := []string{"output.txt", "nonexistent.txt"}
	for _, filename := range files {
		if _, err := os.Stat(filename); err == nil {
			fmt.Printf("✓ %s exists\n", filename)
		} else if os.IsNotExist(err) {
			fmt.Printf("✗ %s does not exist\n", filename)
		} else {
			fmt.Printf("? Error checking %s: %v\n", filename, err)
		}
	}
}


