package Evaluator

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements"
	"reflect"
	"testing"
)

func TestLetAssigment(t *testing.T) {

	ist := `let a = 5+3*(3*(4+2));
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	programParse, e := Statements.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallita")
	}
	root := programParse
	program, e := Eval(root.(*Statements.StatementNode), &Environment{variables: make(map[string]IObject), internals: nil})
	if e != nil {
		t.Error("eval fallita", e)
	}
	object, isLet := program.(*LetObject)
	if !isLet {
		t.Error("is not a let")
	}
	v := object.Value.(*NumberObject).Value
	if v != 59 {
		t.Errorf("value is not %v got %v", 59, v)
	}
	t.Log(reflect.TypeOf(program).String())
}

func TestAssigment(t *testing.T) {

	ist := `
	let a = 5+3*(3*(4+2));
	a=5;
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	programParse, e := Statements.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallita")
	}
	root := programParse
	env := &Environment{variables: make(map[string]IObject), internals: nil}
	_, e = Eval(root.(*Statements.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	envObject, e := env.GetVariable("a")
	v := envObject.(*NumberObject).Value
	if v != 5 {
		t.Errorf("value is not %v got %v", 5, v)
	}
}

func TestAssigmentAndReuse(t *testing.T) {
	ist := `
	let a = 5+3*(3*(4+2));
	a=a*3;
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	programParse, e := Statements.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallita")
	}
	root := programParse
	env := &Environment{variables: make(map[string]IObject), internals: nil}
	_, e = Eval(root.(*Statements.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	envObject, e := env.GetVariable("a")
	v := envObject.(*NumberObject).Value
	if v != 177 {
		t.Errorf("value is not %v got %v", 177, v)
	}
}

func TestCallFunc(t *testing.T) {

	ist := `
	Ff prova (){
	   let a = 5+3*(3*(4+2));		
	}
	prova();
	END
	`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	programParse, e := Statements.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := &Environment{variables: make(map[string]IObject), functions: make(map[string]Statements.FuncDeclarationStatement), internals: nil}
	_, e = Eval(root.(*Statements.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	v, _ := env.internals.GetVariable("a")
	if v.(*NumberObject).Value != 59 {
		t.Errorf("value is not %v got %v", 177, v)
	}
}
