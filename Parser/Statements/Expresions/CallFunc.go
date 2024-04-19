package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
)

func parseCallFunc(l *Lexer.Lexer, _ IExpresion, _ ...Token.TokenType) (IExpresion, error) {

	callFunc := ExpresionCallFunc{NameFunc: l.LookCurrent().Value}
	l.IncrP()
	l.IncrP()
	parms, e := ParseExpresionsGroup(l, nil, Token.CLOSE_CIRCLE_BRACKET, Token.COMMA)
	callFunc.Values = parms
	if e != nil {
		return nil, e
	}
	l.IncrP()
	return callFunc, nil
}
