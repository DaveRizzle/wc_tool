package lines

import (
	"bufio"
	"fmt"
	"os"
)

func Main() {
	// Import the file
	test := "test.txt"

	// Handle open file errors
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

	if err := scanner.Err(); err != nil {
		fmt.Println("Error: Cannot read file", err)
		return
	}

	fmt.Printf("Total lines in the file: %d\n", lineCount)
}
