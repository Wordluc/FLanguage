package Evaluator

import (
	"FLanguage/Parser"
	"errors"
	"strings"
)

func evalCallFunc(expression Parser.ExpresionCallFunc, env *Environment) (iObject, error) {
	envFunc := &Environment{
		variables:   make(map[string]iObject),
		functions:   make(map[string]Parser.FuncDeclarationStatement),
		externals:   env,
		builtInVar:  env.builtInVar,
		builtInFunc: env.builtInFunc,
	}
	var fun Parser.FuncDeclarationStatement
	var e error
	switch ident := expression.Identifier.(type) {
	case Parser.ExpresionLeaf:
		funcBuiltInObject, ok := env.getBuiltInFunc(ident.Value)
		if ok == nil {
			return callBuiltInFunc(expression, funcBuiltInObject, envFunc)
		}
		fun, e = env.getFunction(ident.Value)
		if e != nil {
			inlineVar, e := env.getVariable(ident.Value)
			if e != nil {
				return nil, e
			}
			inlineFun, ok := inlineVar.(Parser.FuncDeclarationStatement)
			if !ok {
				return nil, errors.New("not a function")
			}
			fun = inlineFun
		}
	default:
		hashGet, ok := expression.Identifier.(Parser.ExpresionGetValueHash)
		if ok {
			hash, e := evalExpresion(hashGet.Value, env)
			if e != nil {
				return nil, e
			}
			if lib, ok := hash.(libraryObject); ok {
				envFunc.functions = lib.env.functions
				fun, e = lib.env.getFunction(strings.Split(hashGet.Index.ToString(), `"`)[1])
			}
			envFunc.addVariable("this", hash)
		}
		funct, e := evalExpresion(expression.Identifier, env)
		if e != nil {
			return nil, e
		}
		tfun := fun
		fun, ok = funct.(Parser.FuncDeclarationStatement)
		if !ok {
			fun = tfun
		}
	}
	if len(fun.Params) != len(expression.Values) {
		return nil, errors.New("not enough parms")
	}
	evalParms(expression.Values, fun.Params, envFunc)

	valExp, e := Eval(fun.Body.(*Parser.StatementNode), envFunc)
	if e != nil {
		return nil, e

	}
	v, isReturn := valExp.(returnObject)
	if !isReturn {
		return nil, nil
	}
	return v.Value, nil
}

func callBuiltInFunc(expression Parser.ExpresionCallFunc, funcBuiltInObject builtInFuncObject, env *Environment) (iObject, error) {
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
func evalParms(values []Parser.IExpresion, nameParms []string, env *Environment) error {
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
