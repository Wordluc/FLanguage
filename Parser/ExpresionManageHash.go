package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"errors"
)

func parseHash(l *Lexer.Lexer, back IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {
	if back == nil {
		return parseDeclareHash(l, back)
	}
	switch back.(type) {
	case ExpresionLeaf:
		return parseGetValueHash(l, back)
	case ExpresionDeclareHash:
		return parseGetValueHash(l, back)
	case ExpresionGetValueHash:
		return parseGetValueHash(l, back)
	case ExpresionCallFunc:
		return parseGetValueHash(l, back)
	}
	return nil, errors.New("parseHash: invalid expresion")
}

func parseDeclareHash(l *Lexer.Lexer, _ IExpresion) (IExpresion, error) {
	array := ExpresionDeclareHash{}
	l.IncrP()
	values, e := parseExpresionsGroupHash(l, nil, Token.CLOSE_GRAP_BRACKET, Token.COMMA)
	if e != nil {
		return nil, e
	}
	l.IncrP()
	array.Values = values
	return array, nil
}
func parseExpresionsGroupHash(l *Lexer.Lexer, _ IExpresion, exist Token.TokenType, delimiter Token.TokenType) (map[IExpresion]IExpresion, error) {
	var values map[IExpresion]IExpresion
	values = make(map[IExpresion]IExpresion)
	for {
		if l.LookCurrent().Type == exist {
			break
		}
		parmname, e := ParseExpresion(l, Token.DOUBLE_DOT)
		if e != nil {
			return nil, e
		}
		parmName, isLeaf := parmname.(ExpresionLeaf)
		if !isLeaf {
			return nil, errors.New("parseExpresionsGroupHash: invalid expresion as identifier")
		}
		l.IncrP()
		parmValue, e := ParseExpresion(l, delimiter, exist)
		if e != nil {
			return nil, e
		}
		values[parmName] = parmValue
		if l.LookCurrent().Type == exist {
			break
		}
		l.IncrP()
	}
	return values, nil
}
func parseGetValueHash(l *Lexer.Lexer, back IExpresion) (IExpresion, error) {
	hash := ExpresionGetValueHash{}
	hash.Value = back
	l.IncrP()
	values, e := ParseExpresion(l, Token.CLOSE_GRAP_BRACKET)
	if e != nil {
		return nil, e
	}
	hash.Index = values
	token, e := l.LookNext()
	if e == nil && token.Type == Token.OPEN_CIRCLE_BRACKET {

		return ParseCallFunc(l, hash)
	}
	l.IncrP()
	return hash, nil
}
