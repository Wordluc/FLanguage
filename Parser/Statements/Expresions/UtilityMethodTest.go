package Expresions

import (
	"FLanguage/Lexer/Token"
)

type CompareWith struct {
	TypeA         Token.TokenType
	ValueA        string
	Operator      Token.TokenType
	ValueOperator string
}
