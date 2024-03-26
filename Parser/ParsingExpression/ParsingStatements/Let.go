package ParsingStatements

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Attraction.go"
	"FLanguage/Parser/ParsingExpression"
	"errors"
)

type LetStatement struct {
	Type      Token.TokenType
	Variable  string
	Expresion ParsingExpression.Expresion
}

func (I *LetStatement) GetStatement() string {
	return "LET " + I.Variable + "=" + I.Expresion.GetExpresion()
}
func (I *LetStatement) GetTokenType() Token.TokenType {
	return I.Type
}
func ParseLetStatement(lexer *Lexer.Lexer) (*LetStatement, error) {
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
	head := &ParsingExpression.Expresion{}
	expresion := head
	for {
		head.NextExpresion, e = ParsingExpression.ParseExpresion(Attraction.F0, lexer, &ParsingExpression.Expresion{})
		if e != nil {
			break
		}
		nextToken, e = lexer.LookNext()
		if e != nil || nextToken.Type == Token.DOT_COMMA {
			break
		}
		head = head.NextExpresion
		head.NextExpresion = &ParsingExpression.Expresion{TypeToken: nextToken.Type, Value: nextToken.Value}
		lexer.IncrP()
		head = head.NextExpresion
	}
	statement.Expresion = *expresion
	return &statement, nil
}
