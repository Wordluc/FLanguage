package Statements

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Expresions"
	"errors"
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
		case Token.END:
			return program, nil
		default:
			return nil, errors.New("ParsingStatement: unexpected statement token")
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
	if curToken.Type != Token.EQUAL {
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
