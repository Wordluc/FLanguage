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
	root := programParse["root"]
	program, e := Eval(root.(*Statements.StatementNode), &VariableEnvironment{variables: make(map[string]IObject), externals: nil})
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
	root := programParse["root"]
	env := &VariableEnvironment{variables: make(map[string]IObject), externals: nil}
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
	root := programParse["root"]
	env := &VariableEnvironment{variables: make(map[string]IObject), externals: nil}
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
