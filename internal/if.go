package internal

import (
	"strings"
)

func If(command string, vars map[string]string) (bool, string) {
	splited := strings.Split(command[2:], "#")
	var1, exists := vars[splited[0]]
	if !exists {
		panic(splited[0] + ": Variable does not exist")
	}
	switch command[0:2] {
	case "!=":
		if var1 != splited[1] {
			return true, splited[2]
		}
	case "==":
		if var1 == splited[1] {
			return true, splited[2]
		}
	}
	return false, ""
}
