package Expresions

import (
	"FLanguage/Lexer/Token"
	"errors"
)

type CompareWith struct {
	TypeA         Token.TokenType
	ValueA        string
	Operator      Token.TokenType
	ValueOperator string
}

func (e BinaryExpresion) Is(c CompareWith) error {
	if !(e.TypeA == c.TypeA && e.ValueA == c.ValueA) {
		return errors.New("A:type or value are not equal, got " + e.ValueA + " expected " + c.ValueA)
	}
	if !(e.Operator.Type == c.Operator && e.Operator.Value == c.ValueOperator) {
		return errors.New("Operator:type or value are not equal, got " + e.ValueA + "expected " + c.ValueA)
	}
	return nil
}
