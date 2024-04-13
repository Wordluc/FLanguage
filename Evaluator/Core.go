package Evaluator

import (
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements"
	"FLanguage/Parser/Statements/Expresions"
	"errors"
	"reflect"
	"strconv"
)

type Environment struct {
	variables map[string]IObject
	functions map[string]Statements.FuncDeclarationStatement
	internals *Environment
}

func (v *Environment) AddVariable(name string, value IObject) error {
	v.variables[name] = value
	return nil //check if already exists
}

func (v *Environment) SetVariable(name string, value IObject) error {
	if v.variables[name] == nil {
		return errors.New("variable not defined")
	}
	v.variables[name] = value
	return nil
}

func (v *Environment) GetVariable(name string) (IObject, error) {
	return v.variables[name], nil
}

func (v *Environment) GetFunction(name string) (Statements.FuncDeclarationStatement, error) {
	return v.functions[name], nil
}

func (v *Environment) SetFunction(name string, value Statements.FuncDeclarationStatement) error {
	v.functions[name] = value
	return nil //check if already exists
}

func Eval(program *Statements.StatementNode, env *Environment) (IObject, error) {
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
	if program.Next.Statement == nil {
		return r, nil
	}
	return Eval(program.Next, env)
}

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
		env.SetFunction(statement.(Statements.FuncDeclarationStatement).Identifier, statement.(Statements.FuncDeclarationStatement))
		return &ReturnObject{}, nil
	case Statements.CallFuncStatement: //to comple
		value, err := evalExpresion(statement.(Statements.CallFuncStatement).Expresion, env)
		if err != nil {
			return nil, err
		}

		return value, nil
	case *Statements.ReturnStatement:
		value, err := evalExpresion(statement.(Statements.ReturnStatement).Expresion, env)
		if err != nil {
			return nil, err
		}
		return value, nil
	case Statements.AssignExpresionStatement:
		value, err := evalExpresion(statement.(Statements.AssignExpresionStatement).Expresion, env)
		if err != nil {
			return nil, err
		}
		env.SetVariable(statement.(Statements.AssignExpresionStatement).Identifier, value)
		ob := &LetObject{
			Name:  statement.(Statements.AssignExpresionStatement).Identifier,
			Value: value,
		}
		return ob, nil
	}
	return nil, errors.New("invalid statement" + reflect.TypeOf(statement).String())
}

func evalCallFuncStatement(expression Expresions.ExpresionCallFunc, env *Environment) (IObject, error) {
	env.internals = &Environment{
		variables: make(map[string]IObject),

		functions: make(map[string]Statements.FuncDeclarationStatement),
	}
	for i, v := range expression.Values {
		value, e := evalExpresion(v, env)
		if e != nil {
			return nil, e
		}
		env.internals.AddVariable(strconv.Itoa(i), value)
	}

	return Eval(env.functions[expression.NameFunc].Body.(*Statements.StatementNode), env.internals) //env.functions[expCallFunc.NameFunc]
}

func evalExpresion(expresion Expresions.IExpresion, env *Environment) (IObject, error) {
	switch expresion.(type) {
	case *Expresions.ExpresionCallFunc:
		v, e := evalCallFuncStatement(*expresion.(*Expresions.ExpresionCallFunc), env)
		if e != nil {
			return nil, e
		}

		return v, nil
	case Expresions.ExpresionLeaf:
		exp := expresion.(Expresions.ExpresionLeaf)
		switch exp.Type {
		case Token.WORD:
			value, err := env.GetVariable(exp.Value)
			if err != nil {
				return nil, err
			}
			return value, nil
		case Token.NUMBER:
			v, _ := strconv.Atoi(exp.Value)
			ob := &NumberObject{
				Value: v,
			}
			return ob, nil
		case Token.STRING:
			ob := &StringObject{
				Value: exp.Value,
			}
			return ob, nil

		} //insert callfunc

	case Expresions.ExpresionNode:
		left, e := evalExpresion(expresion.(Expresions.ExpresionNode).LeftExpresion, env)
		if e != nil {
			return nil, e
		}
		right, e := evalExpresion(expresion.(Expresions.ExpresionNode).RightExpresion, env)
		if e != nil {
			return nil, e
		}
		typeLeft := reflect.TypeOf(left)
		typeRight := reflect.TypeOf(right)
		if typeLeft != typeRight {
			return nil, errors.New("invalid operation")
		}
		return evalBinaryExpresion(left, right, expresion.(Expresions.ExpresionNode).Operator)
		//case Expresions.ExpresionCallFunc:
	}
	return nil, errors.New("invalid expresion")
}

func evalBinaryExpresion(left, right IObject, operator Token.Token) (IObject, error) {
	switch left.(type) {
	case *NumberObject:
		valueLeft := left.(*NumberObject).Value
		valueRight := right.(*NumberObject).Value
		switch operator.Type {
		case Token.PLUS:
			return &NumberObject{valueLeft + valueRight}, nil
		case Token.MINUS:
			return &NumberObject{valueLeft - valueRight}, nil
		case Token.DIV:
			return &NumberObject{valueLeft / valueRight}, nil
		case Token.MULT:
			return &NumberObject{valueLeft * valueRight}, nil
		}
	case *StringObject:
		valueLeft := left.(*StringObject).Value
		valueRight := right.(*StringObject).Value
		switch operator.Type {
		case Token.PLUS:
			return &StringObject{valueLeft + valueRight}, nil
		}
	}
	return nil, errors.New("invalid operation")
}
