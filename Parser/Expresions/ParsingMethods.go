package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Attraction.go"
	"errors"
)

type ParseExpresion func(l *Lexer.Lexer) (IExpresion, error)

func ParseExpresionEmpty(l *Lexer.Lexer) (IExpresion, error) {
	return EmptyExpresion{}, nil
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
func ParseWord(l *Lexer.Lexer) (IExpresion, error) {
	nextOp, e := l.NextToken()
	if e != nil {
		return EmptyExpresion{}, e
	}
	f, e := GetParse(nextOp.Type)
	if e != nil {
		return EmptyExpresion{}, And(e, "ParseWord"+nextOp.Value)
	}
	exp, e := f(l)
	if e != nil {
		return EmptyExpresion{}, e
	}
	return exp, nil
}
func ParseBinaryOp(l *Lexer.Lexer) (IExpresion, error) { //sono sul operazione
	backValue, e := l.LookBack()
	if e != nil {
		return EmptyExpresion{}, And(e, "backValue")
	}
	forceCurOp, e := Attraction.GetForce(l.LookCurrent().Type)
	if e != nil {
		return EmptyExpresion{}, And(e, "forceCurOp"+l.LookCurrent().Value)
	}
	expresion := BinaryExpresion{TypeA: backValue.Type, ValueA: backValue.Value, NextExpresion: EmptyExpresion{}}
	expresion.Operator = l.LookCurrent()
	expresion.ValueOperator = l.LookCurrent().Value
	nextValue, e := l.NextToken()
	if e != nil {
		return expresion, And(e, "nextValue")
	}
	nextOp, e := l.LookNext()
	if e != nil {
		return expresion, And(e, "nextOp")
	}
	if nextOp.Type == Token.DOT_COMMA {
		expresion.NextExpresion = BinaryExpresion{TypeA: nextValue.Type, ValueA: nextValue.Value, NextExpresion: EmptyExpresion{}}
		return expresion, nil
	}
	forceNextOp, e := Attraction.GetForce(nextOp.Type)
	if e != nil {
		return expresion, And(e, "forceNextOp")
	}
	if forceCurOp > forceNextOp {
		l.IncrP()
		expresion.NextExpresion = BinaryExpresion{TypeA: nextValue.Type, ValueA: nextValue.Value, NextExpresion: EmptyExpresion{}}
	}
	return expresion, nil
}
