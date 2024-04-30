package Statements

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements/Expresions"
	"errors"
)

func parseSetHashValue(l *Lexer.Lexer) (IStatement, error) {
	array := SetHashValueStatement{
		Identifier: l.LookCurrent().Value,
	}
	l.IncrP()
	l.IncrP()
	index, e := Expresions.ParseExpresion(l, Token.CLOSE_GRAP_BRACKET)
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
	value, e := Expresions.ParseExpresion(l, Token.DOT_COMMA)
	if e != nil {
		return nil, e
	}
	l.IncrP()
	array.Index = index
	array.Value = value
	return array, nil
}
