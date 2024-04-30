package Statements

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements/Expresions"
	"errors"
)

func parseLetStatement(l *Lexer.Lexer) (IStatement, error) {
	let := LetStatement{}
	curToken, e := l.NextToken()
	if e != nil {
		return nil, e
	}
	if curToken.Type != Token.WORD {
		return nil, errors.New("parseLetStatement: expected 'WORD' token")
	}
	let.Identifier = curToken.Value
	curToken, e = l.NextToken()
	if e != nil {
		return nil, e
	}
	if curToken.Type != Token.ASSIGN {
		return nil, errors.New("parseLetStatement: expected '=' token")
	}
	curToken, _ = l.NextToken()
	if curToken.Type == Token.AT {
		l.IncrP()
		let.Expresion, e = ParsingInlineFunction(l)
		if e != nil {
			return nil, e
		}
		l.IncrP()
		return let, nil
	}
	let.Expresion, e = Expresions.ParseExpresion(l, Token.DOT_COMMA)
	if e != nil {
		return nil, e
	}
	l.IncrP()
	return let, nil
}
