package parser

import (
	"reflect"
	"testing"
)

func TestParseSourceBytes_Simple(t *testing.T) {
	src := []byte("++>-.,")

	instrs, err := ParseSourceBytes(src)
	if err != nil {
		t.Fatalf("ParseSourceBytes returned error: %v", err)
	}

	expected := []Instruction{
		IncrementVal{Value: 1},
		IncrementVal{Value: 1},
		IncrementPtr{Count: 1},
		DecrementVal{Value: 1},
		Output{},
		Input{},
	}

	if !reflect.DeepEqual(instrs, expected) {
		t.Errorf("expected %+v, got %+v", expected, instrs)
	}
}

func TestParseSourceBytes_Loop(t *testing.T) {
	src := []byte("[+-]")

	instrs, err := ParseSourceBytes(src)
	if err != nil {
		t.Fatalf("ParseSourceBytes returned error: %v", err)
	}

	// should be a loop with body {IncrementVal, DecrementVal}
	expected := []Instruction{
		Loop{Body: []Instruction{
			IncrementVal{Value: 1},
			DecrementVal{Value: 1},
		}},
	}

	if !reflect.DeepEqual(instrs, expected) {
		t.Errorf("expected %+v, got %+v", expected, instrs)
	}
}

func TestOptimize_Merges(t *testing.T) {
	src := []Instruction{
		IncrementVal{Value: 1},
		IncrementVal{Value: 1},
		IncrementVal{Value: 1},
		DecrementPtr{Count: 1},
		DecrementPtr{Count: 1},
	}

	opt := Optimize(src)

	expected := []Instruction{
		IncrementVal{Value: 3}, // merged
		DecrementPtr{Count: 2}, // merged
	}

	if !reflect.DeepEqual(opt, expected) {
		t.Errorf("expected %+v, got %+v", expected, opt)
	}
}

func TestOptimize_LoopBody(t *testing.T) {
	loop := Loop{Body: []Instruction{
		IncrementVal{Value: 1},
		IncrementVal{Value: 1},
	}}

	opt := Optimize([]Instruction{loop})

	expected := []Instruction{
		Loop{Body: []Instruction{
			IncrementVal{Value: 2}, // merged inside loop
		}},
	}

	if !reflect.DeepEqual(opt, expected) {
		t.Errorf("expected %+v, got %+v", expected, opt)
	}
}
