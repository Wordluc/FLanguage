package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"errors"
)

func parseAssignment(l *Lexer.Lexer) (IStatement, error) {
	ass := AssignExpresionStatement{}
	ass.Identifier = l.LookCurrent().Value
	nextToken, e := l.NextToken()
	if e != nil {
		return nil, e
	}
	if nextToken.Type != Token.ASSIGN {
		return nil, errors.New("parseAssignment: expected '=' token")
	}
	l.IncrP()
	ass.Expresion, e = ParseExpresion(l, Token.DOT_COMMA)

	if e != nil {
		return nil, e
	}
	l.IncrP()
	return ass, nil
}
