# Brainfucker

This project contains an interpreter and an LLVM IR transpiler for [Brainfuck](https://wikipedia.org/wiki/Brainfuck).

## Requirements

* Go 1.20+ installed on your system

## Building

Clone the repository and build the executables:

```bash
git clone https://github.com/gabrielecabrini/brainfucker
cd brainfucker
go build -o brainfuck-interpreter cmd/brainfuck-interpreter/main.go
go build -o brainfuck-transpiler cmd/brainfuck-transpiler/main.go
```

This will generate two executables: `brainfuck-interpreter` and `brainfuck-transpiler`.

## Usage

Run the interpreter with a Brainfuck source file:

```bash
./brainfuck-interpreter path/to/program.b
```

Run the transpiler to generate LLVM IR:

```bash
./brainfuck-transpiler path/to/program.b
```

## Examples

The `examples` folder contains sample Brainfuck programs:

```bash
./brainfuck-interpreter examples/hello_world.b
./brainfuck-transpiler examples/fibonacci.b
```
