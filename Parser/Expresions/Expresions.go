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
	default:
		return "(" + e.ToString() + ")"
	}
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
func (e ExpresionNode) GetWithLeft(left IExpresion) IExpresion {
	e.LeftExpresion = left
	return e
}
func (e ExpresionNode) GetWithRight(right IExpresion) IExpresion {
	e.RightExpresion = right
	return e
}
func (e ExpresionNode) GetWithOperator(operator Token.Token) IExpresion {
	e.Operator = operator
	e.OperatorValue = operator.Value
	return e
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

type ExpresionCallFunc struct {
	Values   []IExpresion
	NameFunc string
}

func (e *ExpresionCallFunc) AddParm(value IExpresion) {
	e.Values = append(e.Values, value)
}
func (e ExpresionCallFunc) ToString() string {
	r := e.NameFunc + "("
	i := 0
	for {
		if i == len(e.Values) {
			break
		}
		r += e.Values[i].ToString() + ","
		i++
	}
	r += ")"
	return r
}

type IExpresion interface {
	ToString() string
}
