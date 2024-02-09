package main

import (
	"ccwc/bytes"
	"ccwc/lines"
	"char"
	"flag"
	"words"
)

func main() {
	// Define flags
	var getBytes = flag.Bool("c", false, "Specify this flag to use the bytes module")
	var getLines = flag.Bool("l", false, "Specify this flag to use the lines module")
	var getWords = flag.Bool("w", false, "Specify this flag to use the words module")
	var getChar = flag.Bool("m", false, "Specify this flag to use the char module")

	// Parse flags
	flag.Parse()

	// Switch or if-else to determine which function to call
	switch {
	case *getBytes:
		bytes.Main() // Function from the bytes module
	case *getLines:
		lines.Main() // Function from the bytes module
	case *getWords:
		words.Main() // Function from the bytes module
	case *getChar:
		char.Main() // Function from the bytes module

	}
}
