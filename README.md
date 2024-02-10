CCWC - Custom Word Count Tool
CCWC is a custom implementation of the classic Unix wc command, designed to count words, lines, characters, and bytes in text files. This tool adheres to the Unix philosophy of writing simple, modular software that does one thing well and can be easily integrated with other programs.

Features
Byte Count: Count the number of bytes in a file with the -c option.
Line Count: Count the number of lines in a file with the -l option.
Word Count: Count the number of words in a file with the -w option.
Character Count: Count the number of characters in a file with the -m option, with support for multibyte characters depending on the locale.
Default Behavior: When no option is provided, CCWC outputs the line, word, and byte counts.
Standard Input: CCWC can read from standard input if no filename is specified, allowing for flexible usage in command pipelines.
Usage
bash
Copy code
ccwc [options] [file]
If no file is specified, CCWC reads from standard input.

Building CCWC
To compile CCWC:

bash
Copy code
go build -o ccwc
Examples
Count bytes in test.txt:

bash
Copy code
ccwc -c test.txt
Count lines in test.txt:

bash
Copy code
ccwc -l test.txt
Count words in test.txt:

bash
Copy code
ccwc -w test.txt
Acknowledgments
This project is inspired by a coding challenge created by John Crickett, focusing on the development of Unix-like command-line tools following the Unix philosophy. The challenge encourages the creation of simple, efficient tools that perform text processing tasks.

For more information on the challenge and other similar projects, visit Coding Challenges.
