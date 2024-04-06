package Statements

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Expresions"
	"errors"
	"slices"
)

func ParsingStatement(l *Lexer.Lexer, exitTokens ...Token.TokenType) (IStatement, error) {
	program := &StatementNode{}
	head := program
	for {
		switch l.LookCurrent().Type {
		case Token.LET:
			letS, e := ParseLetStatement(l)
			if e != nil {

				return nil, e
			}
			head.AddStatement(letS)
			head.AddNext(&StatementNode{})
			head = head.Next
		case Token.IF:
			letS, e := ParseIfStatement(l)
			if e != nil {
				return nil, e
			}
			head.AddStatement(letS)
			head.AddNext(&StatementNode{})
			head = head.Next
		case Token.WORD:
			letS, e := ParseAssignment(l)
			if e != nil {
				return nil, e
			}
			head.AddStatement(letS)
			head.AddNext(&StatementNode{})
			head = head.Next
		case Token.CALL_FUNC:
			letS, e := ParseCallFunc(l)
			if e != nil {
				return nil, e
			}
			head.AddStatement(letS)
			head.AddNext(&StatementNode{})
			head = head.Next
		default:
			if slices.Contains(exitTokens, l.LookCurrent().Type) {
				return program, nil
			}
			return nil, errors.New("ParsingStatement: unexpected statement token,got:" + l.LookCurrent().Value)
		}
	}
}
func ParseLetStatement(l *Lexer.Lexer) (IStatement, error) {
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
	l.IncrP()
	let.Expresion, e = Expresions.ParseExpresion(l, Token.DOT_COMMA)
	if e != nil {
		return nil, e
	}
	l.IncrP()
	return &let, nil
}

func ParseIfStatement(l *Lexer.Lexer) (IStatement, error) {
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
	ifExpr.Body, e = ParsingStatement(l, Token.CLOSE_GRAP_BRACKET)
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
	ifExpr.Else, e = ParsingStatement(l, Token.CLOSE_GRAP_BRACKET)
	if e != nil {
		return nil, e
	}
	l.IncrP()
	return ifExpr, nil
}
func ParseAssignment(l *Lexer.Lexer) (IStatement, error) {
	ass := AssignExpresionStatement{}
	ass.Identifier = l.LookCurrent().Value
	nextToken, e := l.NextToken()
	if e != nil {
		return nil, e
	}
	if nextToken.Type != Token.ASSIGN {
		return nil, errors.New("parseAssignment: expected '=' token")
	}
	l.IncrP()
	ass.Expresion, e = Expresions.ParseExpresion(l, Token.DOT_COMMA)

	if e != nil {
		return nil, e
	}
	l.IncrP()
	return ass, nil
}
func ParseCallFunc(l *Lexer.Lexer) (IStatement, error) {
	callFunc := CallFuncStatement{}
	exp, e := Expresions.ParseExpresion(l, Token.DOT_COMMA)
	if e != nil {
		return nil, e
	}
	callFunc.Expresion = exp
	l.IncrP()
	return callFunc, nil
}
