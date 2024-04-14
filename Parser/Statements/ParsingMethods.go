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
		case Token.FUNC:
			f, e := ParsingFuncDeclaration(l)
			if e != nil {
				return nil, e
			}
			head.addStatement(f)
			head.addNext(&StatementNode{})
			head = head.Next
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
			nextT, _ := l.LookNext()
			var e error
			var res IStatement
			if nextT.Type != Token.OPEN_CIRCLE_BRACKET {
				res, e = parseAssignment(l)
			} else {
				res, e = parseCallFunc(l)
			}

			if e != nil {
				return nil, e
			}
			head.addStatement(res)
			head.addNext(&StatementNode{})
			head = head.Next
			//		case Token.CALL_FUNC:
			//			letS, e := parseCallFunc(l)
			//			if e != nil {
			//				return nil, e
			//			}
			//			head.addStatement(letS)
			//			head.addNext(&StatementNode{})
			//			head = head.Next
		case Token.RETURN:
			runS, e := parseReturnStatement(l)
			if e != nil {
				return nil, e
			}
			head.addStatement(runS)
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
