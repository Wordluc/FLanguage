package Evaluator

type IObject interface {
	Eval() (IObject, error)
}

type LetObject struct {
	Name  string
	Value IObject
}

func (l *LetObject) Eval() (IObject, error) {
	return l.Value.Eval()
}

type StringObject struct {
	Value string
}

func (s *StringObject) Eval() (IObject, error) {
	return s, nil
}

type NumberObject struct {
	Value int
}

func (n *NumberObject) Eval() (IObject, error) {
	return n, nil
}

type ReturnObject struct {
	Value IObject
}

func (r ReturnObject) Eval() (IObject, error) {
	return r, nil
}

type BoolObject struct {
	Value bool
}

func (b BoolObject) Eval() (IObject, error) {
	return b, nil
}

type ArrayObject struct {
	Values []IObject
	Type   string
}

func (a ArrayObject) Eval() (IObject, error) {
	return a, nil
}

type InnerFunc func(env *Environment) (IObject, error)

type InnerFuncObject struct {
	NameParams []string
	innerfunc  InnerFunc
}
