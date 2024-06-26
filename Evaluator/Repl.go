package Evaluator

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser"
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
)

func ReplProgram(env *Environment) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-")
	v, _ := reader.ReadBytes('\n')
	if string(v) == "{{\r\n" {
		v = readBlockLines(reader)
	}
	if string(v) == "clear\r\n" {
		env := NewEnvironment()
		LoadBuiltInFunction(env)
		LoadBuiltInVariable(env)
		return ReplProgram(env)
	}
	v = slices.Concat(v, []byte("\nEND\n"))
	l, e := Lexer.New(v)
	if e != nil {
		return errors.New("Lexer:" + e.Error())
	}
	p, e := Parser.ParsingStatement(&l, Token.END)
	if e != nil {
		return errors.New("Parser:" + e.Error())
	}
	_, e = Eval(p.(*Parser.StatementNode), env)
	if e != nil {
		return errors.New("Eval:" + e.Error())
	}
	println()
	return ReplProgram(env)
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
