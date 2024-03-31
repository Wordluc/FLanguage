package Expresions

import (
	"FLanguage/Lexer"
	"testing"
)

func ParseExpresion_WithOneValue_ShouldPass(t *testing.T) {
	ist := "10;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer)
	if e != nil {
		t.Error(e)
	}
	expected := "10"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func ParseExpresion_Valid_ShouldPass1(t *testing.T) {
	ist := "10+2*1+2+22/2+16*3;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer)
	if e != nil {
		t.Error(e)
	}
	expected := "(((10 + (2 * 1)) + 2) + (22 / 2)) + (16 * 3)"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func ParseExpresion_Valid_ShouldPass2(t *testing.T) {
	ist := "10+22/2+16*3;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer)
	if e != nil {
		t.Error(e)
	}
	expected := "(10 + (22 / 2)) + (16 * 3)"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func ParseExpresion_Invalid_ShouldFail(t *testing.T) {
	ist := "10+22**2+16*3;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	_, e = ParseExpresion(&lexer)
	if e == nil {
		t.Error("should be error")
	}

}
