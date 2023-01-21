package internal

import "fmt"

func Input(command string, vars map[string]string) map[string]string {
	var input string
	fmt.Scanln(&input)
	vars[command] = input
	return vars
}
