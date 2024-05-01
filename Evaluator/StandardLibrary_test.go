package Evaluator

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	ist := `
	import("BinarySearch.txt");
	let list=[2,6,7,9,22,44,55,66,77,88,99];
	let b=BinarySearch_Run(list,66);
	END
	`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	LoadBuiltInFunction(env)
	LoadBuiltInVariable(env)
	_, e = Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error(e)
	}
	a, e := env.getVariable("b")
	if e != nil {
		t.Error(e)
	}
	v := a.(numberObject)
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
	let node=Tree_MakeTree(list);
	let orderedList=Tree_FromTreeToList(node,len(list));
	END
	`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	LoadBuiltInFunction(env)
	LoadBuiltInVariable(env)
	_, e = Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error(e)
	}
	a, e := env.getVariable("orderedList")
	if e != nil {
		t.Error(e)
	}
	v := a.(arrayObject)
	if e != nil {
		t.Error(e)
	}
	expected := []int{-20, 0, 6, 6, 7, 9, 21, 22, 55, 88, 99}
	for i := 0; i < len(expected); i++ {
		if v.Values[i].(numberObject).Value != expected[i] {
			t.Error("error order")
		}
	}
}
