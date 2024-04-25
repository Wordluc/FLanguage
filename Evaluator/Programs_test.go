package Evaluator

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements"
	"testing"
)

func TestFunctionFactorial(t *testing.T) {
	ist := `
	//Calcola il fattoriale di un Number

	Ff factorial(x) {

		if (x == 0) {
			ret 1;
		}
		let c = x-1;
		ret x * factorial(x-1);
	}

	let a = factorial(5);
	//Token che indica la fine del programma
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
	if a.(NumberObject).Value != 120 {
		t.Error("value should be 120")
	}

}

func TestFunctionFibonacci(t *testing.T) {
	ist := `
	/*
	Dato un numero 
	calcola il valore della serie
	*/

	Ff fibonacci(x) {

		if (x == 0) {
			ret 0;
		}

		if (x == 1) {
			ret 1;
		}
		//Continua la ricorsione
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
	if a.(NumberObject).Value != 55 {
		t.Error("value should be 55")
	}
}

func TestElevation(t *testing.T) {
	ist := `	
	Ff eleva(x,i) {

		if (i == 0) {
			ret 1;
		}
		ret x * eleva(x,i-1);
	}
	let a = eleva(3,3);
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
	if a.(NumberObject).Value != 27 {
		t.Error("value should be 27")
	}
}

func TestGetStringWithMoreCharacters(t *testing.T) {
	ist := `	
	let a = ["1","11","22222244","345","oaaaa"];

	Ff search(i,iMax) {

		if (i==len(a)-1) {
			ret  a[iMax];
		}
		i=i+1;

		if (len(a[i])>len(a[iMax])) {
			ret search(i,i);
		}else{
			ret search(i,iMax);
		}
	}

	let b = search(0,0);
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
	env := NewEnvironment()
	LoadBuiltInVariable(env)
	LoadBuiltInFunction(env)
	_, e = Eval(programParse.(*Statements.StatementNode), env)
	if e != nil {
		t.Error(e)
	}

	b, _ := env.GetVariable("b")
	if b.(StringObject).Value != "22222244" {
		t.Error("value should be 8,got:", b.(StringObject).Value)
	}
}

func TestDicotomicSearch(t *testing.T) {
	ist := `
	
	Ff RicercaDicotomica(array,value){
		let low=0;
		let high=len(array)-1;
		let i=len(array)/2;
		while (low<high){
			if(array[i]==value){
				ret i;				
			}

			if(value>array[i]){
                             low=i+1;
			}else{
			     high=i;
			}
			i=(high+low)/2;
		}
		ret -1;
	}
	let array=[1,2,3,7,8,9,10,11,13,14,15,16,17,19,20];
	let result= RicercaDicotomica(array,7);
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

	env := NewEnvironment()
	LoadBuiltInFunction(env)
	LoadBuiltInVariable(env)
	_, e = Eval(root.(*Statements.StatementNode), env)
	if e != nil {
		t.Error(e)
	}

	a, _ := env.GetVariable("result")
	if a.(NumberObject).Value != 3 {
		t.Error("value should be 3,got:", a.(NumberObject).Value)
	}
}
