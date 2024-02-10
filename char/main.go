package char

import (
	"fmt"
	"unicode/utf8"
)

// Process counts and prints the number of characters in the provided content string.
func Process(content string) {
	charCount := utf8.RuneCountInString(content)
	fmt.Printf("The file contains %d characters.\n", charCount)
}
