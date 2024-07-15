package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// go run main.go head -n4 test.txt
func printNLines(fileName string, n int) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		for i := 0; i < n && scanner.Scan(); i++ {
			fmt.Println(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}
}

// go run main.go cat test.txt test2.txt
func catFiles(files []string) {
	for _, file := range files {
		text, err := os.Open(file)
		if err != nil {
			fmt.Println("Error opening file:", err)
		}

		scanner := bufio.NewScanner(text)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Scanner error:", err)
		}

		text.Close()
	}
}

func main() {
	args := os.Args

	if len(args) == 4 && args[1] == "head" && strings.HasPrefix(args[2], "-n") {
		nStr := strings.TrimPrefix(args[2], "-n")
		n, err := strconv.Atoi(nStr)
		if err != nil || n <= 0 {
			fmt.Println("Please enter a valid number of lines:", nStr)
			return
		}
		fileName := args[3]
		printNLines(fileName, n)
	} else if len(args) > 2 && args[1] == "cat" {
		files := args[2:]
		catFiles(files)
	} else {
		fmt.Println("Usage:")
		fmt.Println("  go run main.go head -n1 <filename>")
		fmt.Println("  go run main.go cat <filename1> <filename2> ...")
	}
}
