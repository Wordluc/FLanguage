package Statements

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements/Expresions"
	"errors"
)

func parseIfStatement(l *Lexer.Lexer) (IStatement, error) {

	ifExpr := &IfStatement{}

	curToken, e := l.NextToken()
	if e != nil {
		return nil, e
	}
	if curToken.Type != Token.OPEN_CIRCLE_BRACKET {
		return nil, errors.New("parseIfStatement: expected '(' token")
	}
	l.IncrP()
	ifExpr.FirstExpresion, e = Expresions.ParseExpresion(l,
		Token.EQUAL, Token.LESS_EQUAL, Token.GREATER_EQUAL, Token.NOT_EQUAL, Token.GREATER, Token.LESS) //== <= >= !=
	if e != nil {
		return nil, e
	}
	curToken = l.LookCurrent()
	ifExpr.ConditionType = curToken.Type
	ifExpr.ConditionValue = curToken.Value
	curToken, e = l.NextToken()
	ifExpr.LastExpresion, e = Expresions.ParseExpresion(l, Token.CLOSE_CIRCLE_BRACKET)
	curToken, e = l.NextToken()

	if e != nil {
		return nil, e
	}
	if curToken.Type != Token.OPEN_GRAP_BRACKET {
		return nil, errors.New("parseIfStatement: expected '{' token")
	}
	l.IncrP()
	ifExpr.Body, _, e = ParsingStatement(l, Token.CLOSE_GRAP_BRACKET)
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
	ifExpr.Else, _, e = ParsingStatement(l, Token.CLOSE_GRAP_BRACKET)
	if e != nil {
		return nil, e
	}
	l.IncrP()
	return ifExpr, nil
}
