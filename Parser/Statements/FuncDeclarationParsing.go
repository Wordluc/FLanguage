package Statements

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"errors"
)

func ParsingFuncDeclaration(lexer *Lexer.Lexer, exitTokens ...Token.TokenType) (IStatement, error) {
	funcDeclaration := FuncDeclarationStatement{}
	curToken, e := lexer.NextToken()
	if e != nil {
		return nil, e
	}
	if curToken.Type != Token.WORD {
		return nil, errors.New("ParsingFuncDeclaration: expected 'WORD' token,error Identifier")
	}
	funcDeclaration.Identifier = curToken.Value
	curToken, e = lexer.NextToken()
	if e != nil {
		return nil, e
	}
	if curToken.Type != Token.OPEN_CIRCLE_BRACKET {
		return nil, errors.New("ParsingFuncDeclaration: expected '(' token,error Parameters")
	}
	funcDeclaration.Params, e = ParseParms(lexer)
	if e != nil {
		return nil, e
	}
	curToken, e = lexer.NextToken()
	if curToken.Type != Token.OPEN_GRAP_BRACKET {
		return nil, errors.New("ParsingFuncDeclaration: expected '{' token")
	}

	lexer.IncrP()
	program, e := ParsingStatement(lexer, Token.CLOSE_GRAP_BRACKET)
	funcDeclaration.Body = program["root"]
	if e != nil {
		return nil, e
	}
	curToken = lexer.LookCurrent()
	if curToken.Type != Token.CLOSE_GRAP_BRACKET {
		return nil, errors.New("ParsingFuncDeclaration: expected '}' token")
	}

	lexer.IncrP()
	return funcDeclaration, nil
}

func ParseParms(lexer *Lexer.Lexer) ([]string, error) {
	parms := []string{}
	for {
		lookNextToken, e := lexer.LookNext()
		if e != nil {
			return nil, e
		}
		switch lookNextToken.Type {
		case Token.COMMA:
			lexer.IncrP()
		case Token.CLOSE_CIRCLE_BRACKET:
			lexer.IncrP()
			return parms, nil
		case Token.WORD:
			parms = append(parms, lookNextToken.Value)
			lexer.IncrP()
		default:
			return nil, errors.New("ParseParms: expected ',' or ')' token")
		}
	}
}
