package interpreter

import (
	"brainfucker/pkg/parser"
	"fmt"
)

type VM struct {
	Tape [30000]byte
	Ptr  int
}

func (vm *VM) Execute(instructions []parser.Instruction) {
	for _, instr := range instructions {
		switch v := instr.(type) {
		case parser.IncrementPtr:
			vm.Ptr = vm.Ptr + v.Count
		case parser.DecrementPtr:
			vm.Ptr = vm.Ptr - v.Count
		case parser.IncrementVal:
			vm.Tape[vm.Ptr] = vm.Tape[vm.Ptr] + byte(v.Value)
		case parser.DecrementVal:
			vm.Tape[vm.Ptr] = vm.Tape[vm.Ptr] - byte(v.Value)
		case parser.Output:
			fmt.Printf("%c", vm.Tape[vm.Ptr])
		case parser.Input:
			var b byte
			fmt.Scanf("%c", &b)
			vm.Tape[vm.Ptr] = b
		case parser.Loop:
			for vm.Tape[vm.Ptr] != 0 {
				vm.Execute(v.Body)
			}
		}
	}
}
