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
	var fun Statements.FuncDeclarationStatement
	var e error
	var ok bool
	switch ident := expression.Identifier.(type) {
	case Expresions.ExpresionLeaf:
		funcBuiltInObject, ok := env.getBuiltInFunc(ident.Value)
		if ok == nil {
			return callBuiltInFunc(expression, funcBuiltInObject, envFunc)
		}
		fun, e = env.getFunction(ident.Value)
		if e != nil {
			return nil, e
		}
	case Expresions.ExpresionGetValueHash:
		value, e := evalExpresion(ident, envFunc)
		if e != nil {
			return nil, e
		}
		fun, ok = value.(Statements.FuncDeclarationStatement)
		if !ok {
			return nil, errors.New("not a function")
		}
	case Statements.FuncDeclarationStatement:
		fun = ident
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

func callBuiltInFunc(expression Expresions.ExpresionCallFunc, funcBuiltInObject builtInFuncObject, env *Environment) (iObject, error) {
	err := evalParms(expression.Values, funcBuiltInObject.NameParams, env)
	if err != nil {
		return nil, err
	}
	funcBuiltIn, e := funcBuiltInObject.BuiltInfunc(env)
	if e != nil {
		return nil, e
	}
	return funcBuiltIn, nil

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
