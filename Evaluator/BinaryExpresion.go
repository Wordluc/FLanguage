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

		switch valueRight := right.(type) {
		case NumberObject:
			switch operator.Type {
			case Token.PLUS, Token.MULT, Token.DIV, Token.MINUS:
				return mathOperatorInt(valueLeft, valueRight.Value, operator)
			case Token.GREATER, Token.LESS, Token.EQUAL, Token.NOT_EQUAL, Token.GREATER_EQUAL, Token.LESS_EQUAL:
				return boolOperatorInt(valueLeft, valueRight.Value, operator)
			}
		case FloatNumberObject:
			switch operator.Type {
			case Token.PLUS, Token.MULT, Token.DIV, Token.MINUS:
				return mathOperatorFloat(float64(valueLeft), float64(valueRight.Value), operator)
			case Token.GREATER, Token.LESS, Token.EQUAL, Token.NOT_EQUAL, Token.GREATER_EQUAL, Token.LESS_EQUAL:
				return boolOperatorFloat(float64(valueLeft), float64(valueRight.Value), operator)
			}
		}
	case FloatNumberObject:
		valueLeft := leftObject.Value
		stringValue, isRightString := right.(StringObject)
		if isRightString {
			if operator.Type == Token.PLUS {
				return StringObject{strconv.FormatFloat(valueLeft, 'f', -1, 32) + stringValue.Value}, nil
			} else {
				return nil, errors.New("invalid operator")
			}
		}
		var valueRight float64
		switch rightObject := right.(type) {
		case NumberObject:
			valueRight = float64(rightObject.Value)
		case FloatNumberObject:
			valueRight = rightObject.Value
		}
		switch operator.Type {
		case Token.PLUS, Token.MINUS, Token.DIV, Token.MULT:
			return mathOperatorFloat(valueLeft, valueRight, operator)
		case Token.GREATER, Token.LESS, Token.EQUAL, Token.NOT_EQUAL, Token.GREATER_EQUAL, Token.LESS_EQUAL:
			return boolOperatorFloat(valueLeft, valueRight, operator)
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

func mathOperatorFloat(valueLeft float64, valueRight float64, operator Token.Token) (IObject, error) {
	switch operator.Type {
	case Token.PLUS:
		return FloatNumberObject{valueLeft + valueRight}, nil
	case Token.MINUS:
		return FloatNumberObject{valueLeft - valueRight}, nil
	case Token.DIV:
		return FloatNumberObject{valueLeft / valueRight}, nil
	case Token.MULT:
		return FloatNumberObject{valueLeft * valueRight}, nil

	}
	return nil, errors.New("invalid operation")
}
func mathOperatorInt(valueLeft, valueRight int, operator Token.Token) (IObject, error) {
	switch operator.Type {
	case Token.PLUS:
		return NumberObject{valueLeft + valueRight}, nil
	case Token.MINUS:
		return NumberObject{valueLeft - valueRight}, nil
	case Token.DIV:
		return NumberObject{valueLeft / valueRight}, nil
	case Token.MULT:
		return NumberObject{valueLeft * valueRight}, nil

	}
	return nil, errors.New("invalid operation")
}
func boolOperatorInt(valueLeft, valueRight int, operator Token.Token) (IObject, error) {
	switch operator.Type {
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
	return nil, errors.New("invalid operation")
}
func boolOperatorFloat(valueLeft, valueRight float64, operator Token.Token) (IObject, error) {
	switch operator.Type {
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
	return nil, errors.New("invalid operation")
}
