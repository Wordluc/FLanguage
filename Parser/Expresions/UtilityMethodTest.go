package Expresions

import (
	"FLanguage/Lexer/Token"
	"errors"
)

type CompareWith struct {
	Type  Token.TokenType
	Value string
}

func (e Expresion) Is(c CompareWith) error {
	if e.Type == c.Type && e.Value == c.Value {
		return nil
	}
	return errors.New(string(e.Type) + "is not " + string(c.Type) + "and " + string(e.Value) + "is not " + string(c.Value))
}
