package main

import (
	"flag"
	"github.com/AureumApes/siec-language/cmd"
)

func main() {
	run := flag.Bool("run", true, "")
	// compile := flag.Bool("compile", true, "")
	flag.Parse()
	if *run {
		instructions := cmd.LoadInstructions()
		code := cmd.LoadCode()
		cmd.ExecuteCode(instructions, code)
	}
}
