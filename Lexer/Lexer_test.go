package Lexer

import (
	"FLanguage/Lexer/Token"
	"testing"
)

func TestLexer(t *testing.T) {
	text := `
	let a =4;
	3+4
	Ff prova(){
		ret e;
	}
	END
	`
	exp := [18]Token.TokenType{
		Token.LET, Token.WORD, Token.ASSIGN, Token.NUMBER, Token.DOT_COMMA,
		Token.NUMBER, Token.PLUS, Token.NUMBER,
		Token.FUNC, Token.WORD, Token.OPEN_CIRCLE_BRACKET, Token.CLOSE_CIRCLE_BRACKET, Token.OPEN_GRAP_BRACKET,
		Token.RETURN, Token.WORD, Token.DOT_COMMA,
		Token.CLOSE_GRAP_BRACKET, Token.END,
	}
	l, e := New([]byte(text))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	for _, i := range exp {
		got := l.LookCurrent()
		if got.Type != Token.TokenType(i) {
			t.Errorf("errore parsing: got %v instead %v", got.Type, Token.TokenType(i))
		}
		l.IncrP()

	}
}
func TestLexerNumberWithDot(t *testing.T) {
	text := `
	33.3
	`
	l, e := New([]byte(text))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	got := l.LookCurrent()
	if got.Type != Token.NUMBER_WITH_DOT {
		t.Errorf("errore parsing: got %v instead %v", got.Type, Token.NUMBER_WITH_DOT)
	}
}
func TestLexer_Op(t *testing.T) {
	text := `
	*-/+;
	`
	exp := []Token.TokenType{
		Token.MULT, Token.MINUS, Token.DIV, Token.PLUS, Token.DOT_COMMA,
	}
	l, e := New([]byte(text))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	for _, i := range exp {
		got := l.LookCurrent()
		if got.Type != Token.TokenType(i) {
			t.Errorf("errore parsing: got %v instead %v", got.Type, Token.TokenType(i))
		}
		l.IncrP()

	}
}

func TestLexer_CallFunc(t *testing.T) {
	text := `
	Prova(){
		
	}
	cia1o
	`
	exp := []Token.TokenType{
		Token.WORD, Token.OPEN_CIRCLE_BRACKET, Token.CLOSE_CIRCLE_BRACKET, Token.OPEN_GRAP_BRACKET, Token.CLOSE_GRAP_BRACKET, Token.WORD,
	}
	l, e := New([]byte(text))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	for _, i := range exp {
		got := l.LookCurrent()
		if got.Type != Token.TokenType(i) {
			t.Errorf("errore parsing: got %v instead %v", got.Type, Token.TokenType(i))
		}
		l.IncrP()

	}
}

func TestWordError(t *testing.T) {
	text := `
	33r
	`
	l, e := New([]byte(text))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	got := l.LookCurrent()
	if got.Type != Token.ERROR_L {
		t.Errorf("errore parsing: got %v instead %v", got.Type, Token.STRING)
	}
}

func TestString(t *testing.T) {
	text := `
	"prova" "fff"
	'prova"vvv"'
	`

	l, e := New([]byte(text))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	exp := [3]Token.TokenType{
		Token.STRING, Token.STRING, Token.STRING,
	}
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	for _, i := range exp {
		got := l.LookCurrent()
		if got.Type != Token.TokenType(i) {
			t.Errorf("errore parsing: got %v instead %v", got.Type, Token.TokenType(i))
		}
		l.IncrP()

	}
}

func TestCompare(t *testing.T) {
	text := `
	== != < > <= >=
	`

	l, e := New([]byte(text))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	exp := [6]Token.TokenType{
		Token.EQUAL, Token.NOT_EQUAL, Token.LESS, Token.GREATER, Token.LESS_EQUAL, Token.GREATER_EQUAL,
	}
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	for _, i := range exp {
		got := l.LookCurrent()
		if got.Type != Token.TokenType(i) {
			t.Errorf("errore parsing: got %v instead %v", got.Type, Token.TokenType(i))
		}
		l.IncrP()

	}
}

func TestFuncDefinition(t *testing.T) {
	text := `
 	Ff (a, b, c) {
 		print();
	}
 	END
	`

	l, e := New([]byte(text))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	exp := [15]Token.TokenType{
		Token.FUNC, Token.OPEN_CIRCLE_BRACKET, Token.WORD, Token.COMMA, Token.WORD, Token.COMMA, Token.WORD, Token.CLOSE_CIRCLE_BRACKET,
		Token.OPEN_GRAP_BRACKET, Token.WORD, Token.OPEN_CIRCLE_BRACKET, Token.CLOSE_CIRCLE_BRACKET, Token.DOT_COMMA, Token.CLOSE_GRAP_BRACKET, Token.END,
	}
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	for _, i := range exp {
		got := l.LookCurrent()
		if got.Type != Token.TokenType(i) {
			t.Errorf("errore parsing: got %v instead %v", got.Type, Token.TokenType(i))
		}
		l.IncrP()

	}
}

func TestBoolean(t *testing.T) {
	text := `
 	2>3
	true
 	END
	`

	l, e := New([]byte(text))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	exp := [5]Token.TokenType{
		Token.NUMBER, Token.GREATER, Token.NUMBER,
		Token.BOOLEAN,
		Token.END,
	}
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	for _, i := range exp {
		got := l.LookCurrent()
		if got.Type != Token.TokenType(i) {
			t.Errorf("errore parsing: got %v instead %v", got.Type, Token.TokenType(i))
		}
		l.IncrP()

	}
}
func TestArrayOp(t *testing.T) {
	text := `
 	a[2]
	d=[1,2,3]	
	`

	l, e := New([]byte(text))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	exp := [13]Token.TokenType{
		Token.WORD, Token.OPEN_SQUARE_BRACKET, Token.NUMBER, Token.CLOSE_SQUARE_BRACKET,
		Token.WORD, Token.ASSIGN, Token.OPEN_SQUARE_BRACKET, Token.NUMBER, Token.COMMA, Token.NUMBER, Token.COMMA, Token.NUMBER, Token.CLOSE_SQUARE_BRACKET}
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	for _, i := range exp {
		got := l.LookCurrent()
		if got.Type != Token.TokenType(i) {
			t.Errorf("errore parsing: got %v instead %v", got.Type, Token.TokenType(i))
		}
		l.IncrP()

	}
}
