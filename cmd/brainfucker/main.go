package main

import (
	"brainfucker/internal/interpreter"
	"brainfucker/internal/transpiler"
	"brainfucker/pkg/parser"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "brainfucker"}

	rootCmd.AddCommand(&cobra.Command{
		Use:   "run [file]",
		Short: "Run a Brainfuck program with the interpreter",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			instructions, err := parser.ParseAndOptimizeFile(args[0])
			if err != nil {
				return err
			}
			vm := interpreter.VM{}
			vm.Execute(instructions)
			return nil
		},
	})

	rootCmd.AddCommand(&cobra.Command{
		Use:   "transpile [file]",
		Short: "Transpile Brainfuck to LLVM IR",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inputPath := args[0]
			instructions, err := parser.ParseAndOptimizeFile(inputPath)
			if err != nil {
				return err
			}

			llvm := transpiler.LLVMGenerator{}
			llvmIr := llvm.Generate(instructions)

			base := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
			outputPath := filepath.Join(filepath.Dir(inputPath), base+".ll")

			if err := os.WriteFile(outputPath, []byte(llvmIr), 0o644); err != nil {
				return fmt.Errorf("failed to write LLVM IR: %w", err)
			}

			fmt.Println("LLVM IR saved to", outputPath)
			return nil
		},
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
