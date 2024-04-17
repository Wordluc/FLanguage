package Evaluator

import (
	"FLanguage/Parser/Statements"
	"FLanguage/Parser/Statements/Expresions"
	"errors"
)

func evalCallFunc(expression Expresions.ExpresionCallFunc, env *Environment) (IObject, error) {
	envFunc := &Environment{
		variables: make(map[string]IObject),
		functions: make(map[string]*Statements.FuncDeclarationStatement),
		externals: env,
		innerVar:  env.innerVar,
		innerFunc: env.innerFunc,
	}
	funcInnerObject, ok := env.GetInnerFunc(expression.NameFunc)
	if ok == nil {
		err := evalParms(expression.Values, funcInnerObject.NameParams, envFunc)
		if err != nil {
			return nil, err
		}
		funcInner, e := funcInnerObject.Innerfunc(envFunc)
		if e != nil {
			return nil, e
		}
		return funcInner, nil
	}
	fun, e := env.GetFunction(expression.NameFunc)
	if e != nil {
		return nil, e
	}

	if len(fun.Params) != len(expression.Values) {
		return nil, errors.New("not enough parms")
	}
	evalParms(expression.Values, fun.Params, envFunc)

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
func evalParms(values []Expresions.IExpresion, nameParms []string, env *Environment) error {
	for i, v := range values {

		value, e := evalExpresion(v, env)
		if e != nil {
			return nil
		}
		env.AddVariable(nameParms[i], value)
	}
	return nil
}
