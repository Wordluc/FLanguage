package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"testing"
)

func TestParseExpresion_WithOneValue_ShouldPass(t *testing.T) {
	ist := "10;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	expected := "10"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_Valid_ShouldPass1(t *testing.T) {
	ist := "10+2*1+2+22/2+16*3;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	expected := "(((10 + (2 * 1)) + 2) + (22 / 2)) + (16 * 3)"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_Valid_ShouldPass2(t *testing.T) {
	ist := "10+22/2+16*3;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	expected := "(10 + (22 / 2)) + (16 * 3)"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_Invalid_ShouldFail(t *testing.T) {
	ist := "10+22**2+16*3;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	_, e = ParseExpresion(&lexer, Token.DOT_COMMA)
	if e == nil {
		t.Error("should be error")
	}

}
func TestParseExpresion_WithoutSemicolon_ShouldFail(t *testing.T) {
	ist := "22*2+16*3"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	_, e = ParseExpresion(&lexer, Token.DOT_COMMA)
	if e == nil {
		t.Error("should be error")
	}

}
func TestParseExpresion_WithBracket_ShouldPass(t *testing.T) {
	ist := "22*(2+16)*3;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	expected := "(22 * (2 + 16)) * 3"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_WithBracket_ShouldPass_2(t *testing.T) {
	ist := "22*(2+16*2)*3;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	expected := "(22 * (2 + (16 * 2))) * 3"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
