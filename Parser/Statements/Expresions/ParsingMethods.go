package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"errors"
	"slices"
)

type fParse func(l *Lexer.Lexer, expresion IExpresion, exitTokens ...Token.TokenType) (IExpresion, error)

func And(e error, s string) error {
	v := e.Error()
	return errors.New(v + " " + s)
}

func IsAValidBrach(token Token.Token) bool {
	return token.Type == Token.WORD || token.Type == Token.OPEN_CIRCLE_BRACKET || token.Type == Token.NUMBER || token.Type == Token.STRING
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
	case Token.WORD, Token.NUMBER, Token.STRING, Token.BOOLEAN:
		return parseLeaf, nil
	case Token.OPEN_CIRCLE_BRACKET:
		return parseExpresionBlock, nil
	case Token.OPEN_SQUARE_BRACKET:
		return parseDeclareArray, nil
	}
	return nil, errors.New("GetParse: Operator:" + string(than) + "not implemented")
}

func ParseExpresion(l *Lexer.Lexer, exitTokens ...Token.TokenType) (IExpresion, error) {
	var root IExpresion
	if exitTokens == nil {
		return nil, errors.New("ParseExpresion: no exitTokens defined")
	}
	for {
		lookCurrVar := l.LookCurrent()
		if slices.Contains(exitTokens, lookCurrVar.Type) {
			break
		}
		fVar, e := GetParse(lookCurrVar.Type)
		if e != nil {
			return nil, e
		}
		root, e = fVar(l, root, exitTokens...)
		if e != nil {
			return nil, e
		}
		lookCurrVar = l.LookCurrent()
		if slices.Contains(exitTokens, lookCurrVar.Type) {
			break
		}

	}
	return root, nil
}
