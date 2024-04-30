package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"errors"
)

func parseSetArrayValue(l *Lexer.Lexer) (IStatement, error) {
	array := SetArrayValueStatement{
		Identifier: l.LookCurrent().Value,
	}
	l.IncrP()
	l.IncrP()
	index, e := ParseExpresionsGroup(l, nil, Token.CLOSE_SQUARE_BRACKET, Token.COMMA)
	if e != nil {
		return nil, e
	}
	nextToken, e := l.NextToken()
	if e != nil {
		return nil, e
	}
	if nextToken.Type != Token.ASSIGN {
		return nil, errors.New("parseSetArrayValue: expected '=' token")
	}
	l.IncrP()
	value, e := ParseExpresion(l, Token.DOT_COMMA)
	if e != nil {
		return nil, e
	}
	l.IncrP()
	array.Indexs = index
	array.Value = value
	return array, nil
}
