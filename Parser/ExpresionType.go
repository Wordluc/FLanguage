package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"errors"
)

type fParse func(l *Lexer.Lexer, expresion IExpresion, exitTokens ...Token.TokenType) (IExpresion, error)

func And(e error, s string) error {
	v := e.Error()
	return errors.New(v + " " + s)
}

func IsAValidBrach(token Token.Token) bool {
	return token.Type == Token.WORD || token.Type == Token.OPEN_CIRCLE_BRACKET || token.Type == Token.NUMBER || token.Type == Token.NUMBER_WITH_DOT || token.Type == Token.STRING
}

func IsAValidOperator(token Token.Token) bool {
	return !IsAValidBrach(token)
}

func GetParse(than Token.TokenType) (fParse, error) {
	switch than {
	case Token.DIV:
		return parseTree, nil
	case Token.MULT, Token.MINUS:
		return parseTree, nil
	case Token.EQUAL, Token.NOT_EQUAL, Token.LESS, Token.GREATER, Token.LESS_EQUAL, Token.GREATER_EQUAL:
		return parseTree, nil
	case Token.PLUS:
		return parseTree, nil
	case Token.WORD, Token.NUMBER, Token.STRING, Token.BOOLEAN, Token.NUMBER_WITH_DOT:
		return parseLeaf, nil
	case Token.OPEN_CIRCLE_BRACKET:
		return parseCircleBracket, nil
	case Token.OPEN_SQUARE_BRACKET:
		return parseArray, nil
	case Token.OPEN_GRAP_BRACKET:
		return parseHash, nil
	case Token.AT:
		return parseInlineFunction, nil
	}
	return nil, errors.New("GetParse: Operator:" + string(than) + "not implemented")
}
