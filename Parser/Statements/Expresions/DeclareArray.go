package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"errors"
)

func parseDeclareArray(l *Lexer.Lexer, _ IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {
	array := ExpresionDeclareArray{}
	back, e := l.LookBack()
	if e == nil {
		if back.Type != Token.ASSIGN && back.Type != Token.COMMA && back.Type != Token.OPEN_SQUARE_BRACKET {
			return nil, errors.New("parseDeclareArray: impossible create array")
		}
	}
	l.IncrP()
	values, e := ParseExpresionsGroup(l, nil, Token.CLOSE_SQUARE_BRACKET, Token.COMMA)
	if e != nil {
		return nil, e
	}
	l.IncrP()
	array.Values = values
	return array, nil

}
