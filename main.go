package main

import (
	"FLanguage/Evaluator"
	Lexer "FLanguage/Lexer"
	"fmt"
)

func lexerFromFile() {
	file, _ := Lexer.GetByteFromFile("prova.txt")
	lexer, _ := Lexer.New(file)
	for {
		v, e := lexer.NextToken()
		if e != nil {
			break
		}
		fmt.Println(v)
	}
}
func main() {
	a := Evaluator.New()
	fmt.Println(a)
}
