package Evaluator

import (
	"FLanguage/Lexer/Token"
	"FLanguage/Parser"
	"errors"
	"strconv"
)

func evalExpresion(expresion Parser.IExpresion, env *Environment) (iObject, error) {
	switch expObject := expresion.(type) {
	case Parser.ExpresionCallFunc:
		v, e := evalCallFunc(expObject, env)
		if e != nil {
			return nil, e
		}

		return v, nil
	case Parser.ExpresionLeaf:
		exp := expObject
		switch exp.Type {
		case Token.WORD:
			value, err := env.getVariable(exp.Value)
			if err != nil {
				return nil, err
			}
			return value, nil
		case Token.NUMBER:
			v, _ := strconv.Atoi(exp.Value)
			ob := numberObject{
				Value: v,
			}
			return ob, nil
		case Token.NUMBER_WITH_DOT:
			v, _ := strconv.ParseFloat(exp.Value, 32)
			ob := floatNumberObject{
				Value: v,
			}
			return ob, nil
		case Token.STRING:
			ob := stringObject{
				Value: exp.Value,
			}
			return ob, nil
		case Token.BOOLEAN:
			ob := boolObject{
				Value: exp.Value == "true",
			}
			return ob, nil
		}
	case Parser.FuncDeclarationStatement:
		return expObject, nil
	case Parser.ExpresionNode:
		var left iObject
		var e error
		right, e := evalExpresion(expObject.RightExpresion, env)
		if e != nil {
			return nil, e
		}
		if expObject.LeftExpresion == nil {
			switch right.(type) {
			case numberObject:
				left = numberObject{
					Value: 0,
				}
			case floatNumberObject:
				left = floatNumberObject{
					Value: 0.0,
				}
			}
		} else {
			left, e = evalExpresion(expObject.LeftExpresion, env)
			if e != nil {
				return nil, e
			}
		}

		return evalBinaryExpresion(left, right, expObject.Operator)
	case Parser.ExpresionDeclareArray:
		array := arrayObject{}
		array.Values = make([]iObject, len(expObject.Values))
		for i, v := range expObject.Values {
			value, e := evalExpresion(v, env)
			if e != nil {
				return nil, e
			}
			array.Values[i] = value
		}
		return array, nil
	case Parser.ExpresionDeclareHash:
		hash := hashObject{}
		hash.Values = make(map[iObject]iObject)
		for i, v := range expObject.Values {
			key, e := evalExpresion(i, env)
			if e != nil {
				return nil, e
			}
			value, e := evalExpresion(v, env)
			if e != nil {
				return nil, e
			}
			hash.Values[key] = value
		}
		return hash, nil
	case Parser.ExpresionGetValueHash:
		hash, e := evalExpresion(expObject.Value, env)
		if e != nil {
			return nil, e
		}
		key, e := evalExpresion(expObject.Index, env)
		if e != nil {
			return nil, e
		}
		if b, ok := hash.(libraryObject); ok {
			return b, nil
		}
		hash, ok := hash.(hashObject).Values[key]

		if !ok {
			return nil, errors.New("element not found")
		}
		return hash, nil

	case Parser.ExpresionGetValueArray:
		value, e := evalExpresion(expObject.Value, env)
		if e != nil {
			return nil, e
		}

		switch v := value.(type) {
		case arrayObject: //From array
			var elem iObject = v
			for _, idObj := range expObject.IndexsValues {
				id, e := evalExpresion(idObj, env)
				if e != nil {
					return nil, e
				}
				index := id.(numberObject)
				if index.Value < 0 || index.Value >= len(elem.(arrayObject).Values) {
					return nil, errors.New("index out of range")
				}
				elem = elem.(arrayObject).Values[index.Value]
			}
			return elem, nil
		case stringObject: //From string
			id, e := evalExpresion(expObject.IndexsValues[0], env)
			if e != nil {
				return nil, e
			}
			index := id.(numberObject)
			if index.Value < 0 || index.Value >= len(v.Value) {
				return nil, errors.New("index out of range")
			}
			return stringObject{Value: string(v.Value[index.Value])}, nil
		}
		return nil, errors.New("invalid get expresion")

	}
	return nil, errors.New("invalid expresion")
}
