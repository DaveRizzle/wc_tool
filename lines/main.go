package lines

import (
	"bufio"
	"fmt"
	"os"
)

func Main() {
	// Import the file
	test := "test.txt"

	// Open the file for reading
	file, err := os.Open(test)
	if err != nil {
		fmt.Println("Error: Cannot open file", err)
		return
	}
	defer file.Close()

	// Create a scanner to read lines from the file
	scanner := bufio.NewScanner(file)

	// Initialize a line count
	lineCount := 0

	// Read lines and count them
	for scanner.Scan() {
		lineCount++
	}

	// Handle errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: Cannot read file: %v\n", err)
		return
	}

	fmt.Printf("Total lines in the file: %d\n", lineCount)
}
