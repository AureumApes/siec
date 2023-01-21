package cmd

import (
	"errors"
	"os"
	"strings"
)

func LoadCode() []string {
	codeBytes, err := os.ReadFile("code.sc")
	if errors.Is(err, os.ErrNotExist) {
		panic("Code file (code.sc) does not exist")
	}
	codeString := string(codeBytes)
	code := strings.Split(codeString, "\n")
	return code
}
