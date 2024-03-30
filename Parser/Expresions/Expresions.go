package Expresions

import "FLanguage/Lexer/Token"

type BinaryExpresion struct {
	ValueA        string
	TypeA         Token.TokenType
	Operator      Token.Token
	ValueOperator string
	NextExpresion IExpresion
}
type EmptyExpresion struct {
}
type IExpresion interface {
}
