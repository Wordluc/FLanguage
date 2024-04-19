package main

import (
	"FLanguage/Evaluator"
	Lexer "FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements"
	"bufio"
	"fmt"
	"os"
)

func lexerFromFile(path string) (Lexer.Lexer, error) {
	file, _ := Lexer.GetByteFromFile(path)
	lexer, e := Lexer.New(file)
	return lexer, e
}
func main() {
	reder := bufio.NewReader(os.Stdin)
	fmt.Print("-")
	v, _ := reder.ReadBytes('\n')
	v = v[:len(v)-2]
	l, e := lexerFromFile(string(v))
	if e != nil {
		fmt.Println(e)
		return
	}

	p, e := Statements.ParsingStatement(&l, Token.END)
	if e != nil {
		fmt.Println(e)
		return
	}
	println("-----INIZIO-----")
	env := Evaluator.NewEnvironment()
	Evaluator.LoadBuiltInFunction(env)
	Evaluator.LoadBuiltInVariable(env)
	_, e = Evaluator.Eval(p.(*Statements.StatementNode), env)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println("-----FINE-----")
	return
}
