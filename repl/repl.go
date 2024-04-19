package main

import (
	"FLanguage/Evaluator"
	"bufio"
	"fmt"
	"slices"
)

func main() {
	env := Evaluator.NewEnvironment()
	Evaluator.LoadBuiltInFunction(env)
	Evaluator.LoadBuiltInVariable(env)
	for {
		fmt.Println(Evaluator.ReplProgram(env))
	}
}
func readBlockLines(reader *bufio.Reader) []byte {
	var text []byte
	for {
		v, _ := reader.ReadBytes('\n')
		if string(v) == "}}\r\n" {
			return text
		}
		text = slices.Concat(text, v)
	}
}
