package Evaluator

import (
	"FLanguage/Parser/Statements"
	"FLanguage/Parser/Statements/Expresions"
	"errors"
)

func evalCallFunc(expression Expresions.ExpresionCallFunc, env *Environment) (IObject, error) {
	envFunc := &Environment{
		variables: make(map[string]IObject),
		functions: make(map[string]Statements.FuncDeclarationStatement),
		externals: env,
	}
	fun, e := env.GetFunction(expression.NameFunc)
	if e != nil {
		return nil, e
	}

	if len(fun.Params) != len(expression.Values) {
		return nil, errors.New("not enough parms")
	}
	for i, v := range expression.Values {
		value, e := evalExpresion(v, env)
		if e != nil {
			return nil, e
		}
		envFunc.AddVariable(fun.Params[i], value)
	}
	valExp, e := Eval(fun.Body.(*Statements.StatementNode), envFunc)
	if e != nil {
		return nil, e

	}
	v, isReturn := valExp.(*ReturnObject)
	if !isReturn {
		return nil, nil
	}
	return v.Value, nil
}
