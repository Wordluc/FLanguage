package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Attraction.go"
	"errors"
)

type ParseExpresion func(l *Lexer.Lexer) (*Expresion, error)

func ParseExpresionEmpty(l *Lexer.Lexer) (*Expresion, error) {
	return &Expresion{}, nil
}
func And(e error, s string) error {
	v := e.Error()
	return errors.New(v + " " + s)
}
func GetParse(than Token.TokenType) (ParseExpresion, error) {
	switch than {
	case Token.DIV:
		return ParseBinaryOp, nil
	case Token.MULT:
		return ParseBinaryOp, nil
	case Token.MINUS:
		return ParseBinaryOp, nil
	case Token.PLUS:
		return ParseBinaryOp, nil
	case Token.WORD:
		return ParseWord, nil
	case Token.DOT_COMMA:
		return ParseExpresionEmpty, nil
	}
	return nil, errors.New("GetParse: " + string(than) + "not implemented")
}
func ParseWord(l *Lexer.Lexer) (*Expresion, error) {
	nextOp, e := l.NextToken()
	if e != nil {
		return nil, e
	}
	f, e := GetParse(nextOp.Type)
	if e != nil {
		return nil, And(e, "ParseWord"+nextOp.Value)
	}
	exp, e := f(l)
	if e != nil {
		return nil, e
	}
	return exp, nil
}
func ParseBinaryOp(l *Lexer.Lexer) (*Expresion, error) { //sono sul operazione
	backValue, e := l.LookBack()
	if e != nil {
		return &Expresion{}, And(e, "backValue")
	}
	forceCurOp, e := Attraction.GetForce(l.LookCurrent().Type)
	if e != nil {
		return &Expresion{}, And(e, "forceCurOp"+l.LookCurrent().Value)
	}
	expresion := &Expresion{Type: backValue.Type, Value: backValue.Value, NextExpresion: nil}
	expresion.NextExpresion = &Expresion{Type: l.LookCurrent().Type, Value: l.LookCurrent().Value}
	nextValue, e := l.NextToken()
	if e != nil {
		return expresion, And(e, "nextValue")
	}
	nextOp, e := l.LookNext()
	if e != nil {
		return expresion, And(e, "nextOp")
	}

	forceNextOp, e := Attraction.GetForce(nextOp.Type)
	if e != nil {
		return expresion, And(e, "forceNextOp")
	}
	if forceCurOp > forceNextOp {
		l.IncrP()
		expresion.NextExpresion.NextExpresion = &Expresion{Type: nextValue.Type, Value: nextValue.Value}
	}
	return expresion, nil
}
