package Evaluator

import (
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements/Expresions"
	"errors"
	"strconv"
)

func evalExpresion(expresion Expresions.IExpresion, env *Environment) (IObject, error) {

	switch expObject := expresion.(type) {
	case Expresions.ExpresionCallFunc:
		v, e := evalCallFunc(expObject, env)
		if e != nil {
			return nil, e
		}

		return v, nil
	case Expresions.ExpresionLeaf:
		exp := expObject
		switch exp.Type {
		case Token.WORD:
			value, err := env.GetVariable(exp.Value)
			if err != nil {
				return nil, err
			}
			return value, nil
		case Token.NUMBER:
			v, _ := strconv.Atoi(exp.Value)
			ob := NumberObject{
				Value: v,
			}
			return ob, nil
		case Token.STRING:
			ob := StringObject{
				Value: exp.Value,
			}
			return ob, nil
		case Token.BOOLEAN:
			ob := BoolObject{
				Value: exp.Value == "true",
			}
			return ob, nil
		}
	case Expresions.ExpresionNode:
		left, e := evalExpresion(expObject.LeftExpresion, env)
		if e != nil {
			return nil, e
		}
		right, e := evalExpresion(expObject.RightExpresion, env)
		if e != nil {
			return nil, e
		}
		return evalBinaryExpresion(left, right, expObject.Operator)
	case Expresions.ExpresionDeclareArray:
		array := ArrayObject{}
		for _, v := range expObject.Values {
			value, e := evalExpresion(v, env)
			if e != nil {
				return nil, e
			}
			array.Values = append(array.Values, value)
		}
		return array, nil
	case Expresions.ExpresionGetValueArray:
		array, e := env.GetVariable(expObject.Name)
		if e != nil {
			return nil, e
		}
		valueId, e := evalExpresion(expObject.ValueId, env)
		if e != nil {
			return nil, e
		}
		if valueId, ok := valueId.(NumberObject); ok {
			if valueId.Value < 0 || valueId.Value >= len(array.(ArrayObject).Values) {
				return nil, errors.New("index out of range")
			}
			value := array.(ArrayObject).Values[valueId.Value]
			return value, nil
		}
		return nil, errors.New("not implemented")
	}
	return nil, errors.New("invalid expresion")
}
