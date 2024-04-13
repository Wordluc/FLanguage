package Evaluator

import (
	"FLanguage/Parser/Statements"
	"FLanguage/Parser/Statements/Expresions"
	"errors"
)

type VariableEnvironment struct {
	variables map[string]IObject
	externals *VariableEnvironment
}

func (v *VariableEnvironment) AddVariable(name string, value IObject) error {
	v.variables[name] = value
	return nil //check if already exists
}

func (v *VariableEnvironment) SetVariable(name string, value IObject) error {
	if v.variables[name] == nil {
		return errors.New("variable not defined")
	}
	v.variables[name] = value
	return nil
}

func (v *VariableEnvironment) GetVariable(name string) (IObject, error) {
	return v.variables[name], nil
}

func Eval(program Statements.StatementNode, env *VariableEnvironment) (IObject, error) {
	r, e := evalStatement(program.Statement, env)
	if e != nil {
		return nil, e
	}
	_, isReturn := program.Statement.(Statements.ReturnStatement)
	if isReturn {
		return r, nil
	}

	if program.Next == nil {
		return r, nil
	}
	return Eval(*program.Next, env)
}

func evalStatement(statement Statements.IStatement, env *VariableEnvironment) (IObject, error) {
	switch statement.(type) {
	case Statements.LetStatement:
		value, err := evalExpresion(statement.(Statements.LetStatement).Expresion, env)
		if err != nil {
			return nil, err
		}
		env.SetVariable(statement.(Statements.LetStatement).Identifier, value)
		ob := &LetObject{
			Name:  statement.(Statements.LetStatement).Identifier,
			Value: value,
		}
		return ob, nil
	case Statements.ReturnStatement:
		value, err := evalExpresion(statement.(Statements.ReturnStatement).Expresion, env)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
	return nil, errors.New("invalid statement")
}

func evalExpresion(expresion Expresions.IExpresion, env *VariableEnvironment) (IObject, error) {
	switch expresion.(type) {
	case Expresions.ExpresionLeaf:
	case Expresions.ExpresionNode:
	case Expresions.ExpresionCallFunc:
	}
	return nil, errors.New("invalid expresion")
}
