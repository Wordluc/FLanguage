package Expresions

import "FLanguage/Lexer/Token"

type Expresion struct {
	Type          Token.TokenType
	Value         string
	NextExpresion *Expresion
}
