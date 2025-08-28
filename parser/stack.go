package parser

import (
	"errors"
)

type InstructionStack struct {
	frames [][]Instruction
}

func NewInstructionStack() *InstructionStack {
	return &InstructionStack{frames: [][]Instruction{{}}}
}

func (s *InstructionStack) Push() {
	s.frames = append(s.frames, []Instruction{})
}

func (s *InstructionStack) Pop() ([]Instruction, error) {
	if len(s.frames) < 2 {
		return nil, errors.New("syntax error, missing brackets")
	}
	top := s.frames[len(s.frames)-1]
	s.frames = s.frames[:len(s.frames)-1]
	return top, nil
}

func (s *InstructionStack) Append(instr Instruction) {
	s.frames[len(s.frames)-1] = append(s.frames[len(s.frames)-1], instr)
}

func (s *InstructionStack) Current() []Instruction {
	return s.frames[len(s.frames)-1]
}

func (s *InstructionStack) Root() ([]Instruction, error) {
	if len(s.frames) != 1 {
		return nil, errors.New("syntax error, missing brackets")
	}
	return s.frames[0], nil
}
