package Evaluator

import (
	"FLanguage/Parser/Statements"
	"errors"
	"reflect"
)

func evalStatement(statement Statements.IStatement, env *Environment) (IObject, error) {

	switch stat := statement.(type) {
	case Statements.LetStatement:
		value, err := evalExpresion(stat.Expresion, env)

		if err != nil {
			return nil, err
		}
		e := env.AddVariable(stat.Identifier, value)
		if e != nil {
			return nil, e
		}
		ob := LetObject{
			Name:  stat.Identifier,
			Value: value,
		}
		return ob, nil
	case Statements.FuncDeclarationStatement:
		env.AddFunction(stat.Identifier, &stat)
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
		ob := ReturnObject{
			Value: value,
		}
		return ob, nil
	case Statements.AssignExpresionStatement:
		value, err := evalExpresion(stat.Expresion, env)
		if err != nil {
			return nil, err
		}
		e := env.SetVariable(stat.Identifier, value)
		if e != nil {
			return nil, e
		}
		return nil, nil
	case Statements.IfStatement:
		obCondition, e := evalExpresion(stat.Expresion, env)
		if e != nil {
			return nil, e
		}
		cond, isBool := obCondition.(BoolObject)
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

	}
	//todo: inserire l`assegnamento al array
	return nil, errors.New("invalid statement" + reflect.TypeOf(statement).String())
}
