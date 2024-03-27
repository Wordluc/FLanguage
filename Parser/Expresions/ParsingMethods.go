package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Attraction.go"
	"errors"
)

type ParseExpresion func(a Token.Token, l *Lexer.Lexer) (*Expresion, error)

func GetParse(than Token.TokenType) (ParseExpresion, error) {
	switch than {
	case Token.LET:
		return ParseBinaryOp, nil
	}
	return nil, errors.New("GetParse: " + string(than) + "not implemented")
}

func ParseBinaryOp(a Token.Token, l *Lexer.Lexer) (*Expresion, error) { //sono sul operazione
	forceCurOp, e := Attraction.GetForce(l.LookCurrent().Type)
	if e != nil {
		return &Expresion{}, e
	}
	expresion := &Expresion{Type: a.Type, Value: a.Value, NextExpresion: nil}
	expresion.NextExpresion = &Expresion{Type: l.LookCurrent().Type, Value: l.LookCurrent().Value}
	nextValue, e := l.NextToken()
	if e != nil {
		return expresion, e
	}

	nextOp, e := l.NextToken()
	if e != nil {
		return expresion, e
	}
	forceNextOp, e := Attraction.GetForce(nextOp.Type)
	if e != nil {
		return expresion, e
	}
	for {
		if forceCurOp > forceNextOp {
			expresion.NextExpresion = &Expresion{Type: nextValue.Type, Value: nextValue.Value}
			break
		}
		f, e := GetParse(nextOp.Type)
		expresion.NextExpresion, e = f(nextValue, l)
		if e != nil {
			return expresion, e
		}

	}
	return expresion, nil
}
