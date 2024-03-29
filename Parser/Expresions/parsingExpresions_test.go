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
	program, e := ParseWord(&lexer)
	if e != nil {
		t.Error(e)
	}
	if program == nil {
		t.Error("program is nil")
	}
	CompareWith := []CompareWith{
		{Type: Token.WORD, Value: "10"},
		{Type: Token.PLUS, Value: "+"},
		{Type: Token.WORD, Value: "2"},
	}
	head := program
	i := 0
	for {
		e := head.Is(CompareWith[i])
		if e != nil {
			t.Error(e)
		}
		head = head.NextExpresion
		if head == nil {
			break
		}
		i++

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
	t.Log("mi manca" + lexer.LookCurrent().Value)
	expected := []CompareWith{
		{Type: Token.WORD, Value: "10"},
		{Type: Token.PLUS, Value: "+"},
	}
	head := program
	for {

		if head == nil {
			break
		}
		t.Log("log", head.Value)
		head = head.NextExpresion
	}
	head = program
	i := 0
	for {
		e := head.Is(expected[i])
		if e != nil {
			t.Error(e)
		}
		head = head.NextExpresion
		if head == nil {
			break
		}
		i++

	}
	//	lexer.IncrP()
	program, e = ParseWord(&lexer)
	//print program
	head = program
	t.Log("print")
	if e != nil {
		t.Error(e)
	}
	for {

		if head == nil {
			break
		}
		t.Log(head.Value)
		head = head.NextExpresion
	}

	expected = []CompareWith{
		{Type: Token.WORD, Value: "2"},
		{Type: Token.MULT, Value: "*"},
		{Type: Token.WORD, Value: "1"},
	}
	head = program
	i = 0
	for {
		e := head.Is(expected[i])
		if e != nil {
			t.Error(e)
		}
		head = head.NextExpresion
		if head == nil {
			break
		}
		i++

	}
}
