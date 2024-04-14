package Evaluator

import (
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements"
	"FLanguage/Parser/Statements/Expresions"
	"errors"
	"reflect"
	"strconv"
)

//-la funzione ha accesso alle variabili del chiamante<--chose
//-la funzione non ha accesso alle variabili del chiamante
//-la funzione ha accesso solo ad alcune funzioni del chiamante

type Environment struct {
	variables map[string]IObject
	functions map[string]Statements.FuncDeclarationStatement
	//Environment caller
	internals *Environment //change, i want have access to calling Environment, so in a function i can use a global variable
}

func (v *Environment) AddVariable(name string, value IObject) error {
	if v.variables[name] != nil {
		return errors.New("variable already exists:" + name)
	}
	v.variables[name] = value
	return nil //check if already exists
}

func (v *Environment) SetVariable(name string, value IObject) error {
	if v.variables[name] == nil {
		return errors.New("variable not defined")
	}

	if reflect.TypeOf(v.variables[name]) != reflect.TypeOf(value) {
		return errors.New("should have same type")
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
	_, isReturn := program.Statement.(*Statements.ReturnStatement)
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
		return nil, nil
	case Statements.CallFuncStatement: //to comple
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
	}
	//todo: inserire if statement
	return nil, errors.New("invalid statement" + reflect.TypeOf(statement).String())
}

func evalCallFuncStatement(expression Expresions.ExpresionCallFunc, env *Environment) (IObject, error) {
	env.internals = &Environment{
		variables: make(map[string]IObject),
		functions: make(map[string]Statements.FuncDeclarationStatement),
	}
	fun, e := env.GetFunction(expression.NameFunc)

	if len(fun.Params) != len(expression.Values) {
		return nil, errors.New("not enough parms")
	}
	for i, v := range expression.Values {
		value, e := evalExpresion(v, env)
		if e != nil {
			return nil, e
		}
		env.internals.AddVariable(fun.Params[i], value)
	}
	if e != nil {
		return nil, e
	}
	valExp, e := Eval(fun.Body.(*Statements.StatementNode), env.internals)
	if e != nil {
		return nil, e

	}
	v, isReturn := valExp.(*ReturnObject)
	if !isReturn {
		return nil, nil
	}
	return v.Value, nil
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
		case Token.BOOLEAN:
			ob := &BoolObject{
				Value: exp.Value == "true",
			}
			return ob, nil
		}
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
		case Token.GREATER:
			return &BoolObject{Value: valueLeft > valueRight}, nil
		case Token.LESS:
			return &BoolObject{Value: valueLeft < valueRight}, nil
		case Token.EQUAL:
			return &BoolObject{Value: valueLeft == valueRight}, nil
		case Token.NOT_EQUAL:
			return &BoolObject{Value: valueLeft != valueRight}, nil
		case Token.GREATER_EQUAL:
			return &BoolObject{Value: valueLeft >= valueRight}, nil
		case Token.LESS_EQUAL:
			return &BoolObject{Value: valueLeft <= valueRight}, nil
		}
	case *StringObject:
		valueLeft := left.(*StringObject).Value
		valueRight := right.(*StringObject).Value
		switch operator.Type {
		case Token.PLUS:
			return &StringObject{valueLeft + valueRight}, nil
		case Token.EQUAL:
			return &BoolObject{valueLeft == valueRight}, nil
		case Token.NOT_EQUAL:
			return &BoolObject{valueLeft != valueRight}, nil
		}
	case *BoolObject:
		valueLeft := left.(*BoolObject).Value
		valueRight := right.(*BoolObject).Value
		if operator.Type == Token.EQUAL {
			return &BoolObject{
				Value: valueLeft == valueRight,
			}, nil
		}
	}
	return nil, errors.New("invalid operation")
}
