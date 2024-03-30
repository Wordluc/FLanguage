package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"testing"
)

func TestBinaryOperation_OneOperation(t *testing.T) {
	ist := "10+2;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		//t.Error(e)
	}
	lexer.IncrP()
	program, e := ParseBinaryOp(&lexer)
	if e != nil {
		t.Error(e)
	}
	if program == nil {
		t.Error("program is nil")
	}
	CompareWith := []CompareWith{
		{TypeA: Token.WORD, ValueA: "10", Operator: Token.PLUS, ValueOperator: "+"},
		{TypeA: Token.WORD, ValueA: "2"},
	}
	head := program.(BinaryExpresion)
	_ = head
	i := 0
	//t.SkipNow()
	for {
		e := head.Is(CompareWith[i])
		if e != nil {
			t.Error(e)
			break
		}
		switch head.NextExpresion.(type) {
		case BinaryExpresion:
			head = head.NextExpresion.(BinaryExpresion)
			i++
		case nil:
			return
		case EmptyExpresion:
			return
		default:
			return
		}

	}
}
func TestBinaryOperation_TwoOperation(t *testing.T) {
	ist := "10+2*1;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		//t.Error(e)
	}
	program, e := ParseWord(&lexer)
	if e != nil {
		t.Error(e)
	}
	expected := []CompareWith{
		{TypeA: Token.WORD, ValueA: "10", Operator: Token.PLUS, ValueOperator: "+"},
	}
	head := program.(BinaryExpresion)
	i := 0
GO:
	for {
		e := head.Is(expected[i])
		t.Log(head.ValueA)
		if e != nil {
			t.Error(e)
			break
		}

		switch head.NextExpresion.(type) {
		case BinaryExpresion:
			head = head.NextExpresion.(BinaryExpresion)
			i++
		case EmptyExpresion:
			break GO
		case nil:
			return
		default:
			t.Error("error parsing")
			return
		}

	}

	//	lexer.IncrP()
	program, e = ParseWord(&lexer)
	//print program
	head = program.(BinaryExpresion)
	if e != nil {
		t.Error(e)
	}

	expected = []CompareWith{
		{TypeA: Token.WORD, ValueA: "2", Operator: Token.MULT, ValueOperator: "*"},
		{TypeA: Token.WORD, ValueA: "1"},
	}
	head = program.(BinaryExpresion)
	i = 0
	for {
		e := head.Is(expected[i])
		if e != nil {
			t.Error(e)
			break
		}
		switch head.NextExpresion.(type) {
		case BinaryExpresion:
			head = head.NextExpresion.(BinaryExpresion)
			i++
		case EmptyExpresion:
			return
		case nil:
			t.Log("nil")
			return
		}

	}
}
