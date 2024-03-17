package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Attraction.go"
	"errors"
	"fmt"
)

func ParseProgram(lexer *Lexer.Lexer) (*Nodo, error) {
	start := &Nodo{}
	head := start
	for {
		statement, e := parseStatement(lexer)
		if e != nil {
			return head, e
		}
		head.Statement = statement
		if (statement).getTokenType() == Token.END {
			return start, nil
		}
		head.Next = &Nodo{}
	}
}
func parseStatement(lexer *Lexer.Lexer) (IStatement, error) {
	token, e := lexer.NextToken()
	if e != nil {
		return nil, e
	}
	switch token.Type {
	case Token.LET:
		return parseLetStatement(lexer)
	case Token.FUNC:
	case Token.IF:
	}
	return nil, errors.New("ParseStatement: " + string(token.Type) + "not implemented")

}
func parseExpresion(f Attraction.Force, lexer *Lexer.Lexer, exp *Expresion) (*Expresion, error) {
	token, e := lexer.NextToken()
	if e != nil {
		return nil, e
	}
	fmt.Println("selected token", token.Value)
	lookedToken, e := lexer.LookNext()

	if e != nil {
		return nil, e
	}
	fmt.Println("lookked token", lookedToken.Value)
	if lookedToken.Type == Token.END || lookedToken.Type == Token.DOT_COMMA {
		fmt.Println("prossimo token no", token.Type, token.Value)
		exp.NextExpresion = &Expresion{TypeToken: token.Type, Value: token.Value}
		lexer.IncrP()
		return exp, nil

	}
	forceNext, e := GetForce(lookedToken.Type)
	if e != nil {
		return nil, e
	}
	if forceNext > f {
		fmt.Println("crea inner", token.Type, token.Value)
		inerExp := &Expresion{TypeToken: token.Type, Value: token.Value}
		inerExp.NextExpresion = &Expresion{TypeToken: lookedToken.Type, Value: lookedToken.Value}
		lexer.IncrP()
		_, e = parseExpresion(forceNext, lexer, inerExp.NextExpresion)
		exp.InnerExpresion = inerExp
		return exp, e
	} else if forceNext == f {
		exp.NextExpresion = &Expresion{TypeToken: token.Type, Value: token.Value}
		exp.NextExpresion.NextExpresion = &Expresion{TypeToken: lookedToken.Type, Value: lookedToken.Value}
		lexer.IncrP()
		_, e = parseExpresion(f, lexer, exp.NextExpresion.NextExpresion)
		return exp, e
	}
	fmt.Println("crea next")
	exp.NextExpresion = &Expresion{TypeToken: token.Type, Value: token.Value}
	//_, e = parseExpresion(f, lexer, exp.NextExpresion)
	return exp, nil
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
func parseLetStatement(lexer *Lexer.Lexer) (*LetStatement, error) {
	statement := LetStatement{Type: Token.LET}
	nextToken, e := lexer.NextToken()
	if e != nil || nextToken.Type != Token.WORD {
		return &statement, errors.New("expected word identifier for let")
	}
	statement.Variable = nextToken.Value
	nextToken, e = lexer.NextToken()
	if e != nil || nextToken.Type != Token.EQUAL {
		return &statement, errors.New("expected equal sign for let")
	}
	head := &Expresion{}
	expresion := head
	for {
		head.NextExpresion, e = parseExpresion(Attraction.F0, lexer, &Expresion{})
		if e != nil {
			break
		}
		nextToken, e = lexer.LookNext()
		if e != nil || nextToken.Type == Token.DOT_COMMA {
			break
		}
		head = head.NextExpresion
		head.NextExpresion = &Expresion{TypeToken: nextToken.Type, Value: nextToken.Value}
		lexer.IncrP()
		head = head.NextExpresion
	}
	statement.Expresion = *expresion
	return &statement, nil
}
