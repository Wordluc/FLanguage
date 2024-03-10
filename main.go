package main

import (
	Lexer "FLanguage/Lexer"
	"fmt"
)

func lexerFromFile() {
	file, _ := Lexer.OpenFile("prova.txt")
	lexer, _ := Lexer.New(file)
	for {
		v, e := lexer.NextToken(0)
		if e != nil {
			break
		}
		fmt.Println(v)
	}
}
func main() {
	Lexer.ReplLexer()
}
