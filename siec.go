package main

import (
	"flag"
	"github.com/AureumApes/siec/cmd"
)

func main() {
	run := flag.Bool("run", false, "")
	build := flag.Bool("build", true, "")
	runbuild := flag.String("runbuild", "", "")
	flag.Parse()
	instructions := cmd.LoadInstructions()
	code := cmd.LoadCode()
	if *runbuild != "" {
		cmd.Decompile(*runbuild)
	}
	if *build {
		cmd.CompileCode(instructions, code)
	}
	if *run {
		cmd.ExecuteCode(instructions, code)
	}
}
