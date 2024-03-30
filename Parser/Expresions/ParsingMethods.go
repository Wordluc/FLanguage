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
	return nil, errors.New("GetParse: Operator:" + string(than) + "not implemented")
}
func ParseProgram(l *Lexer.Lexer) (IExpresion, error) {
	f, e := GetParse(l.LookCurrent().Type)
	if e != nil {
		return EmptyExpresion{}, And(e, "ParseProgram")
	}
	exp, e := f(l)
	program := &MainExpresion{Expresion: exp}
	head := program
	for {
		//	head.NextExpresion = &MainExpresion{Expresion: &OperatorExpresion{Operator: l.LookCurrent()}}
		//	head = head.GetNextExpresion().(*MainExpresion)
		//	l.IncrP()
		f, e := GetParse(l.LookCurrent().Type)
		if e != nil {
			return EmptyExpresion{}, And(e, "ParseProgram")
		}
		ex, e := f(l)
		if e != nil {
			return program, e
		}
		head.NextExpresion = &MainExpresion{Expresion: ex}
		head = head.GetNextExpresion().(*MainExpresion)
		if l.LookCurrent().Type == Token.DOT_COMMA { //TODO: permettere di definire l'uscita dal for
			break
		}
	}
	return program, nil
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
	expresion := BinaryExpresion{}.New(backValue, l.LookCurrent())
	nextValue, e := l.NextToken()
	if e != nil {
		return expresion, And(e, "nextValue")
	}
	nextOp, e := l.LookNext()
	if e != nil {
		return expresion, And(e, "nextOp")
	}
	if nextOp.Type == Token.DOT_COMMA {
		l.IncrP()
		expresion.NextExpresion = BinaryExpresion{}.NewValue(nextValue)
		return expresion, nil
	}
	forceNextOp, e := Attraction.GetForce(nextOp.Type)
	if e != nil {
		return expresion, And(e, "forceNextOp")
	}
	l.IncrP()
	if forceCurOp > forceNextOp {
		expresion.NextExpresion = BinaryExpresion{}.NewValue(nextValue)
		return expresion, nil
	}

	f, e := GetParse(nextOp.Type)
	if e != nil {
		return expresion, And(e, "f")
	}
	exp, e := f(l)
	if e != nil {
		return expresion, And(e, "exp")
	}
	expresion.NextExpresion = exp
	return expresion, nil
}
