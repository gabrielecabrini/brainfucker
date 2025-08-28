package parser

type Instruction interface{}

type IncrementPtr struct {
	Count int
}
type DecrementPtr struct {
	Count int
}
type IncrementVal struct {
	Value int
}
type DecrementVal struct {
	Value int
}
type Output struct{}
type Input struct{}
type Loop struct {
	Body []Instruction
}

func ParseSourceBytes(src []byte) ([]Instruction, error) {
	stack := NewInstructionStack()

	for _, char := range src {
		switch char {
		case '>':
			stack.Append(IncrementPtr{Count: 1})
		case '<':
			stack.Append(DecrementPtr{Count: 1})
		case '+':
			stack.Append(IncrementVal{Value: 1})
		case '-':
			stack.Append(DecrementVal{Value: 1})
		case '.':
			stack.Append(Output{})
		case ',':
			stack.Append(Input{})
		case '[':
			stack.Push()
		case ']':
			body, err := stack.Pop()
			if err != nil {
				return nil, err
			}
			stack.Append(Loop{Body: body})
		}
	}

	root, err := stack.Root()
	if err != nil {
		return nil, err
	}

	return Optimize(root), nil
}

func Optimize(instrs []Instruction) []Instruction {
	var optimized []Instruction

	for i := 0; i < len(instrs); i++ {
		switch inst := instrs[i].(type) {
		case IncrementPtr:
			count := inst.Count
			for i+1 < len(instrs) {
				if next, ok := instrs[i+1].(IncrementPtr); ok {
					count += next.Count
					i++
				} else {
					break
				}
			}
			optimized = append(optimized, IncrementPtr{Count: count})

		case DecrementPtr:
			count := inst.Count
			for i+1 < len(instrs) {
				if next, ok := instrs[i+1].(DecrementPtr); ok {
					count += next.Count
					i++
				} else {
					break
				}
			}
			optimized = append(optimized, DecrementPtr{Count: count})

		case IncrementVal:
			val := inst.Value
			for i+1 < len(instrs) {
				if next, ok := instrs[i+1].(IncrementVal); ok {
					val += next.Value
					i++
				} else {
					break
				}
			}
			optimized = append(optimized, IncrementVal{Value: val})

		case DecrementVal:
			val := inst.Value
			for i+1 < len(instrs) {
				if next, ok := instrs[i+1].(DecrementVal); ok {
					val += next.Value
					i++
				} else {
					break
				}
			}
			optimized = append(optimized, DecrementVal{Value: val})

		case Loop:
			optimized = append(optimized, Loop{Body: Optimize(inst.Body)})

		default:
			optimized = append(optimized, inst)
		}
	}

	return optimized
}
