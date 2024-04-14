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
