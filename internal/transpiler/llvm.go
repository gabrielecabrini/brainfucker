package transpiler

import (
	"brainfucker/pkg/parser"
	"fmt"
	"strings"
)

type LLVMGenerator struct {
	output       strings.Builder
	labelCounter int
	tempCounter  int
}

func (g *LLVMGenerator) Generate(instructions []parser.Instruction) string {
	g.output.Reset()
	g.labelCounter = 0
	g.tempCounter = 0

	// header
	g.output.WriteString(`
		@tape = dso_local global [30000 x i8] zeroinitializer
		@ptr = dso_local global i32 0

		declare dso_local i32 @putchar(i32 noundef)
		declare dso_local i32 @getchar()

		define dso_local i32 @main() {
	`)

	// generate instructions code
	g.genInstructions(instructions)

	// exit
	g.output.WriteString(`
			ret i32 0
		}
	`)

	return g.output.String()
}

func (g *LLVMGenerator) genInstructions(instructions []parser.Instruction) {
	for _, instr := range instructions {
		switch v := instr.(type) {
		case parser.IncrementPtr:
			g.genIncrementPtr(v)
		case parser.DecrementPtr:
			g.genDecrementPtr(v)
		case parser.IncrementVal:
			g.genIncrementVal(v)
		case parser.DecrementVal:
			g.genDecrementVal(v)
		case parser.Output:
			g.genOutput()
		case parser.Input:
			g.genInput()
		case parser.Loop:
			g.genLoop(v)
		}
	}
}

func (g *LLVMGenerator) newTemp() string {
	g.tempCounter++
	return fmt.Sprintf("%%t%d", g.tempCounter)
}

func (g *LLVMGenerator) newLabel() string {
	g.labelCounter++
	return fmt.Sprintf("L%d", g.labelCounter)
}

func (g *LLVMGenerator) genIncrementPtr(instr parser.IncrementPtr) {
	temp1 := g.newTemp()
	temp2 := g.newTemp()
	g.output.WriteString(fmt.Sprintf(`
		%s = load i32, i32* @ptr
		%s = add i32 %s, %d
		store i32 %s, i32* @ptr
	`, temp1, temp2, temp1, instr.Count, temp2))
}

func (g *LLVMGenerator) genDecrementPtr(instr parser.DecrementPtr) {
	temp1 := g.newTemp()
	temp2 := g.newTemp()
	g.output.WriteString(fmt.Sprintf(`
		%s = load i32, i32* @ptr
		%s = sub i32 %s, %d
		store i32 %s, i32* @ptr
	`, temp1, temp2, temp1, instr.Count, temp2))
}

func (g *LLVMGenerator) genIncrementVal(instr parser.IncrementVal) {
	temp1 := g.newTemp()
	temp2 := g.newTemp()
	temp3 := g.newTemp()
	temp4 := g.newTemp()
	g.output.WriteString(fmt.Sprintf(`
		%s = load i32, i32* @ptr
		%s = getelementptr [30000 x i8], [30000 x i8]* @tape, i32 0, i32 %s
		%s = load i8, i8* %s
		%s = add i8 %s, %d
		store i8 %s, i8* %s
	`, temp1, temp2, temp1, temp3, temp2, temp4, temp3, instr.Value, temp4, temp2))
}

func (g *LLVMGenerator) genDecrementVal(instr parser.DecrementVal) {
	temp1 := g.newTemp()
	temp2 := g.newTemp()
	temp3 := g.newTemp()
	temp4 := g.newTemp()
	g.output.WriteString(fmt.Sprintf(`
		%s = load i32, i32* @ptr
		%s = getelementptr [30000 x i8], [30000 x i8]* @tape, i32 0, i32 %s
		%s = load i8, i8* %s
		%s = sub i8 %s, %d
		store i8 %s, i8* %s
	`, temp1, temp2, temp1, temp3, temp2, temp4, temp3, instr.Value, temp4, temp2))
}

func (g *LLVMGenerator) genOutput() {
	temp1 := g.newTemp()
	temp2 := g.newTemp()
	temp3 := g.newTemp()
	temp4 := g.newTemp()
	g.output.WriteString(fmt.Sprintf(`
		%s = load i32, i32* @ptr
		%s = getelementptr [30000 x i8], [30000 x i8]* @tape, i32 0, i32 %s
		%s = load i8, i8* %s
		%s = zext i8 %s to i32
		call i32 @putchar(i32 %s)
	`, temp1, temp2, temp1, temp3, temp2, temp4, temp3, temp4))
}

func (g *LLVMGenerator) genInput() {
	temp1 := g.newTemp()
	temp2 := g.newTemp()
	temp3 := g.newTemp()
	temp4 := g.newTemp()
	g.output.WriteString(fmt.Sprintf(`
		%s = call i32 @getchar()
		%s = trunc i32 %s to i8
		%s = load i32, i32* @ptr
		%s = getelementptr [30000 x i8], [30000 x i8]* @tape, i32 0, i32 %s
		store i8 %s, i8* %s
	`, temp1, temp2, temp1, temp3, temp4, temp3, temp2, temp4))
}

func (g *LLVMGenerator) genLoop(instr parser.Loop) {
	startLabel := g.newLabel()
	bodyLabel := g.newLabel()
	endLabel := g.newLabel()

	// loop start
	g.output.WriteString(fmt.Sprintf(`
		br label %%%s
	%s:
	`, startLabel, startLabel))

	// loop conditions
	temp1 := g.newTemp()
	temp2 := g.newTemp()
	temp3 := g.newTemp()
	temp4 := g.newTemp()
	g.output.WriteString(fmt.Sprintf(`
		%s = load i32, i32* @ptr
		%s = getelementptr [30000 x i8], [30000 x i8]* @tape, i32 0, i32 %s
		%s = load i8, i8* %s
		%s = icmp ne i8 %s, 0
		br i1 %s, label %%%s, label %%%s
	`, temp1, temp2, temp1, temp3, temp2, temp4, temp3, temp4, bodyLabel, endLabel))

	// loop body
	g.output.WriteString(fmt.Sprintf(`
	%s:
	`, bodyLabel))
	g.genInstructions(instr.Body)
	g.output.WriteString(fmt.Sprintf(`
		br label %%%s
	`, startLabel))

	// loop end
	g.output.WriteString(fmt.Sprintf(`
	%s:
	`, endLabel))
}
