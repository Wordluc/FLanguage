package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"errors"
)

func parseDeclareArray(l *Lexer.Lexer, back IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {
	array := ExpresionDeclareArray{}
	switch back.(type) {
	case ExpresionLeaf:
		return parseGetValueArray(l)
	case ExpresionDeclareArray, ExpresionGetValueArray:
		return nil, errors.New("parseDeclareArray: unexpected array")
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
