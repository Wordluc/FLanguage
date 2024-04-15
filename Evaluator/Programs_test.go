package Evaluator

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements"
	"testing"
)

func TestFunctionFactorial(t *testing.T) {
	ist := `

	Ff factorial(x) {

		if (x == 0) {
			ret 1;
		}
		let c = x-1;
		ret x * factorial(x-1);
	}

	let a = factorial(5);
	END
	`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}

	programParse, e := Statements.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error(e)
	}
	env := &Environment{
		variables: make(map[string]IObject),
		functions: make(map[string]Statements.FuncDeclarationStatement),
	}
	_, e = Eval(programParse.(*Statements.StatementNode), env)
	if e != nil {
		t.Error(e)
	}

	a, _ := env.GetVariable("a")
	if a.(*NumberObject).Value != 120 {
		t.Error("value should be 120")
	}

}

func TestFunctionFibonacci(t *testing.T) {
	ist := `

	Ff fibonacci(x) {

		if (x == 0) {
			ret 0;
		}
		if (x == 1) {
			ret 1;
		}

		ret fibonacci(x-1) + fibonacci(x-2);
	}

	let a = fibonacci(10);
	END
	`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error(e)
	}

	programParse, e := Statements.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error(e)
	}
	env := &Environment{
		variables: make(map[string]IObject),
		functions: make(map[string]Statements.FuncDeclarationStatement),
	}
	_, e = Eval(programParse.(*Statements.StatementNode), env)
	if e != nil {
		t.Error(e)
	}

	a, _ := env.GetVariable("a")
	if a.(*NumberObject).Value != 55 {
		t.Error("value should be 55")
	}
}
