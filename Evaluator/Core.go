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
	externals *Environment
}

func (v *Environment) AddVariable(name string, value IObject) error {
	_, exist := v.variables[name]
	if exist {
		return errors.New("variable already exists:" + name)
	}
	v.variables[name] = value
	return nil
}

func (v *Environment) SetVariable(name string, value IObject) error {
	variable, exist := v.variables[name]
	if !exist {
		return errors.New("variable not defined")
	}

	if reflect.TypeOf(variable) != reflect.TypeOf(value) {
		return errors.New("should have same type")
	}
	v.variables[name] = value
	return nil
}

func (v *Environment) GetVariable(name string) (IObject, error) {
	variable, exist := v.variables[name]
	if !exist {
		variable, existEx := v.externals.GetVariable(name)
		if existEx != nil {
			return nil, errors.New("variable not defined")
		}
		return variable, nil
	}
	return variable, nil
}

func (v *Environment) GetFunction(name string) (Statements.FuncDeclarationStatement, error) {
	funct, exist := v.functions[name]
	if !exist {
		funct, existEx := v.externals.GetFunction(name)
		if existEx != nil {
			return Statements.FuncDeclarationStatement{}, errors.New("function not defined")
		}
		return funct, nil
	}
	return funct, nil
}

func (v *Environment) SetFunction(name string, value Statements.FuncDeclarationStatement) error {
	_, exist := v.functions[name]
	if exist {
		return errors.New("function already exists:" + name)
	}
	v.functions[name] = value
	return nil

}

func Eval(program *Statements.StatementNode, env *Environment) (IObject, error) {
	r, e := evalStatement(program.Statement, env)
	if e != nil {
		return nil, e
	}
	_, isReturn := r.(*ReturnObject)
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

func evalCallFuncStatement(expression Expresions.ExpresionCallFunc, env *Environment) (IObject, error) {
	envFunc := &Environment{
		variables: make(map[string]IObject),
		functions: make(map[string]Statements.FuncDeclarationStatement),
		externals: env,
	}
	fun, e := env.GetFunction(expression.NameFunc)
	if e != nil {
		return nil, e
	}

	if len(fun.Params) != len(expression.Values) {
		return nil, errors.New("not enough parms")
	}
	for i, v := range expression.Values {
		value, e := evalExpresion(v, env)
		if e != nil {
			return nil, e
		}
		envFunc.AddVariable(fun.Params[i], value)
	}
	valExp, e := Eval(fun.Body.(*Statements.StatementNode), envFunc)
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
		return evalBinaryExpresion(left, right, expresion.(Expresions.ExpresionNode).Operator)
		//case Expresions.ExpresionCallFunc:
	case *Expresions.ExpresionDeclareArray:
		exp := expresion.(*Expresions.ExpresionDeclareArray)
		array := ArrayObject{}
		for _, v := range exp.Values {
			value, e := evalExpresion(v, env)
			if e != nil {
				return nil, e
			}
			array.Values = append(array.Values, value)
		}
		return array, nil
	case *Expresions.ExpresionGetValueArray:
		exp := expresion.(*Expresions.ExpresionGetValueArray)
		array, e := env.GetVariable(exp.Name)
		if e != nil {
			return nil, e
		}
		valueId, e := evalExpresion(exp.ValueId, env)
		if e != nil {
			return nil, e
		}
		if valueId, ok := valueId.(*NumberObject); ok {
			return array.(ArrayObject).Values[valueId.Value], nil
		}
		return nil, errors.New("not implemented")
	}
	return nil, errors.New("invalid expresion")
}

func evalBinaryExpresion(left, right IObject, operator Token.Token) (IObject, error) {

	switch left.(type) {
	case *NumberObject:
		valueLeft := left.(*NumberObject).Value
		stringValue, isRightString := right.(*StringObject)
		if isRightString {
			if operator.Type == Token.PLUS || stringValue != nil {
				return &StringObject{strconv.Itoa(valueLeft) + stringValue.Value}, nil
			} else {
				return nil, errors.New("invalid operator")
			}
		}
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
		stringValue, isRightNumber := right.(*NumberObject)
		if isRightNumber {
			if operator.Type == Token.PLUS || stringValue != nil {
				return &StringObject{valueLeft + strconv.Itoa(stringValue.Value)}, nil
			} else {
				return nil, errors.New("invalid operation")
			}
		}
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
		_, isRightBool := right.(*BoolObject)
		if !isRightBool {
			return nil, errors.New("invalid operation")
		}
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
