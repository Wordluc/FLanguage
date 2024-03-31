package Expresions

import "FLanguage/Lexer/Token"

type ExpresionNode struct {
	LeftExpresion  IExpresion
	Operator       Token.Token
	OperatorValue  string
	RightExpresion IExpresion
}

func (e ExpresionNode) ToString() string {
	r := ""
	if e.LeftExpresion != nil {
		r += PrintLeafOrExpresion(e.LeftExpresion)
	}
	r += " " + e.Operator.Value + " "
	if e.RightExpresion != nil {
		r += PrintLeafOrExpresion(e.RightExpresion)
	}
	return r
}
func PrintLeafOrExpresion(e IExpresion) string {
	switch e.(type) {
	case ExpresionLeaf:
		return e.ToString()
	case ExpresionNode:
		return "(" + e.ToString() + ")"
	}
	return ""
}
func (e *ExpresionNode) SetLeft(left IExpresion) {
	e.LeftExpresion = left
}
func (e *ExpresionNode) SetRight(right IExpresion) {
	e.RightExpresion = right
}
func (e *ExpresionNode) SetOperator(operator Token.Token) {
	e.Operator = operator
	e.OperatorValue = operator.Value
}

type ExpresionLeaf struct {
	Value string
	Type  Token.Token
}

func (e ExpresionLeaf) ToString() string {
	return e.Value
}
func (_ ExpresionLeaf) New(t Token.Token) ExpresionLeaf {
	e := ExpresionLeaf{}
	e.Type = t
	e.Value = t.Value
	return e
}

type IExpresion interface {
	ToString() string
}
