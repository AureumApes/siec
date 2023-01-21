package cmd

import (
	goutils "github.com/AureumApes/goutil"
	"os"
	"strings"
)

func Decompile(fileName string) {
	fileBytes, _ := os.ReadFile(fileName)
	fileContent := strings.Split(goutils.Decrypt([]rune("siec"), string(fileBytes)), "\n")
	instructions := []int64{}
	for i, _ := range fileContent {
		instructions = append(instructions, int64(i+1))
	}
	ExecuteCode(instructions, fileContent)
}
