package Evaluator

import (
	"FLanguage/Parser"
	"errors"
	"reflect"
)

func evalStatement(statement Parser.IStatement, env *Environment) (iObject, error) {

	switch stat := statement.(type) {
	case Parser.LetStatement:
		inlineFunc, isInlineFunc := stat.Expresion.(Parser.FuncDeclarationStatement)
		if isInlineFunc {
			inlineFunc.Identifier = stat.Identifier
			env.addFunction(stat.Identifier, inlineFunc)
			env.addVariable(stat.Identifier, inlineFunc)
			return nil, nil
		}
		value, err := evalExpresion(stat.Expresion, env)
		if err != nil {
			return nil, err
		}
		inlineFunc, isInlineFunc = value.(Parser.FuncDeclarationStatement)
		if isInlineFunc {
			inlineFunc.Identifier = stat.Identifier
			env.addFunction(stat.Identifier, inlineFunc)
		}
		e := env.addVariable(stat.Identifier, value)
		if e != nil {
			return nil, e
		}
		ob := letObject{
			Name:  stat.Identifier,
			Value: value,
		}
		return ob, nil
	case Parser.FuncDeclarationStatement:
		env.addFunction(stat.Identifier, stat)
		return nil, nil
	case Parser.CallFuncStatement:
		value, err := evalExpresion(stat.Expresion, env)
		if err != nil {
			return nil, err
		}
		return value, nil
	case Parser.ReturnStatement:
		value, err := evalExpresion(stat.Expresion, env)
		if err != nil {
			return nil, err
		}
		ob := returnObject{
			Value: value,
		}
		return ob, nil
	case Parser.AssignExpresionStatement:
		value, err := evalExpresion(stat.Expresion, env)
		if err != nil {
			return nil, err
		}
		inlineFunc, isInlineFunc := value.(Parser.FuncDeclarationStatement)
		if isInlineFunc {
			inlineFunc.Identifier = stat.Identifier
			env.addFunction(stat.Identifier, inlineFunc)
		}
		e := env.setVariable(stat.Identifier, value)
		if e != nil {
			return nil, e
		}
		return nil, nil
	case Parser.IfStatement:
		obCondition, e := evalExpresion(stat.Expresion, env)
		if e != nil {
			return nil, e
		}
		cond, isBool := obCondition.(boolObject)
		if !isBool {
			return nil, errors.New("invalid condition" + reflect.TypeOf(obCondition).String())
		}
		if cond.Value {
			v, e := Eval(stat.Body.(*Parser.StatementNode), env)
			return v, e
		} else {
			if stat.Else == nil {
				return nil, nil
			}
			return Eval(stat.Else.(*Parser.StatementNode), env)
		}
	case Parser.SetArrayValueStatement:
		exp, e := evalExpresion(stat.Value, env)
		if e != nil {
			return nil, e
		}
		array, e := env.getVariable(stat.Identifier)
		if e != nil {
			return nil, e
		}
		var elem iObject = array.(arrayObject)
		for i, idObj := range stat.Indexs {
			id, e := evalExpresion(idObj, env)
			if e != nil {
				return nil, e
			}
			index := id.(numberObject)
			if index.Value < 0 || index.Value >= len(elem.(arrayObject).Values) {
				return nil, errors.New("index out of range")
			}

			if i == len(stat.Indexs)-1 {
				elem.(arrayObject).Values[index.Value] = exp
			} else {
				elem = elem.(arrayObject).Values[index.Value]
			}
		}
		return nil, nil

	case Parser.SetHashValueStatement:
		exp, e := evalExpresion(stat.Value, env)
		if e != nil {
			return nil, e
		}
		hash, e := env.getVariable(stat.Identifier)
		if e != nil {
			return nil, e
		}
		elem, ok := hash.(hashObject)
		if !ok {
			return nil, errors.New("invalid hash")
		}
		key, e := evalExpresion(stat.Index, env)
		if e != nil {
			return nil, e
		}
		elem.Values[key] = exp

		return nil, nil
	case Parser.WhileStatement:
		obCondition, e := evalExpresion(stat.Cond, env)
		if e != nil {
			return nil, e
		}
		cond, isBool := obCondition.(boolObject)
		if !isBool {
			return nil, errors.New("invalid condition")
		}
		for cond.Value {
			rObject, e := Eval(stat.Body.(*Parser.StatementNode), env)
			if e != nil {
				return nil, e
			}
			r, isReturn := rObject.(returnObject)
			if isReturn {
				return r, nil
			}
			obCondition, e = evalExpresion(stat.Cond, env)
			if e != nil {
				return nil, e
			}
			cond, isBool = obCondition.(boolObject)
			if !isBool {
				return nil, errors.New("invalid condition")
			}
		}
		return nil, nil
	}
	return nil, errors.New("invalid statement" + reflect.TypeOf(statement).String())
}
