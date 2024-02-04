package words

import (
	"bufio"
	"fmt"
	"os"
)

func Main() {

	//open the file

	test := "test.txt"

	//error handling for file import

	file, err := os.Open(test)
	if err != nil {
		fmt.Println("Error: Cannot open file.")
		return
	}
	defer file.Close()

	//create scanner

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	//initialise the counter

	wordCount := 0

	//initialise the for loop
	for scanner.Scan() {
		wordCount++

	}

	//error handling for file read error

	if err := scanner.Err(); err != nil {
		fmt.Println("Error: Cannot read the file", err)
		return
	}

	//print out the total words in the file
	fmt.Printf("The total number of words in the file is: %d\n", wordCount)

}
