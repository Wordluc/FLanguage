package Evaluator

import (
	"FLanguage/Lexer/Token"
	"errors"
	"strconv"
)

func evalBinaryExpresion(left, right IObject, operator Token.Token) (IObject, error) {
	switch leftObject := left.(type) {
	case NumberObject:
		valueLeft := leftObject.Value
		stringValue, isRightString := right.(StringObject)
		if isRightString {
			if operator.Type == Token.PLUS {
				return StringObject{strconv.Itoa(valueLeft) + stringValue.Value}, nil
			} else {
				return nil, errors.New("invalid operator")
			}
		}
		valueRight := right.(NumberObject).Value
		switch operator.Type {
		case Token.PLUS:
			return NumberObject{valueLeft + valueRight}, nil
		case Token.MINUS:
			return NumberObject{valueLeft - valueRight}, nil
		case Token.DIV:
			return NumberObject{valueLeft / valueRight}, nil
		case Token.MULT:
			return NumberObject{valueLeft * valueRight}, nil
		case Token.GREATER:
			return BoolObject{Value: valueLeft > valueRight}, nil
		case Token.LESS:
			return BoolObject{Value: valueLeft < valueRight}, nil
		case Token.EQUAL:
			return BoolObject{Value: valueLeft == valueRight}, nil
		case Token.NOT_EQUAL:
			return BoolObject{Value: valueLeft != valueRight}, nil
		case Token.GREATER_EQUAL:
			return BoolObject{Value: valueLeft >= valueRight}, nil
		case Token.LESS_EQUAL:
			return BoolObject{Value: valueLeft <= valueRight}, nil
		}
	case StringObject:
		valueLeft := leftObject.Value
		stringValue, isRightNumber := right.(NumberObject)
		if isRightNumber {
			if operator.Type == Token.PLUS {
				return StringObject{valueLeft + strconv.Itoa(stringValue.Value)}, nil
			} else {
				return nil, errors.New("invalid operation")
			}
		} //todo: sistemare string+bool
		valueRight := right.(StringObject).Value
		switch operator.Type {
		case Token.PLUS:
			return StringObject{valueLeft + valueRight}, nil
		case Token.EQUAL:
			return BoolObject{valueLeft == valueRight}, nil
		case Token.NOT_EQUAL:
			return BoolObject{valueLeft != valueRight}, nil
		}
	case BoolObject:
		valueLeft := leftObject.Value
		valueRight := right.(BoolObject).Value
		if operator.Type == Token.EQUAL {
			return BoolObject{
				Value: valueLeft == valueRight,
			}, nil
		}
	}
	return nil, errors.New("invalid operation")
}
