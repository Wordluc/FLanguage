package Evaluator

import (
	"bufio"
	"errors"
	"os"
	"reflect"
	"strconv"
)

func newArray(env *Environment) (IObject, error) {
	typeObject, e := env.GetVariable("type") //todo: processare il type in caso volessi creare un array dentro un array
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
	if aObject == nil {
		return nil, errors.New("is nil")
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
func Int(env *Environment) (IObject, error) {
	v, e := env.GetVariable("a")
	if e != nil {
		return nil, e
	}
	switch a := v.(type) {
	case *NumberObject:
		return a, nil
	case *StringObject:
		n, e := strconv.Atoi(a.Value)
		if e != nil {
			return nil, errors.New("not a number")
		}
		return &NumberObject{Value: n}, nil
	case *BoolObject:
		if a.Value {
			return &NumberObject{Value: 1}, nil
		}
		return &NumberObject{Value: 0}, nil
	default:
		return nil, errors.New("not a number")
	}
}
func String(env *Environment) (IObject, error) {
	v, e := env.GetVariable("a")
	if e != nil {
		return nil, e
	}
	switch a := v.(type) {
	case *NumberObject:
		return &StringObject{Value: strconv.Itoa(a.Value)}, nil
	case *StringObject:
		return a, nil
	case *BoolObject:
		if a.Value {
			return &StringObject{Value: "true"}, nil
		}
		return &StringObject{Value: "false"}, nil
	default:
		return nil, errors.New("not a string")
	}
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
func Input(env *Environment) (IObject, error) {
	reader := bufio.NewReader(os.Stdin)
	v, _ := reader.ReadBytes('\n')
	return &StringObject{Value: string(v[:len(v)-2])}, nil
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
	env.AddInnerFunc("int", &InnerFuncObject{NameParams: []string{"a"}, Innerfunc: Int})
	env.AddInnerFunc("string", &InnerFuncObject{NameParams: []string{"a"}, Innerfunc: String})
	env.AddInnerFunc("print", &InnerFuncObject{NameParams: []string{"a"}, Innerfunc: print})
	env.AddInnerFunc("read", &InnerFuncObject{NameParams: []string{}, Innerfunc: Input})
	return nil
}
