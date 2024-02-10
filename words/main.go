package words

import (
	"bufio"
	"bytes"
	"fmt"
)

// Process counts and prints the number of words in the provided content bytes.
func Process(content []byte) {
	// Create a scanner to read the content
	scanner := bufio.NewScanner(bytes.NewReader(content))
	scanner.Split(bufio.ScanWords) // Set the scanner to split the input into words

	wordCount := 0
	for scanner.Scan() {
		wordCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error processing the content: %v\n", err)
		return
	}

	fmt.Printf("The total number of words in the content is: %d\n", wordCount)
}
