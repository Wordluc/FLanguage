package Parser

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

			nextT, e := l.LookNext()
			tl := *l
			var res IStatement
			switch nextT.Type {
			case Token.OPEN_CIRCLE_BRACKET:
				res, e = parseCallFunc(&tl)
			case Token.OPEN_SQUARE_BRACKET:
				res, e = parseSetArrayValue(&tl)
			case Token.OPEN_GRAP_BRACKET:
				res, e = parseSetHashValue(&tl)
			case Token.ASSIGN:
				res, e = parseAssignment(&tl)
			default:
				return nil, errors.New("ParsingStatement: unexpected token,got:" + nextT.Value)
			}
			if e != nil {
				res, e = ParseExpresion(l, Token.DOT_COMMA)
				if e != nil {
					return nil, e
				}
				l.IncrP()
			} else {
				pc, _ := tl.GetP()
				l.SetpCurrent(pc)
			}
			head.addStatement(res)
			head.addNext(&StatementNode{})
			head = head.Next
		case Token.RETURN:
			runS, e := parseReturnStatement(l)
			if e != nil {
				return nil, e
			}
			head.addStatement(runS)
			head.addNext(&StatementNode{})
			head = head.Next
		case Token.OPEN_COMM:
			for {
				l.IncrP()
				if l.LookCurrent().Type == Token.CLOSE_COMM {
					l.IncrP()
					break
				}
			}
		case Token.LINE_COMM:
			l.IncrP()
		case Token.WHILE:
			runS, e := parseWhileStatement(l)
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
