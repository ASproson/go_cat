package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	// data, err := os.ReadFile("test.txt")
	// if(err != nil){
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(string(data))
	// }
	if(len(os.Args) != 4){
		fmt.Println("Usage: go run main.go head -n1 [file]")
	}

	if os.Args[2] != "-n1" {
		fmt.Println("Usage: go run main.go head -n1 [file]")
	}

	fileName := os.Args[3]
	fmt.Println(fileName)

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if scanner.Scan() {
		firstLine := scanner.Text()
		fmt.Print(firstLine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}


}

// go run main.go head -n1 test.txt 