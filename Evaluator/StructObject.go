package Evaluator

import "strconv"

type IObject interface {
	ToString() string
}

type LetObject struct {
	Name  string
	Value IObject
}

func (l LetObject) ToString() string {
	return "let " + l.Name + " = " + l.Value.ToString()
}

type StringObject struct {
	Value string
}

func (s StringObject) ToString() string {
	return s.Value
}

type NumberObject struct {
	Value int
}

func (n NumberObject) ToString() string {
	return strconv.Itoa(n.Value)
}

type ReturnObject struct {
	Value IObject
}

func (r ReturnObject) ToString() string {
	return r.Value.ToString()
}

type BoolObject struct {
	Value bool
}

func (b BoolObject) ToString() string {
	if b.Value {
		return "true"
	}
	return "false"
}

type ArrayObject struct {
	Values []IObject
	Type   string
}

func (a ArrayObject) ToString() string {
	r := "["
	for _, v := range a.Values {
		r += v.ToString() + ","
	}
	return r + "]"
}

type BuiltInFunc func(env *Environment) (IObject, error)

type BuiltInFuncObject struct {
	Name        string
	NameParams  []string
	BuiltInfunc BuiltInFunc
}

func (b BuiltInFuncObject) ToString() string {
	r := b.Name + "("
	for _, v := range b.NameParams {
		r += v + ","
	}
	return r + ")"
}
