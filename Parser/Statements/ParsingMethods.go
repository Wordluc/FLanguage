package Statements

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"errors"
	"slices"
)

func ParsingStatement(l *Lexer.Lexer, exitTokens ...Token.TokenType) (map[string]IStatement, error) {
	program := make(map[string]IStatement)
	main := &StatementNode{}
	head := main
	for {
		switch l.LookCurrent().Type {
		case Token.FUNC:
			f, e := ParsingFuncDeclaration(l)
			if e != nil {
				return nil, e
			}
			program[f.(FuncDeclarationStatement).Identifier] = f
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
			head = head.Next //Inserire return statement
		case Token.RETURN:
			runS, e := parseReturnStatement(l)
			if e != nil {
				return nil, e
			}
			head.addStatement(runS)
		default:
			if slices.Contains(exitTokens, l.LookCurrent().Type) {
				program["root"] = main
				return program, nil
			}
			return nil, errors.New("ParsingStatement: unexpected statement token,got:" + l.LookCurrent().Value)
		}
	}
}
