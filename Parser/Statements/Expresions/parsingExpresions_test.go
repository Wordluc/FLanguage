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
func TestParseExpresion_WithDecimalNumber(t *testing.T) {
	ist := "10.3;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	expected := "10.3"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_WithNegativeDecimalNumber(t *testing.T) {
	ist := "-10.3;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	expected := " - 10.3"
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
	expected := `(prova(3 + 3,"ciao",)) + (3 * 4)`
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}

func TestParseBooleanExpresion_ShouldPass(t *testing.T) {
	ist := "2>4;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	expected := "2 > 4"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}

func TestParseBooleanExpresionEqual_ShouldPass(t *testing.T) {
	ist := "2==4;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	expected := "2 == 4"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}

func TestParseBooleanValue_ShouldPass(t *testing.T) {
	ist := "true;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	expected := "true"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
	if program.(ExpresionLeaf).Type != Token.BOOLEAN {
		t.Error("error parsing expected:Bool, got", program.(ExpresionLeaf).Type)
	}
}

func TestParseBooleanComplexExpresion_ShouldPass(t *testing.T) {
	ist := "4*2>4+1;"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	expected := "(4 * 2) > (4 + 1)"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}

}

func TestParseStringComparison(t *testing.T) {
	ist := "'ciao'=='ciao';"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	expected := `"ciao" == "ciao"`
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}

func TestParseGetArrayOp(t *testing.T) {
	ist := `
	a[2];
	`
	//[1,2,3,4];/
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	expected := "a[2,]"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseDeclareArray(t *testing.T) {
	ist := `
	[1,2,3,4];
	`
	//[1,2,3,4];/
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	expected := "[1,2,3,4,]"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestNegativeNumber(t *testing.T) {
	ist := `
	-1;
	`
	//[1,2,3,4];/
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	expected := " - 1"
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestDeclareHash(t *testing.T) {
	ist := `
	{"cioa":3,"pep":4};
	`
	//[1,2,3,4];/
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	expected := `{"cioa":3,"pep":4,}`
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestGetHash(t *testing.T) {
	ist := `
	a{"cioa"};
	`
	//[1,2,3,4];/
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, e := ParseExpresion(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	expected := `a{"cioa"}`
	if program.ToString() != expected {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
