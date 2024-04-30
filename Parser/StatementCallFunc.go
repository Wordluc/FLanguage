package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
)

func parseCallFunc(l *Lexer.Lexer) (IStatement, error) {
	callFunc := CallFuncStatement{}
	exp, e := ParseExpresion(l, Token.DOT_COMMA)
	if e != nil {
		return nil, e
	}
	callFunc.Expresion = exp
	l.IncrP()
	return callFunc, nil
}
