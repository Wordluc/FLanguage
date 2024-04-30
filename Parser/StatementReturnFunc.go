package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
)

func parseReturnStatement(l *Lexer.Lexer) (IStatement, error) {
	r := ReturnStatement{}

	l.IncrP()
	exp, e := ParseExpresion(l, Token.DOT_COMMA)
	if e != nil {
		return nil, e
	}
	l.IncrP()
	r.Expresion = exp
	return r, nil
}
