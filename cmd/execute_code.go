package cmd

func ExecuteCode(instructions []int64, code []string) {
	var instructedCode []string
	vars := map[string]string{}
	for _, instruction := range instructions {
		instructedCode = append(instructedCode, code[instruction-1])
	}
	for _, command := range instructedCode {
		vars = ExecuteCommand(command, vars)
	}
}
