package bytes

import (
	"fmt"
	"os"
)

func Main() {
	test := "test.txt"

	// Open the file
	file, err := os.Open(test)
	if err != nil {
		fmt.Printf("Error: Cannot open file: %v\n", err)
		return
	}
	defer file.Close()
	fmt.Println("Your file has been opened successfully.")

	// Get file statistics
	fileInfo, err := os.Stat(test)
	if err != nil {
		fmt.Printf("Error: Cannot retrieve file stats: %v\n", err)
		return
	}

	// Print the total bytes in the file
	stats := fileInfo.Size()
	fmt.Printf("Total bytes in the file: %d\n", stats)
}
