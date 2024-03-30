package Statements

import (
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Expresions"
)

type IStatement interface {
	ToString() string
	TokenType() Token.TokenType
	ToExpression() Expresions.IExpresion
}
