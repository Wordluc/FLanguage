package Parser

import (
	"FLanguage/Lexer/Token"
)

type Nodo struct {
	Statement IStatement
	Next      *Nodo
}
type IStatement interface {
	getStatement() string
	getTokenType() Token.TokenType
}
type LetStatement struct {
	Type      Token.TokenType
	Variable  string
	Expresion Expresion
}

func (I *LetStatement) getStatement() string {
	return "LET " + I.Variable + "=" + I.Expresion.getExpresion()
}
func (I *LetStatement) getTokenType() Token.TokenType {
	return I.Type
}

type Expresion struct {
	TypeToken      Token.TokenType
	Value          string
	InnerExpresion *Expresion
	NextExpresion  *Expresion
}

func (E *Expresion) getExpresion() string {
	r := E.Value
	if E.InnerExpresion != nil {
		r += "(" + E.InnerExpresion.getExpresion() + ")"
	}
	if E.NextExpresion != nil {
		r += E.NextExpresion.getExpresion()
	}
	return r
}
