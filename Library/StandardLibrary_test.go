package Library

import (
	"FLanguage/Evaluator"
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	ist := `
	import("BinarySearch.txt");
	let list=[2,6,7,9,22,44,55,66,77,88,99];
	let b=BinarySearch(list,66);
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

	env := Evaluator.NewEnvironment()
	Evaluator.LoadBuiltInFunction(env)
	Evaluator.LoadBuiltInVariable(env)
	_, e = Evaluator.Eval(root.(*Statements.StatementNode), env)
	if e != nil {
		t.Error(e)
	}
	a, e := env.GetVariable("b")
	if e != nil {
		t.Error(e)
	}
	v := a.(Evaluator.NumberObject)
	if e != nil {
		t.Error(e)
	}
	if v.Value != 7 {
		t.Error("should be '7' ,got:", v.Value)
	}
}

func TestTree(t *testing.T) {
	ist := `
	import("Tree.txt");
	let list=[22,21,6,7,9,0,55,6,-20,88,99];
	let node=MakeTree(list);
	let orderedList=FromTreeToList(node,len(list));
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

	env := Evaluator.NewEnvironment()
	Evaluator.LoadBuiltInFunction(env)
	Evaluator.LoadBuiltInVariable(env)
	_, e = Evaluator.Eval(root.(*Statements.StatementNode), env)
	if e != nil {
		t.Error(e)
	}
	a, e := env.GetVariable("orderedList")
	if e != nil {
		t.Error(e)
	}
	v := a.(Evaluator.ArrayObject)
	if e != nil {
		t.Error(e)
	}
	expected := []int{-20, 0, 6, 6, 7, 9, 21, 22, 55, 88, 99}
	for i := 0; i < len(expected); i++ {
		if v.Values[i].(Evaluator.NumberObject).Value != expected[i] {
			t.Error("error order")
		}
	}
}
