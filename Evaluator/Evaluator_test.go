package Evaluator

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements"
	"errors"
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
	program, e := Eval(root.(*Statements.StatementNode), &Environment{variables: make(map[string]IObject), externals: nil})
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
	env := &Environment{variables: make(map[string]IObject), externals: nil}
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
	env := &Environment{variables: make(map[string]IObject), externals: nil}
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

	env := NewEnvironment()
	program, e := Eval(root.(*Statements.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	if (program.(*NumberObject).Value) != 59 {
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

	env := NewEnvironment()
	program, e := Eval(root.(*Statements.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	if (program) != nil {
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

	env := NewEnvironment()
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

	env := NewEnvironment()
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

	env := NewEnvironment()
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

	env := NewEnvironment()
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

	env := NewEnvironment()
	_, e = Eval(root.(*Statements.StatementNode), env)
	if e == nil {
		t.Error("should have error")
	}
}

func TestUseResultFunction(t *testing.T) {
	ist := `

	Ff add(){
		ret 3;
	}
	let a=2+add();
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
	_, e = Eval(root.(*Statements.StatementNode), env)
	v, _ := env.GetVariable("a")
	if v.(*NumberObject).Value != 5 {
		t.Error("should be 5,got:", v)
	}
}

func TestSumStrings(t *testing.T) {
	ist := `
	let a="ciao "+"bene";
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
	_, e = Eval(root.(*Statements.StatementNode), env)
	v, _ := env.GetVariable("a")
	if v.(*StringObject).Value != "ciao bene" {
		t.Error("should be ciao bene ,got:", v.(*StringObject).Value)
	}
}

func TestSumStringsFromFunc(t *testing.T) {
	ist := `

	Ff getString(){
		ret "prova";
	}
	let a="ciao "+getString();
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
	_, e = Eval(root.(*Statements.StatementNode), env)
	v, _ := env.GetVariable("a")
	if v.(*StringObject).Value != "ciao prova" {
		t.Error("should be ciao bene ,got:", v.(*StringObject).Value)
	}
}

func TestPassValueThroughtFunc(t *testing.T) {
	ist := `

	Ff getValue(n){
		ret n+1;
	}
	let a=getValue(3);
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
	_, e = Eval(root.(*Statements.StatementNode), env)
	v, _ := env.GetVariable("a")
	if v.(*NumberObject).Value != 4 {
		t.Error("should be 4 ,got:", v.(*NumberObject).Value)
	}
}

func TestIfStatement(t *testing.T) {
	ist := `

	let a=2;

	if (4>2){
	   a=a+2;
	}
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
	_, e = Eval(root.(*Statements.StatementNode), env)
	v, _ := env.GetVariable("a")
	if v.(*NumberObject).Value != 4 {
		t.Error("should be 4 ,got:", v.(*NumberObject).Value)
	}
}

func TestElseStatement(t *testing.T) {
	ist := `

	let a=2;

	if (4<2){
	   a=a+2;
	}else{
	   a=a*4;	
	}
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
	_, e = Eval(root.(*Statements.StatementNode), env)
	v, _ := env.GetVariable("a")
	if v.(*NumberObject).Value != 8 {
		t.Error("should be 8 ,got:", v.(*NumberObject).Value)
	}
}

func TestCombineStringAndNumber(t *testing.T) {
	ist := `

	let a="ciao per";
	a=a+2;
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
	_, e = Eval(root.(*Statements.StatementNode), env)
	if e != nil {
		t.Error(e)
	}
	a, _ := env.GetVariable("a")
	if a.(*StringObject).Value != "ciao per2" {
		t.Error("should be 'ciao per2' ,got:", a.(*StringObject).Value)
	}
}

func TestCombineNumberAndString(t *testing.T) {
	ist := `

	let a=2+"ciao per";
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
	_, e = Eval(root.(*Statements.StatementNode), env)
	if e != nil {
		t.Error(e)
	}
	a, _ := env.GetVariable("a")
	if a.(*StringObject).Value != "2ciao per" {
		t.Error("should be '2ciao per' ,got:", a.(*StringObject).Value)
	}
}
func TestDeclareAndGetFromArray(t *testing.T) {
	ist := `
	let a=[1,2,3,"cioa"];
	let b=a[2];
	let c=a[3];
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
	_, e = Eval(root.(*Statements.StatementNode), env)
	if e != nil {
		t.Error(e)
	}
	b, _ := env.GetVariable("b")
	if b.(*NumberObject).Value != 3 {
		t.Error("should be '3' ,got:", b.(*NumberObject).Value)
	}
	c, _ := env.GetVariable("c")
	if c.(*StringObject).Value != "cioa" {
		t.Error("should be 'cioa' ,got:", c.(*StringObject).Value)
	}
}

func provaLen(env *Environment) (IObject, error) {
	aObject, e := env.GetVariable("a")
	if e != nil {
		return nil, e
	}
	switch a := aObject.(type) {
	case ArrayObject:
		return &NumberObject{Value: (len(a.Values))}, nil
	case *StringObject:
		return &NumberObject{Value: (len(a.Value))}, nil
	default:
		return nil, errors.New("not an array or string")
	}
}
func TestInnerFunc(t *testing.T) {
	ist := `
	let a=[1,2,3,"cioa"];
	let b=len(a);
	let c=len("prova");
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
	env.SetInnerFunc("len", &InnerFuncObject{NameParams: []string{"a"}, innerfunc: provaLen})
	_, e = Eval(root.(*Statements.StatementNode), env)
	if e != nil {
		t.Error(e)
	}
	b, _ := env.GetVariable("b")
	if b.(*NumberObject).Value != 4 {
		t.Error("should be '4' ,got:", b.(*NumberObject).Value)
	}
	c, _ := env.GetVariable("c")
	if c.(*NumberObject).Value != 5 {
		t.Error("should be '5' ,got:", c.(*NumberObject).Value)
	}
}
