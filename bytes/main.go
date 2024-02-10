package bytes

import (
	"fmt"
	"os"
)

// This is a specialized function within the bytes module for getting file size, distinct from processing the content.
func ProcessFileStats(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: Cannot open file: %v\n", err)
		return
	}
	defer file.Close()

	// Use file.Stat() to get the fileInfo.
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Error: Cannot retrieve file stats: %v\n", err)
		return
	}

	// Print the total bytes in the file
	fmt.Printf("Total bytes in the file: %d\n", fileInfo.Size())
}
