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
func TestParseExpresion_WithoutCloseBrackets(t *testing.T) {
	ist := "prova(3+3,'ciao'+3*4;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	_, e = ParseExpresion(&lexer, Token.DOT_COMMA)
	if e == nil {
		t.Error("expected error")
		return
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
func TestParseExpresion_WithBracketAtTheEnd_ShouldPass(t *testing.T) {
	ist := "22*(2+16)*2;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	expected := "(22 * (2 + 16)) * 2"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_CallFunc_NoParms_ShouldPass(t *testing.T) {
	ist := "prova();"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	_, e = ParseExpresion(&lexer, Token.DOT_COMMA)
}
func TestParseExpresion_CallFunc_WithParm_ShouldPass(t *testing.T) {
	ist := "prova(3);"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	expected := "prova(3,)"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_CallFuncAtEnd_WithParm_ShouldPass(t *testing.T) {
	ist := "4*prova(3);"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
		return
	}
	expected := "4 * (prova(3,))"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_CallFunc_With3Parms_ShouldPass(t *testing.T) {
	ist := "prova(3,4,ciao);"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	expected := "prova(3,4,ciao,)"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_CallFunc_WithExpressionParms_ShouldPass(t *testing.T) {
	ist := "prova(3+3,4*2+(2+2),ciao);"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
		return
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
		return
	}
	expected := "prova(3 + 3,(4 * 2) + (2 + 2),ciao,)"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_CallFunc_WithExpressionParmsAndExpresion_ShouldPass(t *testing.T) {
	ist := "prova(3+3,ciao)+3*4;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	expected := "(prova(3 + 3,ciao,)) + (3 * 4)"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_StringParameter(t *testing.T) {
	ist := "prova(3+3,'ciao')+3*4;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
		return
	}
	expected := "(prova(3 + 3,'ciao',)) + (3 * 4)"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
