package main

import (
	"brainfucker/internal/interpreter"
	"brainfucker/pkg/parser"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You have to specify a file path")
		return
	}

	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	instructions, err := parser.ParseSourceBytes(bytes)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}
	instructions = parser.Optimize(instructions)

	vm := interpreter.VM{}
	vm.Execute(instructions)

}
