package cmd

import (
	"github.com/AureumApes/goutil"
	"os"
	"strings"
)

func CompileCode(instructions []int64, code []string) {
	var instructedCode []string
	for _, instruction := range instructions {
		instructedCode = append(instructedCode, code[instruction-1])
	}
	instructedString := strings.Join(instructedCode, "\n")
	os.WriteFile("compiled.scc", []byte(goutil.Encrypt([]rune("siec"), instructedString)), os.ModePerm)
}
