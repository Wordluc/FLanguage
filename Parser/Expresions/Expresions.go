package Expresions

import (
	"FLanguage/Lexer/Token"
	"reflect"
)

type BinaryExpresion struct {
	ValueA        string
	TypeA         Token.TokenType
	Operator      Token.Token
	ValueOperator string
	NextExpresion IExpresion
}

func (e BinaryExpresion) GetNextExpresion() IExpresion {
	return e.NextExpresion
}
func (e BinaryExpresion) SetNextExpresion(exp IExpresion) IExpresion {
	e.NextExpresion = exp
	return e
}

func (e BinaryExpresion) GetString() string {
	r := e.ValueA + e.ValueOperator
	if e.NextExpresion != nil && reflect.TypeOf(e.NextExpresion).Name() != "EmptyExpresion" {
		r += "[" + e.NextExpresion.GetString() + "]"
	}
	return r
}

func (e BinaryExpresion) New(tokenA, tokenOp Token.Token) BinaryExpresion {
	e.ValueA = tokenA.Value
	e.TypeA = tokenA.Type
	e.Operator = tokenOp
	e.ValueOperator = tokenOp.Value
	e.NextExpresion = EmptyExpresion{}
	return e
}
func (e BinaryExpresion) NewValue(tokenA Token.Token) BinaryExpresion {
	e.ValueA = tokenA.Value
	e.TypeA = tokenA.Type
	e.NextExpresion = EmptyExpresion{}
	return e
}

type OperatorExpresion struct {
	Operator Token.Token
}

func (e OperatorExpresion) GetNextExpresion() IExpresion {
	return nil
}
func (e OperatorExpresion) SetNextExpresion(IExpresion) IExpresion {
	return e
}
func (e OperatorExpresion) GetString() string {
	return e.Operator.Value
}

type MainExpresion struct {
	Operator      Token.Token
	Expresion     IExpresion
	NextExpresion *MainExpresion
}

func (e MainExpresion) GetNextExpresion() IExpresion {
	return e.NextExpresion
}
func (e MainExpresion) SetNextExpresion(exp IExpresion) IExpresion {
	e.NextExpresion = &MainExpresion{Expresion: exp} //TODO: verificare l'utilità di usare la stessa interface
	return e
}
func (e MainExpresion) GetString() string {
	r := "{" + e.Expresion.GetString() + "}"
	if e.NextExpresion != nil {
		r += e.NextExpresion.GetString()
	}
	return r
}

type EmptyExpresion struct {
}

func (e EmptyExpresion) GetNextExpresion() IExpresion {
	return nil
}
func (e EmptyExpresion) SetNextExpresion(IExpresion) IExpresion {
	return e
}
func (e EmptyExpresion) GetString() string {
	return ""
}

type IExpresion interface {
	GetNextExpresion() IExpresion
	SetNextExpresion(IExpresion) IExpresion
	GetString() string
}
