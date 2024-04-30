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
	curToken := l.LookCurrent()
	leaf := ExpresionLeaf{}.New(curToken)
	if l.LookCurrent().Type == Token.WORD {
		if nextT.Type == Token.OPEN_CIRCLE_BRACKET {
			return parseCallFunc(l, leaf)
		}

	}
	l.IncrP()
	return leaf, nil
}
