package main

import (
	"brainfucker/internal/transpiler"
	"brainfucker/pkg/parser"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You have to specify a file path")
		return
	}

	inputPath := os.Args[1]

	bytes, err := os.ReadFile(inputPath)
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

	llvm := transpiler.LLVMGenerator{}
	llvmIr := llvm.Generate(instructions)

	// change extension
	outputPath := filepath.Base(os.Args[1])
	outputPath = filepath.Join(filepath.Dir(inputPath),
		filepath.Base(inputPath[:len(inputPath)-len(filepath.Ext(inputPath))]+".ll"))

	err = os.WriteFile(outputPath, []byte(llvmIr), 0644)
	if err != nil {
		fmt.Println("Error writing LLVM IR to file:", err)
		return
	}

	fmt.Println("LLVM IR saved to", outputPath)
}
