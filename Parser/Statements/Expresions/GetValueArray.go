package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
)

func parseGetValueArray(l *Lexer.Lexer) (IExpresion, error) {
	array := ExpresionGetValueArray{}
	array.Name = l.LookCurrent().Value
	l.IncrP()
	l.IncrP()
	values, e := ParseExpresionsGroup(l, nil, Token.CLOSE_SQUARE_BRACKET, Token.COMMA)
	if e != nil {
		return nil, e
	}
	array.IndexsValues = values
	l.IncrP()
	return array, nil

}
