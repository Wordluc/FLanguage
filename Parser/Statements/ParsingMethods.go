package Statements

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"errors"
	"slices"
)

func ParsingStatement(l *Lexer.Lexer, exitTokens ...Token.TokenType) (IStatement, error) {
	program := &StatementNode{}
	head := program
	for {
		switch l.LookCurrent().Type {
		case Token.LET:
			letS, e := parseLetStatement(l)
			if e != nil {

				return nil, e
			}
			head.addStatement(letS)
			head.addNext(&StatementNode{})
			head = head.Next
		case Token.IF:
			letS, e := parseIfStatement(l)
			if e != nil {
				return nil, e
			}
			head.addStatement(letS)
			head.addNext(&StatementNode{})
			head = head.Next
		case Token.WORD:
			letS, e := parseAssignment(l)
			if e != nil {
				return nil, e
			}
			head.addStatement(letS)
			head.addNext(&StatementNode{})
			head = head.Next
		case Token.CALL_FUNC:
			letS, e := parseCallFunc(l)
			if e != nil {
				return nil, e
			}
			head.addStatement(letS)
			head.addNext(&StatementNode{})
			head = head.Next
		default:
			if slices.Contains(exitTokens, l.LookCurrent().Type) {
				return program, nil
			}
			return nil, errors.New("ParsingStatement: unexpected statement token,got:" + l.LookCurrent().Value)
		}
	}
}
