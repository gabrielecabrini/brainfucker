package main

import (
	"brainfuck-interpreter/interpreter"
	"brainfuck-interpreter/parser"
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

	vm := interpreter.VM{}
	vm.Execute(instructions)

}
