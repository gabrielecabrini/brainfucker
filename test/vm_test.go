package test

import (
	"brainfuck-interpreter/interpreter"
	"brainfuck-interpreter/parser"
	"os"
	"testing"
)

func BenchmarkVM(b *testing.B) {
	data, err := os.ReadFile("../examples/fibonacci.b")
	if err != nil {
		b.Fatal(err)
	}

	instrs, err := parser.ParseSourceBytes(data)
	if err != nil {
		b.Fatal(err)
	}

	vm := interpreter.VM{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		vm.Execute(instrs)
	}
}
