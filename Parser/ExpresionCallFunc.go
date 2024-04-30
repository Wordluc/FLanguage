package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
)

func ParseCallFunc(l *Lexer.Lexer, back IExpresion, _ ...Token.TokenType) (IExpresion, error) {
	callFunc := ExpresionCallFunc{Identifier: back}
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
