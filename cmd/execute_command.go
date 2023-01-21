package cmd

import (
	"github.com/AureumApes/siec/internal"
)

func ExecuteCommand(command string, vars map[string]string) map[string]string {
	switch command[0] {
	case '.':
		internal.Print(command[1:], vars)
	case ',':
		internal.Println(command[1:], vars)
	case '^':
		vars = internal.Set(command[1:], vars)
	case '-':
		vars = internal.Subtraction(command[1:], vars)
	case '+':
		vars = internal.Addition(command[1:], vars)
	case '/':
		vars = internal.Division(command[1:], vars)
	case '*':
		vars = internal.Multiplication(command[1:], vars)
	case '?':
		if ifTrue, ifCode := internal.If(command[1:], vars); ifTrue {
			vars = ExecuteCommand(ifCode, vars)
		}
	case '~':
		vars = internal.Input(command[1:], vars)
	}
	return vars
}
