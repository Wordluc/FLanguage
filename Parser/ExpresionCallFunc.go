package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
)

func expresionCallFunc(l *Lexer.Lexer, back IExpresion, _ ...Token.TokenType) (IExpresion, error) {
	callFunc := ExpresionCallFunc{Called: back}
	l.IncrP()
	parms, e := ParseExpresionsGroup(l, nil, Token.CLOSE_CIRCLE_BRACKET, Token.COMMA)
	callFunc.Params = parms
	if e != nil {
		return nil, e
	}
	l.IncrP()
	return callFunc, nil
}
