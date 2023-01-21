package cmd

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func LoadInstructions() []int64 {
	instructionBytes, err := os.ReadFile("inst.sc")
	if errors.Is(err, os.ErrNotExist) {
		panic("Instruction file (inst.sc) does not exist")
	}
	instructionNoNewln := strings.ReplaceAll(string(instructionBytes), "\n", "")
	instructionStrings := make([]string, len([]rune(string(instructionBytes)))/2)
	instructions := make([]int64, len(instructionStrings))
	for i := 0; i < len([]rune(string(instructionBytes)))/2; i++ {
		instructionStrings[i] = instructionNoNewln[2*i : i*2+2]
	}
	for i, instruction := range instructionStrings {
		instructions[i], _ = strconv.ParseInt(instruction, 20, 64)
	}
	return instructions
}
