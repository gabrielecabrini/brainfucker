package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You have to specify a file path")
		return
	}

	fileContent, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	bracketsMap, err := getBracketsMap(fileContent)
	if err != nil {
		fmt.Println("Error parsing file:", err)
	}

	programIndex := 0
	index := 0
	data := make([]byte, 30000)
	for programIndex < len(fileContent) {
		char := fileContent[programIndex]
		switch char {
		case 62: // >
			if index == len(data)-1 {
				index = 0
			} else {
				index++
			}
		case 60: // <
			if index == 0 {
				index = len(data) - 1
			} else {
				index--
			}
		case 43: // +
			data[index]++
		case 44: // ,
			var input [1]byte
			_, err := os.Stdin.Read(input[:])
			if err == nil {
				data[index] = input[0]
			} else {
				data[index] = 0 // or EOF
			}
		case 45: // -
			data[index]--
		case 46: // .
			fmt.Printf("%c", data[index])
		case 91: // [
			if data[index] == 0 {
				programIndex = bracketsMap[programIndex] // jump to ]
			}
		case 93: // ]
			if data[index] != 0 {
				programIndex = bracketsMap[programIndex] // jump to [
			}
		}
		programIndex++
	}
}

func getBracketsMap(content []byte) (map[int]int, error) {
	var stack []int
	result := make(map[int]int)
	for index, char := range content {
		switch char {
		case 91: // [
			stack = append(stack, index)
		case 93: // ]
			if len(stack) == 0 {
				return nil, errors.New("unmatched brackets")
			}

			openBracketIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result[openBracketIndex] = index
			result[index] = openBracketIndex
		}
	}
	if len(stack) != 0 {
		return nil, errors.New("unmatched brackets")
	}
	return result, nil
}
