package ParsingExpression

import "FLanguage/Lexer/Token"

type Expresion struct {
	TypeToken      Token.TokenType
	Value          string
	InnerExpresion *Expresion
	NextExpresion  *Expresion
}

func (E *Expresion) GetExpresion() string {
	r := E.Value
	if E.InnerExpresion != nil {
		r += "(" + E.InnerExpresion.GetExpresion() + ")"
	}
	if E.NextExpresion != nil {
		r += E.NextExpresion.GetExpresion()
	}
	return r
}
