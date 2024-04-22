package main

import (
	"FLanguage/Evaluator"
	Lexer "FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements"
	"fmt"
	"os"
)

func lexerFromFile(path string) (Lexer.Lexer, error) {
	file, _ := Lexer.GetByteFromFile(path)
	lexer, e := Lexer.New(file)
	return lexer, e
}

func main() {
	path := os.Args[1]
	l, e := lexerFromFile(path)
	if e != nil {
		fmt.Println(e)
		return
	}

	p, e := Statements.ParsingStatement(&l, Token.END)
	if e != nil {
		fmt.Println(e)
		return
	}
	env := Evaluator.NewEnvironment()
	Evaluator.LoadBuiltInFunction(env)
	Evaluator.LoadBuiltInVariable(env)
	_, e = Evaluator.Eval(p.(*Statements.StatementNode), env)
	if e != nil {
		fmt.Println(e)
		return
	}
	return
}
