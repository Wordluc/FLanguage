package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
)

func parseDeclareArray(l *Lexer.Lexer, _ IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {
	array := &ExpresionDeclareArray{}
	l.IncrP()
	values, e := parseExpresionsGroup(l, nil, Token.CLOSE_SQUARE_BRACKET, Token.COMMA)
	if e != nil {
		return nil, e
	}
	l.IncrP()
	array.Values = values
	return array, nil

}
