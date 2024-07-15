package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// go run main.go head -n4 test.txt -n -b
// printNLines prints the first n lines of a file, optionally numbering the lines and/or blank lines
func printNLines(fileName string, n int, numberLines bool, showBlanks bool) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for i := 0; i < n && scanner.Scan(); i++ {
		if numberLines && !showBlanks {
			fmt.Printf("%d: %s\n", i+1, scanner.Text())
		} else if numberLines && showBlanks {

			if len(scanner.Text()) > 0 {
				fmt.Printf("%d: %s\n", i+1, scanner.Text())
			} else {
				i--
			}
		} else {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}
}

// go run main.go cat test.txt test2.txt
// catFiles prints the content of multiple files
func catFiles(files []string) {
	for _, file := range files {
		if err := catFile(file); err != nil {
			fmt.Println("Error processing file:", file, err)
		}
	}
}

// catFile prints the content of a single file
func catFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func main() {
	args := os.Args

	if len(args) >= 4 && args[1] == "head" {
		nStr := ""
		showNumberLines := false
		showBlanks := false

		// Parse arguments to get the number of lines and the flag for numbering lines
		if strings.HasPrefix(args[2], "-n") {
			nStr = strings.TrimPrefix(args[2], "-n")
			if len(args) > 4 && args[4] == "-n" {
				showNumberLines = true
			}
		}

		if args[len(args)-1] == "-b" {
			showBlanks = true
		}

		n, err := strconv.Atoi(nStr)
		if err != nil || n <= 0 {
			fmt.Println("Please enter a valid number of lines:", nStr)
			return
		}
		fileName := args[3]
		printNLines(fileName, n, showNumberLines, showBlanks)
	} else if len(args) > 2 && args[1] == "cat" {
		files := args[2:]
		catFiles(files)
	} else {
		fmt.Println("Usage:")
		fmt.Println("  go run main.go head -n1 <filename>")
		fmt.Println("  go run main.go cat <filename1> <filename2> ...")
	}
}
