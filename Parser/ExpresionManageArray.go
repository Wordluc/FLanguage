package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"errors"
)

func parseArray(l *Lexer.Lexer, back IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {
	if back == nil {
		return parseDeclareArray(l, back)
	}
	switch back.(type) {
	case ExpresionLeaf:
		return parseGetValueArray(l, back)
	case ExpresionDeclareArray:
		return parseGetValueArray(l, back)
	case ExpresionGetValueArray:
		return parseGetValueArray(l, back)
	case ExpresionCallFunc:
		return parseGetValueArray(l, back)
	}
	return nil, errors.New("parseArray: invalid expresion")
}

func parseDeclareArray(l *Lexer.Lexer, _ IExpresion) (IExpresion, error) {
	array := ExpresionDeclareArray{}
	l.IncrP()
	values, e := ParseExpresionsGroup(l, nil, Token.CLOSE_SQUARE_BRACKET, Token.COMMA)
	if e != nil {
		return nil, e
	}
	l.IncrP()
	array.Values = values
	return array, nil
}
func parseGetValueArray(l *Lexer.Lexer, back IExpresion) (IExpresion, error) {
	array := ExpresionGetValueArray{}
	array.Value = back
	l.IncrP()
	values, e := ParseExpresionsGroup(l, nil, Token.CLOSE_SQUARE_BRACKET, Token.COMMA)
	if e != nil {
		return nil, e
	}
	array.IndexsValues = values
	token, e := l.LookNext()
	if e == nil && token.Type == Token.OPEN_CIRCLE_BRACKET {
		return ParseCallFunc(l, array)
	}
	l.IncrP()
	return array, nil

}
