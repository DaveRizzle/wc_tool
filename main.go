package main

import (
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

	// Get filename from command-line arguments
	if len(flag.Args()) == 0 {
		fmt.Println("Error: No file specified")
		return
	}
	filename := flag.Arg(0)

	// Read file into memory
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	// Default behavior or specific flags
	if noFlagsProvided || *getBytes {
		bytes.ProcessFileStats(filename)
	}
	if noFlagsProvided || *getLines {
		lines.Process(content)
	}
	if noFlagsProvided || *getWords {
		words.Process(content)
	}
	// The -m (character count) option is not included in the default behavior
	if *getChar {
		char.Process(string(content)) // Ensure char.Process accepts a string
	}
}
