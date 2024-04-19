package Statements

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements/Expresions"
)

func parseCallFunc(l *Lexer.Lexer) (IStatement, error) {
	callFunc := CallFuncStatement{}
	exp, e := Expresions.ParseExpresion(l, Token.DOT_COMMA)
	if e != nil {
		return nil, e
	}
	callFunc.Expresion = exp
	l.IncrP()
	return callFunc, nil
}
