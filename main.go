package main

import (
	"bufio"
	"fmt"
	"os"
)

// go run main.go head -n1 test.txt
func printFirstLine(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		firstLine := scanner.Text()
		fmt.Println("\n" + firstLine + "\n")
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

	if len(args) == 4 && args[1] == "head" && args[2] == "-n1" {
		fileName := args[3]
		printFirstLine(fileName)
	} else if len(args) > 2 && args[1] == "cat" {
		files := args[2:]
		catFiles(files)
	} else {
		fmt.Println("Usage:")
		fmt.Println("  go run main.go head -n1 <filename>")
		fmt.Println("  go run main.go cat <filename1> <filename2> ...")
	}
}
