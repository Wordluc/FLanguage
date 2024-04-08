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
	program, _, e := ParsingStatement(&lexer, Token.END)
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
	program, _, e := ParsingStatement(&lexer, Token.END)
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
	program, _, e := ParsingStatement(&lexer, Token.END)
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
	program, _, e := ParsingStatement(&lexer, Token.END)
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
	program, _, e := ParsingStatement(&lexer, Token.END)
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
	_, _, e = ParsingStatement(&lexer, Token.END)
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
	_, _, e = ParsingStatement(&lexer, Token.END)
	if e == nil {
		t.Error("expected error")
	}
}
func TestParseExpresion_IF(t *testing.T) {
	ist := `
	if (x > 0) {
		let x = x + 1;	
	}
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, _, e := ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error(e)
		return
	}
	expected := `
	IF ( x > 0 ) {
		LET x = x + 1		
	}`
	if !IsEqual(program.ToString(), expected) {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_IFANDELSE(t *testing.T) {
	ist := `
	if (x > 0) {
		let x = x + 1;	
	}else{
		let a=prova();
	}
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, _, e := ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error(e)
		return
	}
	expected := `
	IF ( x > 0 ) {
		LET x = x + 1		
	} ELSE {
		LET a = prova()
	}`
	if !IsEqual(program.ToString(), expected) {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_WITHWORD(t *testing.T) {
	ist := `
	if (x > 0) {
		x = x + 1;	
	}else{
		a=prova();
	}
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, _, e := ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error(e)
		return
	}
	expected := `
	IF ( x > 0 ) {
		x = x + 1		
	} ELSE {
		a = prova()
	}`
	if !IsEqual(program.ToString(), expected) {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_CallFunc(t *testing.T) {
	ist := `
	Prova("cioa","frfr",3);
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, _, e := ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error(e)
		return
	}
	expected := `
	Prova("cioa","frfr",3,)`
	if !IsEqual(program.ToString(), expected) {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_FuncDefinition(t *testing.T) {
	ist := `
	Ff prova (a){
		Prova("cioa","frfr",3);
	}
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	_, funcs, e := ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error(e)
		return
	}
	expected := `
	Ff prova ( a ) {
	        Prova("cioa","frfr",3,)
	
	}`
	if !IsEqual(funcs["prova"].ToString(), expected) {
		t.Error("error parsing", "expected: ", expected, "got: ", funcs["prova"].ToString())
	}
}
func TestParseExpresion_TwoFuncDefinition(t *testing.T) {
	ist := `
	Ff prova (a){
		Prova("cioa","frfr",3);
	}
	Ff pippo (a,b,c){
		let a=3+4;
	}
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	_, funcs, e := ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error(e)
		return
	}
	expected := `
	Ff prova ( a ) {
	        Prova("cioa","frfr",3,)
	
	}`
	if !IsEqual(funcs["prova"].ToString(), expected) {
		t.Error("error parsing", "expected: ", expected, "got: ", funcs["prova"].ToString())
	}
	expected = `
	Ff pippo ( a, b, c ) {
                LET a = 3 + 4

        }`
	if !IsEqual(funcs["pippo"].ToString(), expected) {
		t.Error("error parsing", "expected: ", expected, "got: ", funcs["pippo"].ToString())
	}
}
func TestParseExpresion_FuncDefinitionWithCall(t *testing.T) {
	ist := `
	Ff prova (a){
		Prova("cioa","frfr",3);
	}
	prova("dio");
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, funcs, e := ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error(e)
		return
	}
	expectedFunc := `
	Ff prova ( a ) {
	        Prova("cioa","frfr",3,)
	
	}`
	expected := `
	prova("dio",) 
	`
	if !IsEqual(funcs["prova"].ToString(), expectedFunc) {
		t.Error("error parsing", "expected: ", expectedFunc, "got: ", funcs["prova"].ToString())
	}
	if !IsEqual(program.ToString(), expected) {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_All(t *testing.T) {
	ist := `
	Ff prova (a){
		Prova("cioa","frfr",3);
	}
	prova("dio");
	if (x > 0) {
		x = x + 1;	
	}else{
		a=prova(2*(3+4));
	}
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	program, funcs, e := ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error(e)
		return
	}
	expectedFunc := `
	Ff prova ( a ) {
	        Prova("cioa","frfr",3,)
	
	}`
	expected := `
	prova("dio",)
        IF ( x > 0 ) {
                x = x + 1

        } ELSE {
                a = prova(2 * (3 + 4),)
        }	`
	if !IsEqual(funcs["prova"].ToString(), expectedFunc) {
		t.Error("error parsing", "expected: ", expectedFunc, "got: ", funcs["prova"].ToString())
	}
	if !IsEqual(program.ToString(), expected) {
		t.Error("error parsing", "expected: ", expected, "got: ", program.ToString())
	}
}
func TestParseExpresion_ErrorDefinitionShouldFail(t *testing.T) {
	ist := `
	Ff (a){
		Prova("cioa","frfr",3);
	}
	prova("dio");
	if (x > 0) {
		x = x + 1;	
	}else{
		a=prova();
	}
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	_, _, e = ParsingStatement(&lexer, Token.END)
	if e == nil {
		t.Error("expected error parsing")
		return
	}
}
func TestParseExpresion_ErrorDefinition2ShouldFail(t *testing.T) {
	ist := `
	Ff prova a){
		Prova("cioa","frfr",3);
	}
	prova("dio");
	if (x > 0) {
		x = x + 1;	
	}else{
		a=prova();
	}
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}
	_, _, e = ParsingStatement(&lexer, Token.END)
	if e == nil {
		t.Error("expected error parsing")
		return
	}
}
