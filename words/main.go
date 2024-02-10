package words

import (
	"bufio"
	"fmt"
	"os"
)

func Main() {

	// Open the file for reading
	test := "test.txt"
	file, err := os.Open(test)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err) // Include the error object for more details
		return
	}
	defer file.Close() // Ensure the file is closed after opening

	// Create a new scanner to read the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords) // Set the scanner to split the input into words

	// Initialize the word counter
	wordCount := 0

	// Count words in the file
	for scanner.Scan() {
		wordCount++
	}

	// Check for errors that occurred during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading the file: %v\n", err) // Include the error object for more details
		return
	}

	// Print out the total number of words in the file
	fmt.Printf("The total number of words in the file is: %d\n", wordCount)
}
