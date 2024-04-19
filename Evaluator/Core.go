package Evaluator

import (
	"FLanguage/Parser/Statements"
)

func Eval(program *Statements.StatementNode, env *Environment) (IObject, error) {
	r, e := evalStatement(program.Statement, env)
	if e != nil {
		return nil, e
	}
	_, isReturn := r.(ReturnObject)
	if isReturn {
		return r, nil
	}

	if program.Next == nil {
		return r, nil
	}
	if program.Next.Statement == nil {

		return r, nil
	}
	return Eval(program.Next, env)
}
