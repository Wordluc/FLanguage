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
	value, e := ParseExpresion(l, Token.CLOSE_SQUARE_BRACKET)
	if e != nil {
		return nil, e
	}
	array.ValueId = value
	l.IncrP()
	return array, nil

}
