package Statements

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"errors"
	"slices"
)

func ParsingStatement(l *Lexer.Lexer, exitTokens ...Token.TokenType) (IStatement, map[string]IStatement, error) {
	funcs := make(map[string]IStatement)
	program := &StatementNode{}
	head := program
	for {
		switch l.LookCurrent().Type {
		case Token.FUNC:
			f, e := ParsingFuncDeclaration(l)
			if e != nil {
				return nil, nil, e
			}
			funcs[f.(FuncDeclarationStatement).Identifier] = f
		case Token.LET:
			letS, e := parseLetStatement(l)
			if e != nil {

				return nil, nil, e
			}
			head.addStatement(letS)
			head.addNext(&StatementNode{})
			head = head.Next
		case Token.IF:
			letS, e := parseIfStatement(l)
			if e != nil {
				return nil, nil, e
			}
			head.addStatement(letS)
			head.addNext(&StatementNode{})
			head = head.Next
		case Token.WORD:
			letS, e := parseAssignment(l)
			if e != nil {
				return nil, nil, e
			}
			head.addStatement(letS)
			head.addNext(&StatementNode{})
			head = head.Next
		case Token.CALL_FUNC:
			letS, e := parseCallFunc(l)
			if e != nil {
				return nil, nil, e
			}
			head.addStatement(letS)
			head.addNext(&StatementNode{})
			head = head.Next //Inserire return statement
		default:
			if slices.Contains(exitTokens, l.LookCurrent().Type) {
				return program, funcs, nil
			}
			return nil, funcs, errors.New("ParsingStatement: unexpected statement token,got:" + l.LookCurrent().Value)
		}
	}
}
