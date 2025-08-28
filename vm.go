package main

import "fmt"

type VM struct {
	Tape [30000]byte
	Ptr  int
}

func (vm *VM) Execute(instructions []Instruction) {
	for _, instr := range instructions {
		switch v := instr.(type) {
		case IncrementPtr:
			vm.Ptr++
		case DecrementPtr:
			vm.Ptr--
		case IncrementVal:
			vm.Tape[vm.Ptr]++
		case DecrementVal:
			vm.Tape[vm.Ptr]--
		case Output:
			fmt.Printf("%c", vm.Tape[vm.Ptr])
		case Input:
			var b byte
			fmt.Scanf("%c", &b)
			vm.Tape[vm.Ptr] = b
		case Loop:
			for vm.Tape[vm.Ptr] != 0 {
				vm.Execute(v.Body)
			}
		}
	}
}
