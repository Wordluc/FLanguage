package Statements

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements/Expresions"
	"errors"
)

func parseWhileStatement(lexer *Lexer.Lexer) (IStatement, error) {
	next, e := lexer.NextToken()
	var cond Expresions.IExpresion
	if next.Type != Token.OPEN_CIRCLE_BRACKET {
		return nil, errors.New("ParseWhileStatement: expected '(' token,got:" + next.Value)
	}

	lexer.IncrP()
	cond, e = Expresions.ParseExpresion(lexer, Token.CLOSE_CIRCLE_BRACKET)
	if e != nil {
		return nil, e
	}
	next, e = lexer.NextToken()
	if next.Type != Token.OPEN_GRAP_BRACKET {
		return nil, errors.New("ParseWhileStatement: expected '{' token,got:" + next.Value)
	}

	lexer.IncrP()
	exp, e := ParsingStatement(lexer, Token.CLOSE_GRAP_BRACKET)
	lexer.IncrP()

	return WhileStatement{Cond: cond, Body: exp}, nil
}
