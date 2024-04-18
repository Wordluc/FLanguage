package Evaluator

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements"
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
	v = slices.Concat(v, []byte("\nEND\n"))
	l, e := Lexer.New(v)
	if e != nil {
		return errors.New("Lexer:" + e.Error())
	}
	p, e := Statements.ParsingStatement(&l, Token.END)
	if e != nil {
		return errors.New("Parser:" + e.Error())
	}

	Eval(p.(*Statements.StatementNode), env)

	return ReplProgram(env)
}

func New() error {
	env := NewEnvironment()
	LoadInnerFunction(env)
	LoadInnerVariable(env)
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
