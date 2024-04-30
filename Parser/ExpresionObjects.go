package Parser

import (
	"FLanguage/Lexer/Token"
)

type IExpresion interface {
	ToString() string
}
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
	Type  Token.TokenType
}

func (e ExpresionLeaf) ToString() string {
	if e.Type == Token.STRING {
		return `"` + e.Value + `"`
	}
	return e.Value
}

func (_ ExpresionLeaf) New(t Token.Token) ExpresionLeaf {
	e := ExpresionLeaf{}
	e.Type = t.Type
	if e.Type == Token.STRING {
		e.Value = t.Value[1 : len(t.Value)-1]
	} else {
		e.Value = t.Value
	}
	return e
}

type ExpresionCallFunc struct {
	Values     []IExpresion
	Identifier IExpresion
}

func (e *ExpresionCallFunc) AddParm(value IExpresion) {
	e.Values = append(e.Values, value)
}
func (e ExpresionCallFunc) ToString() string {
	r := e.Identifier.ToString() + "("
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

type ExpresionGetValueArray struct {
	IndexsValues []IExpresion
	Value        IExpresion
}

func (e ExpresionGetValueArray) ToString() string {
	r := e.Value.ToString() + "["
	i := 0
	for {
		if i == len(e.IndexsValues) {
			break
		}
		r += e.IndexsValues[i].ToString() + ","
		i++
	}
	return r + "]"
}

type ExpresionDeclareArray struct {
	Values []IExpresion
}

func (e ExpresionDeclareArray) AddValue(value IExpresion) {
	e.Values = append(e.Values, value)
}
func (e ExpresionDeclareArray) ToString() string {
	r := "["
	i := 0
	for {
		if i == len(e.Values) {
			break
		}
		r += e.Values[i].ToString() + ","
		i++
	}
	r += "]"
	return r
}

type ExpresionGetValueHash struct {
	Index IExpresion
	Value IExpresion
}

func (e ExpresionGetValueHash) ToString() string {
	r := e.Value.ToString() + "{"
	r += e.Index.ToString()
	r += "}"
	return r
}

type ExpresionDeclareHash struct {
	Values map[IExpresion]IExpresion
}

func (e ExpresionDeclareHash) AddValue(value IExpresion, key IExpresion) {
	e.Values[key] = value
}
func (e ExpresionDeclareHash) ToString() string {
	r := "{"
	i := 0
	for k, v := range e.Values {
		if i == len(e.Values) {
			break
		}
		r += k.ToString() + ":" + v.ToString() + ","
		i++
	}
	r += "}"
	return r
}
