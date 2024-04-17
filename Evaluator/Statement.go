package Evaluator

import (
	"FLanguage/Parser/Statements"
	"errors"
	"reflect"
)

func evalStatement(statement Statements.IStatement, env *Environment) (IObject, error) {

	switch statement.(type) {
	case *Statements.LetStatement:
		value, err := evalExpresion(statement.(*Statements.LetStatement).Expresion, env)

		if err != nil {
			return nil, err
		}
		e := env.AddVariable(statement.(*Statements.LetStatement).Identifier, value)
		if e != nil {
			return nil, e
		}
		ob := &LetObject{
			Name:  statement.(*Statements.LetStatement).Identifier,
			Value: value,
		}
		return ob, nil
	case Statements.FuncDeclarationStatement:
		funcStat, _ := statement.(Statements.FuncDeclarationStatement)
		env.SetFunction(statement.(Statements.FuncDeclarationStatement).Identifier, &funcStat)
		return nil, nil
	case Statements.CallFuncStatement:
		value, err := evalExpresion(statement.(Statements.CallFuncStatement).Expresion, env)
		if err != nil {
			return nil, err
		}
		return value, nil
	case *Statements.ReturnStatement:
		value, err := evalExpresion(statement.(*Statements.ReturnStatement).Expresion, env)
		if err != nil {
			return nil, err
		}
		ob := &ReturnObject{
			Value: value,
		}
		return ob, nil
	case Statements.AssignExpresionStatement:
		value, err := evalExpresion(statement.(Statements.AssignExpresionStatement).Expresion, env)
		if err != nil {
			return nil, err
		}
		e := env.SetVariable(statement.(Statements.AssignExpresionStatement).Identifier, value)
		if e != nil {
			return nil, e
		}
		ob := &LetObject{
			Name:  statement.(Statements.AssignExpresionStatement).Identifier,
			Value: value,
		}
		return ob, nil
	case *Statements.IfStatement:
		ifStat := statement.(*Statements.IfStatement)
		obCondition, _ := evalExpresion(ifStat.Expresion, env)

		cond, isBool := obCondition.(*BoolObject)
		if !isBool {
			return nil, errors.New("invalid condition")
		}
		if cond.Value {
			v, e := Eval(ifStat.Body.(*Statements.StatementNode), env)
			return v, e
		} else {
			if ifStat.Else == nil {
				return nil, nil
			}
			return Eval(ifStat.Else.(*Statements.StatementNode), env)
		}

	}
	//todo: inserire l`assegnamento al array
	return nil, errors.New("invalid statement" + reflect.TypeOf(statement).String())
}
