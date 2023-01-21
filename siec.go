package main

import "github.com/AureumApes/siec-language/cmd"

func main() {
	instructions := cmd.LoadInstructions()
	code := cmd.LoadCode()
	cmd.ExecuteCode(instructions, code)
}
