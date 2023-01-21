package main

import (
	"flag"
	"github.com/AureumApes/siec-language/cmd"
)

func main() {
	run := flag.Bool("run", false, "")
	compile := flag.Bool("compile", true, "")
	flag.Parse()
	instructions := cmd.LoadInstructions()
	code := cmd.LoadCode()
	if *compile {
		cmd.CompileCode(instructions, code)
	}
	if *run {
		cmd.ExecuteCode(instructions, code)
	}
}
