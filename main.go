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

func catFiles(files ...string) {
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

	if len(args) == 4 {
		fileName := args[3]
		if args[1] == "head" && args[2] == "-n1" {
			printFirstLine(fileName)
		}
	}

	if len(args) == 3 {
		f1, f2 := args[1], args[2]
		catFiles(f1, f2)
	}

}

// go run main.go test.txt test2.txt
