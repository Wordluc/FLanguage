package Statements

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"slices"
	"strings"
	"testing"
)

func IsEqual(a, b string) bool {
	return slices.Equal(strings.Fields(a), strings.Fields(b))
}
func TestParsingLetStatements(t *testing.T) {

	ist := "let a = 5;END"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		//t.Error(e)
	}
	program, e := ParsingStatement(&lexer, Token.DOT_COMMA)
	if e != nil {
		t.Error(e)
	}
	if program == nil {
		t.Error("program is nil")
	}

	expected := "LET a = 5"
	if !IsEqual(program.ToString(), expected) {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParsingLetStatements2(t *testing.T) {

	ist := "let a = 5+3*(3*(4+2));END"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		//t.Error(e)
	}
	program, e := ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error(e)
	}
	if program == nil {
		t.Error("program is nil")
	}

	expected := "LET a = 5 + (3 * (3 * (4 + 2)))"

	if !IsEqual(program.ToString(), expected) {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParsingLetStatementsWithoutSemicolon(t *testing.T) {

	ist := "let a = 5+3*(3*(4+2))END"
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		//t.Error(e)
	}
	program, e := ParsingStatement(&lexer, Token.END)
	if e == nil {
		t.Error("expected error")
	}
	if program != nil {
		t.Error("expected not program")
	}
}
func TestParsingWithMoreLetStatements(t *testing.T) {

	ist := `let a = 5+3*(3*(4+2));
		let b=2;
		END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		//t.Error(e)
	}
	program, e := ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error(e)
	}
	expected := `
	LET a = 5 + (3 * (3 * (4 + 2)))
	LET b = 2`
	if !IsEqual(program.ToString(), expected) {
		t.Error("error parsing", "expected:", expected, "\ngot: \n", program.ToString())
	}
}
func TestParsingLetWithCallFunc(t *testing.T) {

	ist := `let a = Pippo("ciao","frfr",3);
		END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		//t.Error(e)
	}
	program, e := ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error(e)
	}
	expected := `
	LET a = Pippo("ciao","frfr",3,)`
	if !IsEqual(program.ToString(), expected) {
		t.Error("error parsing", "expected:", expected, "\ngot: \n", program.ToString())
	}
}
func TestParsingLetWithoutEqual(t *testing.T) {

	ist := `let a Pippo("ciao","frfr",3);
		END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		//t.Error(e)
	}
	_, e = ParsingStatement(&lexer, Token.END)
	if e == nil {
		t.Error("expected error")
	}
}
func TestParsingLetWithoutEND(t *testing.T) {

	ist := `let a= Pippo("ciao","frfr",3);
		`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		//t.Error(e)
	}
	_, e = ParsingStatement(&lexer, Token.END)
	if e == nil {
		t.Error("expected error")
	}
}
