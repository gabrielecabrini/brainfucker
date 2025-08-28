# Brainfuck Interpreter

[Brainfuck](https://wikipedia.org/wiki/Brainfuck) interpreter written in Go. It reads a Brainfuck source file and executes it using a 30,000-cell memory tape.

## Requirements

* Go 1.20+ installed on your system

## Building

Clone the repository and build the executable:

```bash
git clone https://github.com/gabrielecabrini/brainfuck-interpreter
cd brainfuck-interpreter
go build
```

This will generate an executable called `brainfuck-interpreter`.

## Usage

Run the interpreter with a Brainfuck source file as the first argument:

```bash
./brainfuck-interpreter path/to/program.b
```

Example:

```bash
./brainfuck-interpreter examples/hello_world.b
```

## Examples

The `examples` folder contains sample Brainfuck programs, such as:

You can run them using:

```bash
./brainfuck-interpreter examples/fibonacci.b
```
