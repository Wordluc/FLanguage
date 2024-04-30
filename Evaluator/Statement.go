package Evaluator

import (
	"FLanguage/Parser/Statements"
	"errors"
	"reflect"
)

func evalStatement(statement Statements.IStatement, env *Environment) (iObject, error) {

	switch stat := statement.(type) {
	case Statements.LetStatement:
		value, err := evalExpresion(stat.Expresion, env)
		if err != nil {
			return nil, err
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
	case Statements.FuncDeclarationStatement:
		env.addFunction(stat.Identifier, stat)
		return nil, nil
	case Statements.CallFuncStatement:
		value, err := evalExpresion(stat.Expresion, env)
		if err != nil {
			return nil, err
		}
		return value, nil
	case Statements.ReturnStatement:
		value, err := evalExpresion(stat.Expresion, env)
		if err != nil {
			return nil, err
		}
		ob := returnObject{
			Value: value,
		}
		return ob, nil
	case Statements.AssignExpresionStatement:
		value, err := evalExpresion(stat.Expresion, env)
		if err != nil {
			return nil, err
		}
		e := env.setVariable(stat.Identifier, value)
		if e != nil {
			return nil, e
		}
		return nil, nil
	case Statements.IfStatement:
		obCondition, e := evalExpresion(stat.Expresion, env)
		if e != nil {
			return nil, e
		}
		cond, isBool := obCondition.(boolObject)
		if !isBool {
			return nil, errors.New("invalid condition" + reflect.TypeOf(obCondition).String())
		}
		if cond.Value {
			v, e := Eval(stat.Body.(*Statements.StatementNode), env)
			return v, e
		} else {
			if stat.Else == nil {
				return nil, nil
			}
			return Eval(stat.Else.(*Statements.StatementNode), env)
		}
	case Statements.SetArrayValueStatement:
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

	case Statements.SetHashValueStatement:
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
	case Statements.WhileStatement:
		obCondition, e := evalExpresion(stat.Cond, env)
		if e != nil {
			return nil, e
		}
		cond, isBool := obCondition.(boolObject)
		if !isBool {
			return nil, errors.New("invalid condition")
		}
		for cond.Value {
			rObject, e := Eval(stat.Body.(*Statements.StatementNode), env)
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
