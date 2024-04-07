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
	exp := [17]Token.TokenType{
		Token.LET, Token.WORD, Token.ASSIGN, Token.NUMBER, Token.DOT_COMMA,
		Token.NUMBER, Token.PLUS, Token.NUMBER,
		Token.FUNC, Token.CALL_FUNC, Token.CLOSE_CIRCLE_BRACKET, Token.OPEN_GRAP_BRACKET,
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
		Token.CALL_FUNC, Token.CLOSE_CIRCLE_BRACKET, Token.OPEN_GRAP_BRACKET, Token.CLOSE_GRAP_BRACKET, Token.WORD,
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
