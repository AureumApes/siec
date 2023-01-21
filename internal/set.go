package internal

import "strings"

func Set(cmd string, vars map[string]string) map[string]string {
	splitCmd := strings.Split(cmd, "#")
	vars[splitCmd[0]] = splitCmd[1]
	return vars
}
