package char

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func Main() {

	test := "test.txt"

	//error handling for both opening and reading the file

	file, err := os.ReadFile(test)
	if err != nil {
		fmt.Println("Error reading file", err)
		os.Exit(1)
	}

	content := string(file)

	charCount := utf8.RuneCountInString(content)

	fmt.Printf("The file contains %d characters \n.", charCount)
}
