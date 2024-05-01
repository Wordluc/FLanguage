package Evaluator

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser"
	"testing"
)

func TestLetAssigment(t *testing.T) {

	ist := `let a = 5+3*(3*(4+2));
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallita")
	}
	root := programParse
	program, e := Eval(root.(*Parser.StatementNode), &Environment{variables: make(map[string]iObject), externals: nil})
	if e != nil {
		t.Error("eval fallita", e)
	}
	object, isLet := program.(letObject)
	if !isLet {
		t.Error("is not a let")
	}
	v := object.Value.(numberObject).Value
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallita")
	}
	root := programParse
	env := &Environment{variables: make(map[string]iObject), externals: nil}
	_, e = Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	envObject, e := env.getVariable("a")
	v := envObject.(numberObject).Value
	if v != 5 {
		t.Errorf("value is not %v got %v", 5, v)
	}
}
func TestAssigmentNegativeNumber(t *testing.T) {

	ist := `
	let a=-5;
	let b=-(-5);
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallita")
	}
	root := programParse
	env := &Environment{variables: make(map[string]iObject), externals: nil}
	_, e = Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	envObject, e := env.getVariable("a")
	v := envObject.(numberObject).Value
	if v != -5 {
		t.Errorf("value is not %v got %v", -5, v)
	}
	envObject, e = env.getVariable("b")
	v = envObject.(numberObject).Value
	if v != 5 {
		t.Errorf("value is not %v got %v", 5, v)
	}
}
func TestAssigmentDecimaleNumber(t *testing.T) {

	ist := `
	let a=5.5;
	let b=a+float(1);
	END`
	lexer, e := Lexer.New([]byte(ist))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallita")
	}
	root := programParse

	env := NewEnvironment()
	LoadBuiltInFunction(env)
	_, e = Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	envObject, e := env.getVariable("a")
	v := envObject.(floatNumberObject).Value
	if v != 5.5 {
		t.Errorf("value is not %v got %v", 5.5, v)
	}
	envObject, e = env.getVariable("b")
	v = envObject.(floatNumberObject).Value
	if v != 6.5 {
		t.Errorf("value is not %v got %v", 6.5, v)
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallita")
	}
	root := programParse
	env := &Environment{variables: make(map[string]iObject), externals: nil}
	_, e = Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	envObject, e := env.getVariable("a")
	v := envObject.(numberObject).Value
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	program, e := Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	if (program.(numberObject).Value) != 59 {
		t.Errorf("value is not %v got %v", 59, program.(returnObject).Value.(numberObject).Value)
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	program, e := Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	if (program) != nil {
		t.Error("should not return anything")
	}
}
func TestIncVar(t *testing.T) {

	ist := `
	let a=2;
	a=a+1;
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
	_, _ = Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	v, _ := env.getVariable("a")
	if v.(numberObject).Value != 3 {
		t.Error("value should be 3")
	}

}
func TestIncInParm(t *testing.T) {

	ist := `

	Ff add(x){
		ret x+1;
	}
	let a=add(1+2);
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
	_, _ = Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	v, _ := env.getVariable("a")
	if v.(numberObject).Value != 4 {
		t.Error("value should be 4")
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	program, e := Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	val := (program.(returnObject).Value)
	v, _ := val.(boolObject)
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	_, e = Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	a, _ := env.getVariable("a")
	b, _ := env.getVariable("b")
	if !a.(boolObject).Value {
		t.Error("should be true")
	}
	if b.(boolObject).Value {
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	_, e = Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error("eval fallita", e)
	}
	a, _ := env.getVariable("a")
	b, _ := env.getVariable("b")
	c, _ := env.getVariable("c")
	if !a.(boolObject).Value {
		t.Error("should be true, ffff != f")
	}
	if !b.(boolObject).Value {
		t.Error("should be true, ciao==ciao")
	}
	if !c.(boolObject).Value {
		t.Error("should be true, 3>=3")
	}
}

func TestMultipleVariableDeclarationWtihSameName(t *testing.T) {

	ist := `
	let a="ffff"!="f";
	let a="ciao"=="ciao";
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
	_, e = Eval(root.(*Parser.StatementNode), env)
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	_, e = Eval(root.(*Parser.StatementNode), env)
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	_, e = Eval(root.(*Parser.StatementNode), env)
	v, _ := env.getVariable("a")
	if v.(numberObject).Value != 5 {
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	_, e = Eval(root.(*Parser.StatementNode), env)
	v, _ := env.getVariable("a")
	if v.(stringObject).Value != "ciao bene" {
		t.Error("should be ciao bene ,got:", v.(stringObject).Value)
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	_, e = Eval(root.(*Parser.StatementNode), env)
	v, _ := env.getVariable("a")
	if v.(stringObject).Value != "ciao prova" {
		t.Error("should be ciao bene ,got:", v.(stringObject).Value)
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	_, e = Eval(root.(*Parser.StatementNode), env)
	v, _ := env.getVariable("a")
	if v.(numberObject).Value != 4 {
		t.Error("should be 4 ,got:", v.(numberObject).Value)
	}
}
func TestSumStringInFunc(t *testing.T) {
	ist := `

	Ff getString(a,b){
		ret a+b;
	}
	let a=getString("ciao"," prova");
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
	_, e = Eval(root.(*Parser.StatementNode), env)
	v, _ := env.getVariable("a")
	if v.(stringObject).Value != "ciao prova" {
		t.Error("should be ciao prova ,got:", v.(stringObject).Value)
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	_, e = Eval(root.(*Parser.StatementNode), env)
	v, _ := env.getVariable("a")
	if v.(numberObject).Value != 4 {
		t.Error("should be 4 ,got:", v.(numberObject).Value)
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	_, e = Eval(root.(*Parser.StatementNode), env)
	v, _ := env.getVariable("a")
	if v.(numberObject).Value != 8 {
		t.Error("should be 8 ,got:", v.(numberObject).Value)
	}
}

func TestCombineStringAndNumber(t *testing.T) {
	ist := `

	let a="ciao per";
	a=a+string(2);
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
	_, e = Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error(e)
	}
	a, _ := env.getVariable("a")
	if a.(stringObject).Value != "ciao per2" {
		t.Error("should be 'ciao per2' ,got:", a.(stringObject).Value)
	}
}

func TestCombineNumberAndString(t *testing.T) {
	ist := `

	let a=string(2)+"ciao per";
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
	_, e = Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error(e)
	}
	a, _ := env.getVariable("a")
	if a.(stringObject).Value != "2ciao per" {
		t.Error("should be '2ciao per' ,got:", a.(stringObject).Value)
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	LoadBuiltInFunction(env)
	_, e = Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error(e)
	}
	b, _ := env.getVariable("b")
	if b.(numberObject).Value != 3 {
		t.Error("should be '3' ,got:", b.(numberObject).Value)
	}
	c, _ := env.getVariable("c")
	if c.(stringObject).Value != "cioa" {
		t.Error("should be 'cioa' ,got:", c.(stringObject).Value)
	}
}
func TestDeclareArrayIntoArray(t *testing.T) {
	ist := `
	let a=[1,[2,3]];
	let b=a[0];
	let c=a[1,0];
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
	_, e = Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error(e)
	}
	b, _ := env.getVariable("b")
	if b.(numberObject).Value != 1 {
		t.Error("should be '1' ,got:", b.(numberObject).Value)
	}
	c, _ := env.getVariable("c")
	if c.(numberObject).Value != 2 {
		t.Error("should be '2' ,got:", c.(numberObject).Value)
	}
}

func TestBuiltInFunc(t *testing.T) {
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
	programParse, e := Parser.ParsingStatement(&lexer, Token.END)
	if e != nil {
		t.Error("parsing fallito", e)
	}
	root := programParse

	env := NewEnvironment()
	LoadBuiltInFunction(env)
	_, e = Eval(root.(*Parser.StatementNode), env)
	if e != nil {
		t.Error(e)
	}
	b, _ := env.getVariable("b")
	if b.(numberObject).Value != 4 {
		t.Error("should be '4' ,got:", b.(numberObject).Value)
	}
	c, _ := env.getVariable("c")
	if c.(numberObject).Value != 5 {
		t.Error("should be '5' ,got:", c.(numberObject).Value)
	}
}
func TestOutofRangeArray(t *testing.T) {
	ist := `
	let a=[1,2,3,"cioa"];
	let c=a[5];
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
	_, e = Eval(root.(*Parser.StatementNode), env)
	if e == nil {
		t.Error("should be an error")
	}
}
func TestCreateArray(t *testing.T) {
	ist := `
	let a=newArray(4,0);
	let b=len(a);
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
	b, _ := env.getVariable("b")
	a, _ := env.getVariable("a")
	if b.(numberObject).Value != 4 {
		t.Error("should be '4' ,got:", b.(numberObject).Value)
	}
	if _, ok := a.(arrayObject); !ok {
		t.Error("should be a array ")
	}
	if _, ok := a.(arrayObject).Values[0].(numberObject); !ok {
		t.Error("should be a number")
	}
}
func TestGetMatrix(t *testing.T) {
	ist := `
	let a=[[2,4],[2,3,4]];
	let b=a[0][1];
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
	if v.Value != 4 {
		t.Error("should be '4' ,got:", v.Value)
	}
}
func TestGetArrayFromFunc(t *testing.T) {
	ist := `
	Ff getMatrix(){
		ret [[2,4],[2,3,4]];
	}
	let b=getMatrix()[0][1];
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
	if v.Value != 4 {
		t.Error("should be '4' ,got:", v.Value)
	}
}
func TestSetArray(t *testing.T) {
	ist := `
	let a=[1,2,3,5,6];
	a[2]=0;
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
	a, e := env.getVariable("a")
	if e != nil {
		t.Error(e)
	}
	v := a.(arrayObject)
	if e != nil {
		t.Error(e)
	}
	if v.Values[2].(numberObject).Value != 0 {
		t.Error("should be '0' ,got:", v.Values[2].(numberObject).Value)
	}

}
func TestWhile(t *testing.T) {
	ist := `
	let i=0;

	while (i<5){
		i=i+1;
	}
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
	a, e := env.getVariable("i")
	if e != nil {
		t.Error(e)
	}
	v := a.(numberObject)
	if e != nil {
		t.Error(e)
	}
	if v.Value != 5 {
		t.Error("should be '5' ,got:", v.Value)
	}

}

func TestGetCharFromSting(t *testing.T) {
	ist := `
	let a="afrfrfr";
	let b=a[1];
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
	v := a.(stringObject)
	if e != nil {
		t.Error(e)
	}
	if v.Value != "f" {
		t.Error("should be 'f' ,got:", v.Value)
	}
}
func TestGetHashValue(t *testing.T) {
	ist := `
	let a={"prova":3,"ciao":4};
	let b=a{"prova"};
	a{"ciao"}=5;
	let c=a{"ciao"};
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
	b, e := env.getVariable("b")
	if e != nil {
		t.Error(e)
	}
	v := b.(numberObject)
	if e != nil {
		t.Error(e)
	}
	if v.Value != 3 {
		t.Error("should be '3' ,got:", v.Value)
	}
	c, e := env.getVariable("c")
	if e != nil {
		t.Error(e)
	}
	v = c.(numberObject)
	if e != nil {
		t.Error(e)
	}
	if v.Value != 5 {
		t.Error("should be '3' ,got:", v.Value)
	}
}
func TestInlineFunc(t *testing.T) {
	ist := `
	let a=@(a,b){
		ret a+b;
	};
	let b=a(2,1);
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
	b, e := env.getVariable("b")
	if e != nil {
		t.Error(e)
	}
	v := b.(numberObject)
	if e != nil {
		t.Error(e)
	}
	if v.Value != 3 {
		t.Error("should be '1' ,got:", v.Value)
	}
}
func TestInlineFuncInHash(t *testing.T) {
	ist := `
	let a=@(a,b){
		ret a+b;
	};
	let b={"a":a};
	let c=b{"a"}(1,2);
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
	c, e := env.getVariable("c")
	if e != nil {
		t.Error(e)
	}
	v := c.(numberObject)
	if e != nil {
		t.Error(e)
	}
	if v.Value != 3 {
		t.Error("should be '1' ,got:", v.Value)
	}
}
func TestCallInlineFuncInDeclaration(t *testing.T) {
	ist := `
	let a=@(a,b){
		ret a+b;
	}(1,2);
	let b={"prova":@(){ret 7;}};
	let c=b{"prova"}();
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
	a, e := env.getVariable("a")
	if e != nil {
		t.Error(e)
	}
	v := a.(numberObject)
	if e != nil {
		t.Error(e)
	}
	if v.Value != 3 {
		t.Error("should be '3' ,got:", v.Value)
	}
	c, e := env.getVariable("c")
	if e != nil {
		t.Error(e)
	}
	v = c.(numberObject)
	if e != nil {
		t.Error(e)
	}
	if v.Value != 7 {
		t.Error("should be '7' ,got:", v.Value)
	}
}
func TestUseHashInInlineFunc(t *testing.T) {
	ist := `
	let object={
	   "nome":"luca",
	   "eta":22,
	   "compleanno":@(){
		this{"eta"}=this{"eta"}+1;
		ret this{"eta"};
	   }
	};
	object{"compleanno"}();
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
	a, e := env.getVariable("object")
	if e != nil {
		t.Error(e)
	}
	v := a.(hashObject)
	if e != nil {
		t.Error(e)
	}
	if v.Values[stringObject{Value: "eta"}].(numberObject).Value != 23 {
		t.Error("should be '23'")
	}
}
func TestUseinlineFuncFromarray(t *testing.T) {
	ist := `
	let array=[
	   @(){
		ret 2;
	   },
	   @(){
		ret 3;
	   }
	];
	let c1=array[0]();
	let c2=array[1]();
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
	c, e := env.getVariable("c1")
	if e != nil {
		t.Error(e)
	}
	v := c.(numberObject)
	if e != nil {
		t.Error(e)
	}
	if v.Value != 2 {
		t.Error("should be '2'")
	}
	c, e = env.getVariable("c2")
	if e != nil {
		t.Error(e)
	}
	v = c.(numberObject)
	if e != nil {
		t.Error(e)
	}
	if v.Value != 3 {
		t.Error("should be '3'")
	}
}
