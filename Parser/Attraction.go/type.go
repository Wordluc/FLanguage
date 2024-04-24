package Attraction

import (
	"FLanguage/Lexer/Token"
	"errors"
)

type Force uint8

const (
	F0 Force = iota
	F1
	F2
	F3
	F4
	F5
)

func GetForce(than Token.TokenType) (Force, error) {
	switch than {
	case Token.OPEN_CIRCLE_BRACKET:
		return F5, nil
	case Token.CLOSE_CIRCLE_BRACKET:
		return F5, nil
	case Token.OPEN_SQUARE_BRACKET:
		return F5, nil
	case Token.NOT_EQUAL, Token.EQUAL, Token.LESS, Token.GREATER, Token.LESS_EQUAL, Token.GREATER_EQUAL:
		return F1, nil
	case Token.PLUS:
		return F2, nil
	case Token.MINUS:
		return F2, nil
	case Token.DIV:
		return F3, nil
	case Token.MULT:
		return F3, nil
	case Token.FUNC:
		return F4, nil
	}
	return F5, errors.New("GetForce: " + string(than) + "not implemented")
}
