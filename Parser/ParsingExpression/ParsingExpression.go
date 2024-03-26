package ParsingExpression

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Attraction.go"
)

func ParseExpresion(f Attraction.Force, lexer *Lexer.Lexer, exp *Expresion) (*Expresion, error) {
	token, e := lexer.NextToken()
	if e != nil {
		return nil, e
	}
	lookedToken, e := lexer.LookNext()

	if e != nil {
		return nil, e
	}
	if lookedToken.Type == Token.END || lookedToken.Type == Token.DOT_COMMA {
		exp.NextExpresion = &Expresion{TypeToken: token.Type, Value: token.Value}
		lexer.IncrP()
		return exp, nil

	}
	forceNext, e := Attraction.GetForce(lookedToken.Type)
	if e != nil {
		return nil, e
	}
	if forceNext >= f {
		inerExp := &Expresion{TypeToken: token.Type, Value: token.Value}
		inerExp.NextExpresion = &Expresion{TypeToken: lookedToken.Type, Value: lookedToken.Value}
		lexer.IncrP()
		_, e = ParseExpresion(forceNext, lexer, inerExp.NextExpresion)
		exp.InnerExpresion = inerExp
		return exp, e
	}
	exp.NextExpresion = &Expresion{TypeToken: token.Type, Value: token.Value}
	//_, e = parseExpresion(f, lexer, exp.NextExpresion)
	return exp, nil
}
