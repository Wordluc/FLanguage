package Evaluator

import (
	"errors"
	"reflect"
)

func newArray(env *Environment) (IObject, error) {
	typeObject, e := env.GetVariable("type")
	if e != nil {
		return nil, e
	}
	nObject, e := env.GetVariable("n")
	n, ok := nObject.(*NumberObject)
	if !ok {
		return nil, errors.New("not a number")
	}
	a := ArrayObject{Values: []IObject{}, Len: 0}
	for i := 0; i < int(n.Value); i++ {
		a.Len++
		a.Values = append(a.Values, typeObject)
	}
	return &a, nil
}
func print(env *Environment) (IObject, error) {
	aObject, e := env.GetVariable("a")
	if e != nil {
		return nil, e
	}
	switch a := aObject.(type) {
	case *StringObject:
		println(a.Value)
	case *ArrayObject:
		for _, v := range a.Values {
			println(v)
		}
	case *NumberObject:
		println(a.Value)
	case *BoolObject:
		println(a.Value)
	default:
		return nil, errors.New("not an array or string or number,got:" + reflect.TypeOf(a).String())
	}
	return nil, nil
}
func innerLen(env *Environment) (IObject, error) {
	aObject, e := env.GetVariable("a")
	if e != nil {
		return nil, e
	}
	switch a := aObject.(type) {
	case *StringObject:
		return &NumberObject{Value: (len(a.Value))}, nil
	case *ArrayObject:
		return &NumberObject{Value: (len(a.Values))}, nil
	case ArrayObject:
		return &NumberObject{Value: a.Len}, nil
	default:
		return nil, errors.New("not an array or string,got:" + reflect.TypeOf(a).String())
	}
}

var NUMBER = NumberObject{Value: 0}
var STRING = StringObject{Value: ""}

func LoadInnerVariable(env *Environment) error {
	env.AddInnerVar("NUMBER", &NUMBER)
	env.AddInnerVar("STRING", &STRING)
	return nil
}

func LoadInnerFunction(env *Environment) error {
	env.AddInnerFunc("len", &InnerFuncObject{NameParams: []string{"a"}, Innerfunc: innerLen})
	env.AddInnerFunc("newArray", &InnerFuncObject{NameParams: []string{"n", "type"}, Innerfunc: newArray})
	env.AddInnerFunc("print", &InnerFuncObject{NameParams: []string{"a"}, Innerfunc: print})
	return nil
}
