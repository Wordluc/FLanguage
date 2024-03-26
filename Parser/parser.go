package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Attraction.go"
	"FLanguage/Parser/ParsingExpression/ParsingStatements"
	"errors"
)

type Nodo struct {
	Statement ParsingStatements.IStatement
	Next      *Nodo
}

func ParseProgram(lexer *Lexer.Lexer) (*Nodo, error) {
	start := &Nodo{}
	head := start
	for {
		statement, e := parseStatement(lexer)
		if e != nil {
			return head, e
		}
		head.Statement = statement
		if (statement).GetTokenType() == Token.END {
			return start, nil
		}
		head.Next = &Nodo{}
	}
}
func parseStatement(lexer *Lexer.Lexer) (ParsingStatements.IStatement, error) {
	token, e := lexer.NextToken()
	if e != nil {
		return nil, e
	}
	switch token.Type {
	case Token.LET:
		return ParsingStatements.ParseLetStatement(lexer)
	case Token.FUNC:
	case Token.IF:
	}
	return nil, errors.New("ParseStatement: " + string(token.Type) + "not implemented")

}

func GetForce(than Token.TokenType) (Attraction.Force, error) {
	switch than {
	case Token.OPEN_CIRCLE_BRACKET:
		return Attraction.F5, nil
	case Token.CLOSE_CIRCLE_BRACKET:
		return Attraction.F5, nil
	case Token.WORD:
		return Attraction.F0, nil
	case Token.PLUS:
		return Attraction.F2, nil
	case Token.MINUS:
		return Attraction.F2, nil
	case Token.DIV:
		return Attraction.F3, nil
	case Token.MULT:
		return Attraction.F3, nil
	case Token.FUNC:
		return Attraction.F4, nil
	}
	return Attraction.F5, errors.New("GetForce: " + string(than) + "not implemented")
}
