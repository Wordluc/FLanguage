package Statements

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements/Expresions"
	"errors"
)

func parseIfStatement(l *Lexer.Lexer) (IStatement, error) {
	ifExpr := IfStatement{}
	curToken, e := l.NextToken()
	if e != nil {
		return nil, e
	}
	if curToken.Type != Token.OPEN_CIRCLE_BRACKET {
		return nil, errors.New("parseIfStatement: expected '(' token")
	}
	l.IncrP()
	ifExpr.Expresion, e = Expresions.ParseExpresion(l,
		Token.CLOSE_CIRCLE_BRACKET)
	if e != nil {
		return nil, e
	}
	curToken = l.LookCurrent()
	curToken, e = l.NextToken()
	if e != nil {
		return nil, e
	}
	if curToken.Type != Token.OPEN_GRAP_BRACKET {
		return nil, errors.New("parseIfStatement: expected '{' token")
	}
	l.IncrP()
	program, e := ParsingStatement(l, Token.CLOSE_GRAP_BRACKET)
	ifExpr.Body = program
	if e != nil {
		return nil, e
	}
	curlToken, e := l.NextToken()

	if e != nil {
		return nil, e
	}
	if curlToken.Type != Token.ELSE {
		return ifExpr, nil
	}

	curToken, e = l.NextToken()
	if e != nil {
		return nil, e
	}
	if curToken.Type != Token.OPEN_GRAP_BRACKET {
		return nil, errors.New("parseIfStatement: expected '{' token")
	}
	l.IncrP()
	program, e = ParsingStatement(l, Token.CLOSE_GRAP_BRACKET)
	if e != nil {
		return nil, e
	}
	ifExpr.Else = program
	l.IncrP()
	return ifExpr, nil
}
