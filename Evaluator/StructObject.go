package Evaluator

type IObject interface {
}

type LetObject struct {
	Name  string
	Value IObject
}

type StringObject struct {
	Value string
}

type NumberObject struct {
	Value int
}

type ReturnObject struct {
	Value IObject
}

type BoolObject struct {
	Value bool
}

type ArrayObject struct {
	Values []IObject
	Type   string
	Len    int
}

type BuiltInFunc func(env *Environment) (IObject, error)

type BuiltInFuncObject struct {
	NameParams  []string
	BuiltInfunc BuiltInFunc
}
