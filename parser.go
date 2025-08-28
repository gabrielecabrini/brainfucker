package main

import "errors"

type Instruction interface{}

type IncrementPtr struct{}
type DecrementPtr struct{}
type IncrementVal struct{}
type DecrementVal struct{}
type Output struct{}
type Input struct{}
type Loop struct {
	Body []Instruction
}

func ParseSourceBytes(src []byte) ([]Instruction, error) {
	stack := [][]Instruction{{}}
	for _, char := range src {
		switch char {
		case 62: // >
			stack[len(stack)-1] = append(stack[len(stack)-1], IncrementPtr{})
		case 60: // <
			stack[len(stack)-1] = append(stack[len(stack)-1], DecrementPtr{})
		case 43: // +
			stack[len(stack)-1] = append(stack[len(stack)-1], IncrementVal{})
		case 44: // ,
			stack[len(stack)-1] = append(stack[len(stack)-1], Input{})
		case 45: // -
			stack[len(stack)-1] = append(stack[len(stack)-1], DecrementVal{})
		case 46: // .
			stack[len(stack)-1] = append(stack[len(stack)-1], Output{})
		case 91: // [
			stack = append(stack, []Instruction{})
		case 93: // ]
			if len(stack) < 2 {
				return nil, errors.New("syntax error, missing brackets")
			}
			body := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = append(stack[len(stack)-1], Loop{body})
		}
	}
	if len(stack) != 1 {
		return nil, errors.New("syntax error, missing brackets")
	}
	return stack[0], nil
}
