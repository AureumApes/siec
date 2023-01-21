package internal

func Println(cmd string, vars map[string]string) {
	variable, exists := vars[cmd]
	if !exists {
		println(cmd)
	} else {
		println(variable)
	}
}

func Print(cmd string, vars map[string]string) {
	variable, exists := vars[cmd]
	if !exists {
		print(cmd)
	} else {
		print(variable)
	}
}
