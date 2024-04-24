package Evaluator

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

func newArray(env *Environment) (IObject, error) {
	typeObject, e := env.GetVariable("type")
	if e != nil {
		return nil, e
	}
	nObject, e := env.GetVariable("n")
	n, ok := nObject.(NumberObject)
	if !ok {
		return nil, errors.New("not a number")
	}
	a := ArrayObject{Values: make([]IObject, n.Value)}
	for i := 0; i < int(n.Value); i++ {
		a.Values[i] = typeObject
	}
	return a, nil
}

func builtInPrint(env *Environment) (IObject, error) {
	aObject, e := env.GetVariable("a")
	if e != nil {
		return nil, e
	}
	if aObject == nil {
		return nil, errors.New("is nil")
	}
	print(aObject.ToString())
	return nil, nil
}

func builtInPrintln(env *Environment) (IObject, error) {
	aObject, e := env.GetVariable("a")
	if e != nil {
		return nil, e
	}
	if aObject == nil {
		return nil, errors.New("is nil")
	}
	println(aObject.ToString())
	return nil, nil
}

func Int(env *Environment) (IObject, error) {
	v, e := env.GetVariable("a")
	if e != nil {
		return nil, e
	}

	switch a := v.(type) {
	case NumberObject:
		return a, nil
	case StringObject:
		n, e := strconv.Atoi(a.Value)
		if e != nil {
			return nil, errors.New("not a number")
		}
		return NumberObject{Value: n}, nil
	case BoolObject:
		if a.Value {
			return NumberObject{Value: 1}, nil
		}
		return NumberObject{Value: 0}, nil
	default:
		return nil, errors.New("not a number")
	}
}

func String(env *Environment) (IObject, error) {
	v, e := env.GetVariable("a")
	if e != nil {
		return nil, e
	}
	return StringObject{Value: v.ToString()}, nil
}

func builtInLen(env *Environment) (IObject, error) {
	aObject, e := env.GetVariable("a")
	if e != nil {
		return nil, e
	}

	switch a := aObject.(type) {
	case StringObject:
		return NumberObject{Value: len(a.Value)}, nil
	case ArrayObject:
		return NumberObject{Value: len(a.Values)}, nil

	default:
		return nil, errors.New("impossible use len")
	}
}

func Input(env *Environment) (IObject, error) {
	reader := bufio.NewReader(os.Stdin)
	v, _ := reader.ReadBytes('\n')
	return StringObject{Value: string(v[:len(v)-2])}, nil
}

func LoadBuiltInVariable(env *Environment) error {
	return nil
}

func LoadBuiltInFunction(env *Environment) error {
	env.AddBuiltInFunc("len", &BuiltInFuncObject{Name: "len", NameParams: []string{"a"}, BuiltInfunc: builtInLen})
	env.AddBuiltInFunc("newArray", &BuiltInFuncObject{Name: "newArray", NameParams: []string{"n", "type"}, BuiltInfunc: newArray})
	env.AddBuiltInFunc("int", &BuiltInFuncObject{Name: "int", NameParams: []string{"a"}, BuiltInfunc: Int})
	env.AddBuiltInFunc("string", &BuiltInFuncObject{Name: "string", NameParams: []string{"a"}, BuiltInfunc: String})
	env.AddBuiltInFunc("print", &BuiltInFuncObject{Name: "print", NameParams: []string{"a"}, BuiltInfunc: builtInPrint})
	env.AddBuiltInFunc("println", &BuiltInFuncObject{Name: "print", NameParams: []string{"a"}, BuiltInfunc: builtInPrintln})
	env.AddBuiltInFunc("read", &BuiltInFuncObject{Name: "read", NameParams: []string{}, BuiltInfunc: Input})
	return nil
}
