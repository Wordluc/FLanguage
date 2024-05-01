package main

import (
	"FLanguage/Evaluator"
	Lexer "FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser"
	"FLanguage/repl"
	"fmt"
	"os"
)

func lexerFromFile(path string) (Lexer.Lexer, error) {
	file, _ := Lexer.GetByteFromFile(path)
	lexer, e := Lexer.New(file)
	return lexer, e
}

func main() {
	typeOp := os.Args[1]
	if typeOp == "r" {
		repl.Start()
		return
	}
	path := os.Args[2]
	l, e := lexerFromFile(path)
	if e != nil {
		fmt.Println(e)
		return
	}
	p, e := Parser.ParsingStatement(&l, Token.END)
	if e != nil {
		fmt.Println(e)
		return
	}
	env := Evaluator.NewEnvironment()
	Evaluator.LoadBuiltInFunction(env)
	Evaluator.LoadBuiltInVariable(env)
	_, e = Evaluator.Eval(p.(*Parser.StatementNode), env)
	if e != nil {
		fmt.Println(e)
		return
	}
	return
}
