package internal

import (
	goutil "github.com/AureumApes/goutil"
	"strconv"
	"strings"
)

func Subtraction(cmd string, vars map[string]string) map[string]string {
	splitCmd := strings.Split(cmd, "#")
	var1, exists := vars[splitCmd[0]]
	if !exists {
		panic(splitCmd[0] + ": Variable does not exist")
	}
	if !goutil.IsNumeric(var1) {
		panic(splitCmd[0] + ": Variable is not numeric")
	}
	var2, exists2 := vars[splitCmd[1]]
	if !exists2 && !goutil.IsNumeric(splitCmd[1]) {
		panic(splitCmd[1] + ": is not a variable nor numeric")
	}
	if exists2 {
		var1Num, _ := strconv.ParseFloat(var1, 10)
		var2Num, _ := strconv.ParseFloat(var2, 10)
		result := var1Num - var2Num
		vars[splitCmd[0]] = strconv.FormatFloat(result, 'f', -1, 64)
	} else {
		var1Num, _ := strconv.ParseFloat(var1, 10)
		var2Num, _ := strconv.ParseFloat(splitCmd[1], 10)
		result := var1Num - var2Num
		vars[splitCmd[0]] = strconv.FormatFloat(result, 'f', -1, 64)
	}
	return vars
}

func Addition(cmd string, vars map[string]string) map[string]string {
	splitCmd := strings.Split(cmd, "#")
	var1, exists := vars[splitCmd[0]]
	if !exists {
		panic(splitCmd[0] + ": Variable does not exist")
	}
	if !goutil.IsNumeric(var1) {
		panic(splitCmd[0] + ": Variable is not numeric")
	}
	var2, exists2 := vars[splitCmd[1]]
	if !exists2 && !goutil.IsNumeric(splitCmd[1]) {
		panic(splitCmd[1] + ": is not a variable nor numeric")
	}
	if exists2 {
		var1Num, _ := strconv.ParseFloat(var1, 10)
		var2Num, _ := strconv.ParseFloat(var2, 10)
		result := var1Num + var2Num
		vars[splitCmd[0]] = strconv.FormatFloat(result, 'f', -1, 64)
	} else {
		var1Num, _ := strconv.ParseFloat(var1, 10)
		var2Num, _ := strconv.ParseFloat(splitCmd[1], 10)
		result := var1Num + var2Num
		vars[splitCmd[0]] = strconv.FormatFloat(result, 'f', -1, 64)
	}
	return vars
}

func Multiplication(cmd string, vars map[string]string) map[string]string {
	splitCmd := strings.Split(cmd, "#")
	var1, exists := vars[splitCmd[0]]
	if !exists {
		panic(splitCmd[0] + ": Variable does not exist")
	}
	if !goutil.IsNumeric(var1) {
		panic(splitCmd[0] + ": Variable is not numeric")
	}
	var2, exists2 := vars[splitCmd[1]]
	if !exists2 && !goutil.IsNumeric(splitCmd[1]) {
		panic(splitCmd[1] + ": is not a variable nor numeric")
	}
	if exists2 {
		var1Num, _ := strconv.ParseFloat(var1, 10)
		var2Num, _ := strconv.ParseFloat(var2, 10)
		result := var1Num * var2Num
		vars[splitCmd[0]] = strconv.FormatFloat(result, 'f', -1, 64)
	} else {
		var1Num, _ := strconv.ParseFloat(var1, 10)
		var2Num, _ := strconv.ParseFloat(splitCmd[1], 10)
		result := var1Num * var2Num
		vars[splitCmd[0]] = strconv.FormatFloat(result, 'f', -1, 64)
	}
	return vars
}

func Division(cmd string, vars map[string]string) map[string]string {
	splitCmd := strings.Split(cmd, "#")
	var1, exists := vars[splitCmd[0]]
	if !exists {
		panic(splitCmd[0] + ": Variable does not exist")
	}
	if !goutil.IsNumeric(var1) {
		panic(splitCmd[0] + ": Variable is not numeric")
	}
	var2, exists2 := vars[splitCmd[1]]
	if !exists2 && !goutil.IsNumeric(splitCmd[1]) {
		panic(splitCmd[1] + ": is not a variable nor numeric")
	}
	if exists2 {
		var1Num, _ := strconv.ParseFloat(var1, 64)
		var2Num, _ := strconv.ParseFloat(var2, 10)
		result := var1Num / var2Num
		vars[splitCmd[0]] = strconv.FormatFloat(result, 'f', -1, 64)
	} else {
		var1Num, _ := strconv.ParseFloat(var1, 10)
		var2Num, _ := strconv.ParseFloat(splitCmd[1], 10)
		result := var1Num / var2Num
		vars[splitCmd[0]] = strconv.FormatFloat(result, 'f', -1, 64)
	}
	return vars
}
