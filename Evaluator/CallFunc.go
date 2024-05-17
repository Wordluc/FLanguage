package Evaluator

import (
	"FLanguage/Parser"
	"errors"
	"strings"
)

func evalCallFunc(expression Parser.ExpresionCallFunc, env *Environment) (iObject, error) {
	funcEnv := &Environment{
		variables:   make(map[string]iObject),
		functions:   make(map[string]Parser.FuncDeclarationStatement),
		externals:   env,
		builtInVar:  env.builtInVar,
		builtInFunc: env.builtInFunc,
	}
	var fun Parser.FuncDeclarationStatement
	var e error
	obj, e := getFunction(expression, env, funcEnv)
	if e != nil {
		return nil, e
	}
	fun, isFun := obj.(Parser.FuncDeclarationStatement)
	if !isFun {
		return obj, nil
	}
	funcEnv.addFunction(fun.Identifier, fun) //for recursion
	if len(fun.Params) > len(expression.Params) {
		return nil, errors.New("not enough parms")
	}
	evalParms(expression.Params, fun.Params, funcEnv)

	valExp, e := Eval(fun.Body.(*Parser.StatementNode), funcEnv)
	if e != nil {
		return nil, e

	}
	v, isReturn := valExp.(returnObject)
	if !isReturn {
		return nil, nil
	}
	return v.Value, nil
}

func getFunction(expression Parser.ExpresionCallFunc, env *Environment, funcEnv *Environment) (iObject, error) {
	var fun Parser.FuncDeclarationStatement
	switch ident := expression.Called.(type) {
	case Parser.ExpresionLeaf:
		funcBuiltInObject, ok := env.getBuiltInFunc(ident.Value)
		if ok == nil {
			return callBuiltInFunc(expression, funcBuiltInObject, funcEnv)
		}
		if t, e := getFunctionInLeaf(env, ident); e == nil {
			fun = t.(Parser.FuncDeclarationStatement)
		}
	case Parser.ExpresionGetValueHash:
		if t, e := manageFunctionInHashTable(ident, env, funcEnv); e == nil {
			if t != nil {
				fun = t.(Parser.FuncDeclarationStatement)
			}
		}
		funct, e := evalExpresion(expression.Called, env)
		if e != nil {
			return nil, e
		}
		if tfun, ok := funct.(Parser.FuncDeclarationStatement); ok {
			fun = tfun
			println(fun.ToString())
		}
	case Parser.FuncDeclarationStatement:
		fun = ident
	case Parser.ExpresionGetValueArray:
		if t, e := evalExpresion(expression.Called, env); e == nil {
			if tfun, ok := t.(Parser.FuncDeclarationStatement); ok {
				fun = tfun
			}
		}
	default:
		return nil, errors.New("not a function")
	}
	return fun, nil
}

func manageFunctionInHashTable(called Parser.ExpresionGetValueHash, envSource *Environment, funcEnv *Environment) (iObject, error) {
	hash, e := evalExpresion(called.Value, envSource)
	if e != nil {
		return nil, e
	}
	funcEnv.addVariable("this", hash)
	if lib, ok := hash.(libraryObject); ok {
		funcEnv.externals = lib.env
		return lib.env.getFunction(strings.Split(called.Index.ToString(), `"`)[1])
	}
	return nil, nil
}

func getFunctionInLeaf(envSource *Environment, ident Parser.ExpresionLeaf) (iObject, error) {

	fun, e := envSource.getFunction(ident.Value)
	if e != nil {
		inlineVar, e := envSource.getVariable(ident.Value)
		if e != nil {
			return nil, e
		}
		inlineFun, ok := inlineVar.(Parser.FuncDeclarationStatement)
		if !ok {
			return nil, errors.New("not a function")
		}
		return inlineFun, nil
	}
	return fun, e
}

func callBuiltInFunc(expression Parser.ExpresionCallFunc, funcBuiltInObject builtInFuncObject, env *Environment) (iObject, error) {
	err := evalParms(expression.Params, funcBuiltInObject.NameParams, env)
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
