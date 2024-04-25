package Evaluator

import (
	"FLanguage/Lexer/Token"
	"errors"
	"reflect"
)

func evalBinaryExpresion(left, right IObject, operator Token.Token) (IObject, error) {
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
		return BoolObject{Value: !v.(BoolObject).Value}, nil
	case Token.LESS:
		return evalLessOperator(left, right)
	case Token.GREATER:
		v, _ := evalLessOperator(left, right)
		return BoolObject{Value: !v.(BoolObject).Value}, nil
	case Token.LESS_EQUAL:
		vEqual, _ := evalEqualOperator(left, right)
		vLess, _ := evalLessOperator(left, right)
		return BoolObject{Value: vLess.(BoolObject).Value || vEqual.(BoolObject).Value}, nil
	case Token.GREATER_EQUAL:
		vEqual, _ := evalEqualOperator(left, right)
		vLess, _ := evalLessOperator(left, right)
		return BoolObject{Value: !vLess.(BoolObject).Value || vEqual.(BoolObject).Value}, nil
	}
	return nil, errors.New("Wrong type")
}
func evalLessOperator(left, right IObject) (IObject, error) {
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return nil, errors.New("Wrong type")
	}
	switch leftObject := left.(type) {
	case NumberObject:
		rightObject, _ := right.(NumberObject)
		return BoolObject{Value: leftObject.Value < rightObject.Value}, nil
	case FloatNumberObject:
		rightObject, _ := right.(FloatNumberObject)
		return BoolObject{Value: leftObject.Value < rightObject.Value}, nil
	}
	return nil, errors.New("Wrong type")
}
func evalEqualOperator(left, right IObject) (IObject, error) {
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return BoolObject{Value: false}, nil
	}
	switch leftObject := left.(type) {
	case StringObject:
		rightObject, _ := right.(StringObject)
		return BoolObject{Value: leftObject.Value == rightObject.Value}, nil
	case NumberObject:
		rightObject, _ := right.(NumberObject)
		return BoolObject{Value: leftObject.Value == rightObject.Value}, nil
	case FloatNumberObject:
		rightObject, _ := right.(FloatNumberObject)
		return BoolObject{Value: leftObject.Value == rightObject.Value}, nil
	}
	return BoolObject{Value: false}, nil
}
func evalPlusOperator(left, right IObject) (IObject, error) {
	if left == nil {
		{
			rightObject, ok := right.(FloatNumberObject)
			if ok {
				return FloatNumberObject{Value: rightObject.Value}, nil
			}
		}
		{
			rightObject, ok := right.(NumberObject)
			if ok {
				return NumberObject{Value: rightObject.Value}, nil
			}
		}
	}
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return nil, errors.New("Wrong type")
	}
	switch leftObject := left.(type) {
	case StringObject:
		rightObject, _ := right.(StringObject)
		return StringObject{Value: leftObject.Value + rightObject.Value}, nil
	case NumberObject:
		rightObject, _ := right.(NumberObject)
		return NumberObject{Value: leftObject.Value + rightObject.Value}, nil
	case FloatNumberObject:
		rightObject, _ := right.(FloatNumberObject)
		return FloatNumberObject{Value: leftObject.Value + rightObject.Value}, nil
	}
	return nil, errors.New("Wrong type")
}
func evalMinusOperator(left, right IObject) (IObject, error) {
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return nil, errors.New("Wrong type")
	}

	switch leftObject := left.(type) {
	case NumberObject:
		rightObject, _ := right.(NumberObject)
		return NumberObject{Value: leftObject.Value - rightObject.Value}, nil
	case FloatNumberObject:
		rightObject, _ := right.(FloatNumberObject)
		return FloatNumberObject{Value: leftObject.Value - rightObject.Value}, nil
	}
	return nil, errors.New("Wrong type")
}
func evalDivOperator(left, right IObject) (IObject, error) {
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return nil, errors.New("Wrong type")
	}

	switch leftObject := left.(type) {
	case NumberObject:
		rightObject, _ := right.(NumberObject)
		return NumberObject{Value: leftObject.Value / rightObject.Value}, nil
	case FloatNumberObject:
		rightObject, _ := right.(FloatNumberObject)
		return FloatNumberObject{Value: leftObject.Value / rightObject.Value}, nil
	}
	return nil, errors.New("Wrong type")
}
func evalMultOperator(left, right IObject) (IObject, error) {
	if reflect.TypeOf(left) != reflect.TypeOf(right) {
		return nil, errors.New("Wrong type")
	}

	switch leftObject := left.(type) {
	case NumberObject:
		rightObject, _ := right.(NumberObject)
		return NumberObject{Value: leftObject.Value * rightObject.Value}, nil

	case FloatNumberObject:
		rightObject, _ := right.(FloatNumberObject)
		return FloatNumberObject{Value: leftObject.Value * rightObject.Value}, nil
	}
	return nil, errors.New("Wrong type")
}
