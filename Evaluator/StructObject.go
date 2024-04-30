package Evaluator

import "strconv"

type iObject interface {
	ToString() string
}

type letObject struct {
	Name  string
	Value iObject
}

func (l letObject) ToString() string {
	return "let " + l.Name + " = " + l.Value.ToString()
}

type stringObject struct {
	Value string
}

func (s stringObject) ToString() string {
	return s.Value
}

type numberObject struct {
	Value int
}

func (n numberObject) ToString() string {
	return strconv.Itoa(n.Value)
}

type floatNumberObject struct {
	Value float64
}

func (n floatNumberObject) ToString() string {
	return strconv.FormatFloat(n.Value, 'f', -1, 32)
}

type returnObject struct {
	Value iObject
}

func (r returnObject) ToString() string {
	return r.Value.ToString()
}

type boolObject struct {
	Value bool
}

func (b boolObject) ToString() string {
	if b.Value {
		return "true"
	}
	return "false"
}

type arrayObject struct {
	Values []iObject
	Type   string
}

func (a arrayObject) ToString() string {
	r := "["
	for i, v := range a.Values {
		r += v.ToString()
		if i < len(a.Values)-1 {
			r += ","
		}
	}
	return r + "]"
}

type hashObject struct {
	Values map[iObject]iObject
	Type   string
}

func (a hashObject) ToString() string {
	r := "{"
	for i, v := range a.Values {
		r += i.ToString() + ":" + v.ToString() + ","
	}
	return r + "}"
}

type builtInFunc func(env *Environment) (iObject, error)

type builtInFuncObject struct {
	Name        string
	NameParams  []string
	BuiltInfunc builtInFunc
}

func (b builtInFuncObject) ToString() string {
	r := b.Name + "("
	for _, v := range b.NameParams {
		r += v + ","
	}
	return r + ")"
}
