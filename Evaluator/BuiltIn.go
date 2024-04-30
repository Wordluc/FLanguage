package Evaluator

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strconv"
)

func newArrayBuiltIn(env *Environment) (iObject, error) {
	typeObject, e := env.getVariable("type")
	if e != nil {
		return nil, e
	}
	nObject, e := env.getVariable("n")
	n, ok := nObject.(numberObject)
	if !ok {
		return nil, errors.New("not a number")
	}
	a := arrayObject{Values: make([]iObject, n.Value)}
	for i := 0; i < int(n.Value); i++ {
		a.Values[i] = typeObject
	}
	return a, nil
}

func printBuiltIn(env *Environment) (iObject, error) {
	aObject, e := env.getVariable("a")
	if e != nil {
		return nil, e
	}
	if aObject == nil {
		return nil, errors.New("is nil")
	}

	print(aObject.ToString())
	return nil, nil
}

func printlnBuiltIn(env *Environment) (iObject, error) {
	aObject, e := env.getVariable("a")
	if e != nil {
		return nil, e
	}
	if aObject == nil {
		return nil, errors.New("is nil")
	}
	println(aObject.ToString())
	return nil, nil
}

func intBuiltIn(env *Environment) (iObject, error) {
	v, e := env.getVariable("a")
	if e != nil {
		return nil, e
	}

	switch a := v.(type) {
	case floatNumberObject:
		return numberObject{Value: int(a.Value)}, nil
	case numberObject:
		return a, nil
	case stringObject:
		n, e := strconv.Atoi(a.Value)
		if e != nil {
			return nil, errors.New("not a number")
		}
		return numberObject{Value: n}, nil
	case boolObject:
		if a.Value {
			return numberObject{Value: 1}, nil
		}
		return numberObject{Value: 0}, nil
	default:
		return nil, errors.New("not a number")
	}
}
func floatBuiltIn(env *Environment) (iObject, error) {
	v, e := env.getVariable("a")
	if e != nil {
		return nil, e
	}

	switch a := v.(type) {
	case numberObject:
		return floatNumberObject{Value: float64(a.Value)}, nil
	case stringObject:
		n, e := strconv.ParseFloat(a.Value, 32)
		if e != nil {
			return nil, errors.New("not a number")
		}
		return floatNumberObject{Value: n}, nil
	case boolObject:
		if a.Value {
			return floatNumberObject{Value: 1}, nil
		}
		return floatNumberObject{Value: 0}, nil
	default:
		return nil, errors.New("not a number")
	}
}
func stringBuiltIn(env *Environment) (iObject, error) {
	v, e := env.getVariable("a")
	if e != nil {
		return nil, e
	}
	return stringObject{Value: v.ToString()}, nil
}

func lenBuiltin(env *Environment) (iObject, error) {
	aObject, e := env.getVariable("a")
	if e != nil {
		return nil, e
	}

	switch a := aObject.(type) {
	case stringObject:
		return numberObject{Value: len(a.Value)}, nil
	case arrayObject:
		return numberObject{Value: len(a.Values)}, nil

	default:
		return nil, errors.New("impossible use len")
	}
}

func inputBuiltIn(env *Environment) (iObject, error) {
	reader := bufio.NewReader(os.Stdin)
	v, _ := reader.ReadBytes('\n')
	return stringObject{Value: string(v[:len(v)-2])}, nil
}

func LoadBuiltInVariable(env *Environment) error {
	return nil
}
func ImportLibrary(env *Environment) (iObject, error) {
	pathVar, e := env.getVariable("path")
	if e != nil {
		return nil, e
	}
	pathOb, ok := pathVar.(stringObject)
	if !ok {
		return nil, errors.New("not a string")
	}
	dirtyPath, e := os.Getwd()
	if e != nil {
		return nil, e
	}
	pwd := filepath.Dir(dirtyPath)
	path := filepath.Join(pwd, "Library", pathOb.Value)
	_, e = Run(path, env)
	if e != nil {
		return nil, e
	}
	if len(env.variables) > 1 {
		return nil, errors.New("not possible define variables in library")
	}
	for name, funct := range env.functions {
		env.externals.addFunction(name, funct)
	}
	return nil, nil
}
func LoadBuiltInFunction(env *Environment) {
	env.addBuiltInFunc("len", builtInFuncObject{Name: "len", NameParams: []string{"a"}, BuiltInfunc: lenBuiltin})
	env.addBuiltInFunc("newArray", builtInFuncObject{Name: "newArray", NameParams: []string{"n", "type"}, BuiltInfunc: newArrayBuiltIn})
	env.addBuiltInFunc("int", builtInFuncObject{Name: "int", NameParams: []string{"a"}, BuiltInfunc: intBuiltIn})
	env.addBuiltInFunc("float", builtInFuncObject{Name: "float", NameParams: []string{"a"}, BuiltInfunc: floatBuiltIn})
	env.addBuiltInFunc("string", builtInFuncObject{Name: "string", NameParams: []string{"a"}, BuiltInfunc: stringBuiltIn})
	env.addBuiltInFunc("print", builtInFuncObject{Name: "print", NameParams: []string{"a"}, BuiltInfunc: printBuiltIn})
	env.addBuiltInFunc("println", builtInFuncObject{Name: "print", NameParams: []string{"a"}, BuiltInfunc: printlnBuiltIn})
	env.addBuiltInFunc("read", builtInFuncObject{Name: "read", NameParams: []string{}, BuiltInfunc: inputBuiltIn})
	env.addBuiltInFunc("import", builtInFuncObject{Name: "import", NameParams: []string{"path"}, BuiltInfunc: ImportLibrary})
}
