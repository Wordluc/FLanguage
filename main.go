package main

import (
	Lexer "FLanguage/Lexer"
	"fmt"
)

func main() {
	p, _ := Lexer.Parse("prova.txt")
	fmt.Print(p)
}
