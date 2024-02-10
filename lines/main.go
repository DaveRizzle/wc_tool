package lines

import (
	"bufio"
	"bytes"
	"fmt"
)

// Process counts and prints the number of lines in the provided content bytes.
func Process(content []byte) {
	// Convert the []byte content to a Reader to use with bufio.Scanner
	reader := bytes.NewReader(content)
	scanner := bufio.NewScanner(reader)

	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error: Cannot read lines: %v\n", err)
		return
	}

	fmt.Printf("Total lines in the file: %d\n", lineCount)
}
