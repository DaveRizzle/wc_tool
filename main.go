package main

import (
	"bufio"
	"ccwc/bytes"
	"ccwc/char"
	"ccwc/lines"
	"ccwc/words"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define flags
	var getBytes = flag.Bool("c", false, "Specify this flag to use the bytes module")
	var getLines = flag.Bool("l", false, "Specify this flag to use the lines module")
	var getWords = flag.Bool("w", false, "Specify this flag to use the words module")
	var getChar = flag.Bool("m", false, "Specify this flag to use the char module")

	flag.Parse()

	// Determine if no flags were provided
	noFlagsProvided := !(*getBytes || *getLines || *getWords || *getChar)

	var content []byte
	var err error
	var filename string

	if len(flag.Args()) == 0 {
		fmt.Println("No file specified. Please enter a filename:")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			filename = scanner.Text()
		} else {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading input: %s\n", err)
			} else {
				fmt.Println("No input provided.")
			}
			return
		}
	} else {
		filename = flag.Arg(0)
	}

	// Read file into memory
	content, err = os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	// Process file based on flags
	if noFlagsProvided || *getBytes {
		bytes.ProcessFileStats(filename) // Correctly calling your bytes module functionality
	}
	if noFlagsProvided || *getLines {
		lines.Process(content)
	}
	if noFlagsProvided || *getWords {
		words.Process(content)
	}
	if *getChar {
		char.Process(string(content))
	}
}
