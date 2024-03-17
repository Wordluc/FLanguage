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
	`
	exp := [17]Token.TokenType{
		Token.LET, Token.WORD, Token.EQUAL, Token.WORD, Token.DOT_COMMA,
		Token.WORD, Token.PLUS, Token.WORD,
		Token.FUNC, Token.WORD, Token.OPEN_CIRCLE_BRACKET, Token.CLOSE_CIRCLE_BRACKET, Token.OPEN_GRAP_BRACKET,
		Token.RETURN, Token.WORD, Token.DOT_COMMA,
		Token.CLOSE_GRAP_BRACKET,
	}
	l, e := New([]byte(text))
	if e != nil {
		t.Error("creazione Lexer fallita")
	}
	for _, i := range exp {
		got, _ := l.NextToken()
		if got.Type != Token.TokenType(i) {
			t.Errorf("errore parsing: got %v instead %v", got.Type, Token.TokenType(i))
		}

	}
}
