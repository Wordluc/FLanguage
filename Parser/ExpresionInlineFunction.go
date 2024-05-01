package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"errors"
)

func parseInlineFunction(lexer *Lexer.Lexer, _ IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {
	funcDeclaration := FuncDeclarationStatement{}
	curToken, e := lexer.NextToken()
	if e != nil {
		return nil, e
	}
	if curToken.Type != Token.OPEN_CIRCLE_BRACKET {
		return nil, errors.New("ParsingFuncDeclaration: expected '(' token,error Parameters")
	}
	parms, e := parseParms(lexer)
	if e != nil {
		return nil, e
	}
	funcDeclaration.Params = parms
	curToken, e = lexer.NextToken()

	if e != nil {
		return nil, e
	}
	if curToken.Type != Token.OPEN_GRAP_BRACKET {
		return nil, errors.New("ParsingFuncDeclaration: expected '{' token")
	}
	lexer.IncrP()
	program, e := ParsingStatement(lexer, Token.CLOSE_GRAP_BRACKET)
	if e != nil {
		return nil, e
	}
	funcDeclaration.Body = program
	curToken, e = lexer.LookNext()
	if e == nil && curToken.Type == Token.OPEN_CIRCLE_BRACKET {
		return expresionCallFunc(lexer, funcDeclaration)
	}
	lexer.IncrP()
	return funcDeclaration, nil
}

func parseParms(lexer *Lexer.Lexer) ([]string, error) {
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
