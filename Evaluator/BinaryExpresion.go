package Evaluator

import (
	"FLanguage/Lexer/Token"
	"errors"
	"reflect"
)

func evalBinaryExpresion(left, right iObject, operator Token.Token) (iObject, error) {
	switch op := operator.Type; op {
	case Token.PLUS:
		return evalPlusOperator(left, right)
	case Token.MINUS:
		return evalMinusOperator(left, right)
	case Token.DIV:
		return evalDivOperator(left, right)
	case Token.MULT:
		return evalMultOperator(left, right)
	case Token.EQUAL:
		return evalEqualOperator(left, right)
	case Token.NOT_EQUAL:
		v, _ := evalEqualOperator(left, right)
		return boolObject{Value: !v.(boolObject).Value}, nil
	case Token.LESS:
		return evalLessOperator(left, right)
	case Token.GREATER:
		v, _ := evalLessOperator(left, right)
		return boolObject{Value: !v.(boolObject).Value}, nil
	case Token.LESS_EQUAL:
		vEqual, _ := evalEqualOperator(left, right)
		vLess, _ := evalLessOperator(left, right)
		return boolObject{Value: vLess.(boolObject).Value || vEqual.(boolObject).Value}, nil
	case Token.GREATER_EQUAL:
		vEqual, _ := evalEqualOperator(left, right)
		vLess, _ := evalLessOperator(left, right)
		return boolObject{Value: !vLess.(boolObject).Value || vEqual.(boolObject).Value}, nil
	}
	return nil, errors.New("Wrong type")
}
func evalLessOperator(left, right iObject) (iObject, error) {
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return nil, errors.New("Wrong type")
	}
	switch leftObject := left.(type) {
	case numberObject:
		rightObject, _ := right.(numberObject)
		return boolObject{Value: leftObject.Value < rightObject.Value}, nil
	case floatNumberObject:
		rightObject, _ := right.(floatNumberObject)
		return boolObject{Value: leftObject.Value < rightObject.Value}, nil
	}
	return nil, errors.New("Wrong type")
}
func evalEqualOperator(left, right iObject) (iObject, error) {
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return boolObject{Value: false}, nil
	}
	switch leftObject := left.(type) {
	case stringObject:
		rightObject, _ := right.(stringObject)
		return boolObject{Value: leftObject.Value == rightObject.Value}, nil
	case numberObject:
		rightObject, _ := right.(numberObject)
		return boolObject{Value: leftObject.Value == rightObject.Value}, nil
	case floatNumberObject:
		rightObject, _ := right.(floatNumberObject)
		return boolObject{Value: leftObject.Value == rightObject.Value}, nil
	}
	return boolObject{Value: false}, nil
}
func evalPlusOperator(left, right iObject) (iObject, error) {
	if left == nil {
		{
			rightObject, ok := right.(floatNumberObject)
			if ok {
				return floatNumberObject{Value: rightObject.Value}, nil
			}
		}
		{
			rightObject, ok := right.(numberObject)
			if ok {
				return numberObject{Value: rightObject.Value}, nil
			}
		}
	}
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return nil, errors.New("Wrong type")
	}
	switch leftObject := left.(type) {
	case stringObject:
		rightObject, _ := right.(stringObject)
		return stringObject{Value: leftObject.Value + rightObject.Value}, nil
	case numberObject:
		rightObject, _ := right.(numberObject)
		return numberObject{Value: leftObject.Value + rightObject.Value}, nil
	case floatNumberObject:
		rightObject, _ := right.(floatNumberObject)
		return floatNumberObject{Value: leftObject.Value + rightObject.Value}, nil
	}
	return nil, errors.New("Wrong type")
}
func evalMinusOperator(left, right iObject) (iObject, error) {
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return nil, errors.New("Wrong type")
	}

	switch leftObject := left.(type) {
	case numberObject:
		rightObject, _ := right.(numberObject)
		return numberObject{Value: leftObject.Value - rightObject.Value}, nil
	case floatNumberObject:
		rightObject, _ := right.(floatNumberObject)
		return floatNumberObject{Value: leftObject.Value - rightObject.Value}, nil
	}
	return nil, errors.New("Wrong type")
}
func evalDivOperator(left, right iObject) (iObject, error) {
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return nil, errors.New("Wrong type")
	}

	switch leftObject := left.(type) {
	case numberObject:
		rightObject, _ := right.(numberObject)
		return numberObject{Value: leftObject.Value / rightObject.Value}, nil
	case floatNumberObject:
		rightObject, _ := right.(floatNumberObject)
		return floatNumberObject{Value: leftObject.Value / rightObject.Value}, nil
	}
	return nil, errors.New("Wrong type")
}
func evalMultOperator(left, right iObject) (iObject, error) {
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return nil, errors.New("Wrong type")
	}

	switch leftObject := left.(type) {
	case numberObject:
		rightObject, _ := right.(numberObject)
		return numberObject{Value: leftObject.Value * rightObject.Value}, nil

	case floatNumberObject:
		rightObject, _ := right.(floatNumberObject)
		return floatNumberObject{Value: leftObject.Value * rightObject.Value}, nil
	}
	return nil, errors.New("Wrong type")
}
