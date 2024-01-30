package bytes

import (
	"fmt"
	"os"
)

func Main() {
	//import the file
	test := "test.txt"

	//Handle open file errors
	file, err := os.Open(test)
	if err != nil {
		fmt.Println("Error: Cannot open file", err)
		return
	}
	defer file.Close()

	//Handle read file errors

	if err != nil {
		fmt.Println("Error: Cannot read file", err)
	}

	fmt.Println("Your file has been imported successfully.")

	fileInfo, err := os.Stat(test)
	if err != nil {
		fmt.Println("There was an error processing your file.")
		return
	}

	stats := fileInfo.Size()
	fmt.Printf("Total bytes in the file: %d\n", stats)

}
