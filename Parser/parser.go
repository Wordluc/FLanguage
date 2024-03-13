package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"errors"
)

type INodo interface {
	getTokenType()
	setTokenType(Token.TokenType)
	setNextNode(INodo)
}
type Nodo struct {
}

//func (n *Nodo) toString() string {
//	return string(n.typeToken) + n.toString()
//
//}

type letNode struct {
	INodo
	name      string
	expresion *Nodo
	next      *Nodo
}

func ParseProgram(lexer *Lexer.Lexer) {
	root := Nodo{}
}
func parseStatement(emptyNode *Nodo, lexer *Lexer.Lexer) (Nodo, error) {
	token, e := lexer.NextToken()
	if e != nil {
		return Nodo{}, e
	}
	switch token.Type {
	case Token.LET:
	case Token.FUNC:
	case Token.IF:
	default:
	}

}
func parseExpresion(preNode *INodo, lexer *Lexer.Lexer) {
	node := lexer.NextToken()
}
func parseLet(token Token.Token, lexer *Lexer.Lexer) (INodo, error) {
	nodo := letNode{}
	nextToken, e := lexer.NextToken()
	if e != nil {
		return nodo, e
	}
	if nextToken.Type != Token.WORD {
		return nodo, errors.New("expected word identifier for let")
	}
	nodo.name = token.Value
	if nextToken.Type != Token.EQUAL {
		return nodo, errors.New("expected equal sign for let")
	}
	parseExpresion(nodo, lexer)
}
