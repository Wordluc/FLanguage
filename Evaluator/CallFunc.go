package Evaluator

import (
	"FLanguage/Parser/Statements"
	"FLanguage/Parser/Statements/Expresions"
	"errors"
)

func evalCallFunc(expression Expresions.ExpresionCallFunc, env *Environment) (iObject, error) {
	envFunc := &Environment{
		variables:   make(map[string]iObject),
		functions:   make(map[string]Statements.FuncDeclarationStatement),
		externals:   env,
		builtInVar:  env.builtInVar,
		builtInFunc: env.builtInFunc,
	}
	funcBuiltInObject, ok := env.getBuiltInFunc(expression.NameFunc)
	if ok == nil {
		err := evalParms(expression.Values, funcBuiltInObject.NameParams, envFunc)
		if err != nil {
			return nil, err
		}
		funcBuiltIn, e := funcBuiltInObject.BuiltInfunc(envFunc)
		if e != nil {
			return nil, e
		}
		return funcBuiltIn, nil
	}
	fun, e := env.getFunction(expression.NameFunc)
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
	v, isReturn := valExp.(returnObject)
	if !isReturn {
		return nil, nil
	}
	return v.Value, nil
}
func evalParms(values []Expresions.IExpresion, nameParms []string, env *Environment) error {
	for i, v := range values {

		value, e := evalExpresion(v, env)
		if e != nil {
			return e
		}
		e = env.addVariable(nameParms[i], value)
		if e != nil {
			return e
		}
	}
	return nil
}
