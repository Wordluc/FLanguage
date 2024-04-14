package Evaluator

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements"
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

func TestCallFuncSetVariable(t *testing.T) {

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

func TestCallFuncReturn(t *testing.T) {

	ist := `
	Ff prova (){
	   let a = 5+3*(3*(4+2));		
	   ret a;
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
	program, e := Eval(root.(*Statements.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	if (program.(*ReturnObject).Value.(*NumberObject).Value) != 59 {
		t.Errorf("value is not %v got %v", 59, program.(*ReturnObject).Value.(*NumberObject).Value)
	}
}

func TestCallFuncWithoutReturn(t *testing.T) {

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
	program, e := Eval(root.(*Statements.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	if (program.(*ReturnObject).Value) != nil {
		t.Error("should not return anything")
	}
}

func TestBooleanOp(t *testing.T) {

	ist := `
	ret 3>2;
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
	program, e := Eval(root.(*Statements.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	val := (program.(*ReturnObject).Value)
	v, _ := val.(*BoolObject)
	if !v.Value {
		t.Error("3 is greater than 2")
	}

}

func TestNumberComparison(t *testing.T) {

	ist := `
	let a=3==3;
	let b=3==2;
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
	a, _ := env.GetVariable("a")
	b, _ := env.GetVariable("b")
	if !a.(*BoolObject).Value {
		t.Error("should be true")
	}
	if b.(*BoolObject).Value {
		t.Error("should be false")
	}

}

func TestStringComparison(t *testing.T) {

	ist := `
	let a="ffff"!="f";
	let b="ciao"=="ciao";
	let c=3>=3;
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
	a, _ := env.GetVariable("a")
	b, _ := env.GetVariable("b")
	c, _ := env.GetVariable("c")
	if !a.(*BoolObject).Value {
		t.Error("should be true, ffff != f")
	}
	if !b.(*BoolObject).Value {
		t.Error("should be true, ciao==ciao")
	}
	if !c.(*BoolObject).Value {
		t.Error("should be true, 3>=3")
	}
}

func TestMultipleDeclarationWtihSameName(t *testing.T) {

	ist := `
	let a="ffff"!="f";
	let a="ciao"=="ciao";
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
	if e == nil {
		t.Error("should have error")
	}
}
func TestAssigmentDifferentValue(t *testing.T) {

	ist := `
	let a="ffff"!="f";
	a="ciao";
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
	if e == nil {
		t.Error("should have error")
	}
}
