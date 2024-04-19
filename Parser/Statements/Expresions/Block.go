package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
)

func parseExpresionBlock(l *Lexer.Lexer, _ IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {
	l.IncrP()
	block, e := ParseExpresion(l, Token.CLOSE_CIRCLE_BRACKET)
	if e != nil {
		return nil, e
	}
	l.IncrP()
	return block, nil
}
