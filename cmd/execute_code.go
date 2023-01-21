package cmd

import "github.com/AureumApes/siec-language/internal"

func ExecuteCode(instructions []int64, code []string) {
	var instructedCode []string
	vars := map[string]string{}
	for _, instruction := range instructions {
		instructedCode = append(instructedCode, code[instruction-1])
	}
	for _, command := range instructedCode {
		switch command[0] {
		case '.':
			internal.Print(command[1:], vars)
		case ',':
			internal.Println(command[1:], vars)
		case '^':
			vars = internal.Set(command[1:], vars)
		case '-':
			vars = internal.Subtraction(command[1:], vars)
		}
	}
	return
}
