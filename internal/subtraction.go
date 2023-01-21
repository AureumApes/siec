package internal

import (
	GoUtils "github.com/AureumApes/goutil"
	"strconv"
	"strings"
)

func Subtraction(cmd string, vars map[string]string) map[string]string {
	splitCmd := strings.Split(cmd, "#")
	var1, exists := vars[splitCmd[0]]
	if !exists {
		panic(splitCmd[0] + ": Variable does not exist")
	}
	if !GoUtils.IsNumeric(var1) {
		panic(splitCmd[0] + ": Variable is not numeric")
	}
	var2, exists2 := vars[splitCmd[1]]
	if !exists2 && !GoUtils.IsNumeric(splitCmd[1]) {
		panic(splitCmd[1] + ": is not a variable nor numeric")
	}
	if exists2 {
		var1Num, _ := strconv.ParseInt(var1, 10, 64)
		var2Num, _ := strconv.ParseInt(var2, 10, 64)
		result := var1Num - var2Num
		vars[splitCmd[0]] = strconv.FormatInt(result, 10)
	} else {
		var1Num, _ := strconv.ParseInt(var1, 10, 64)
		var2Num, _ := strconv.ParseInt(splitCmd[1], 10, 64)
		result := var1Num - var2Num
		vars[splitCmd[0]] = strconv.FormatInt(result, 10)
	}
	return vars
}
