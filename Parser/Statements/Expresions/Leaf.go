package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
)

func parseLeaf(l *Lexer.Lexer, _ IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {
	nextT, e := l.LookNext()
	if e != nil {
		return nil, e
	}
	if l.LookCurrent().Type == Token.WORD {

		if nextT.Type == Token.OPEN_CIRCLE_BRACKET {
			return parseCallFunc(l, nil)
		}
		if nextT.Type == Token.OPEN_SQUARE_BRACKET {
			return parseGetValueArray(l)
		}
	}
	curToken := l.LookCurrent()
	leaf := ExpresionLeaf{}
	l.IncrP()
	return leaf.New(curToken), nil
}
