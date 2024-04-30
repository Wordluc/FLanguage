package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
)

func ParseExpresionsGroup(l *Lexer.Lexer, _ IExpresion, exist Token.TokenType, delimiter Token.TokenType) ([]IExpresion, error) {
	var values []IExpresion
	for {

		if exist == l.LookCurrent().Type {
			break
		}

		parm, e := ParseExpresion(l, delimiter, exist)
		if e != nil {
			return nil, e
		}
		if parm == nil {
			break
		}
		values = append(values, parm)
		if exist == l.LookCurrent().Type {
			break
		}
		l.IncrP()
	}
	return values, nil
}
