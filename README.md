# Brainfucker

This project contains an interpreter and an LLVM IR transpiler for [Brainfuck](https://wikipedia.org/wiki/Brainfuck).

## Requirements

* Go 1.20+ installed on your system

## Building

Clone the repository and build the executables:

```bash
git clone https://github.com/gabrielecabrini/brainfucker
cd brainfucker
go build -o brainfucker cmd/brainfucker/main.go
```

## Usage

Run the interpreter with a Brainfuck source file:

```bash
./brainfucker run path/to/program.b
```

Run the transpiler to generate LLVM IR:

```bash
./brainfucker transpile path/to/program.b
```

## Examples

The `examples` folder contains sample Brainfuck programs
