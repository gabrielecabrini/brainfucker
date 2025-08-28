package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You have to specify a file path")
		return
	}

	fileContent, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	instructions, err := ParseSourceBytes(fileContent)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	vm := VM{}
	vm.Execute(instructions)

}
