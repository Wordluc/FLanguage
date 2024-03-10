package Lexer

import (
	Token "FLanguage/Lexer/Token"
	"errors"
	"log"
	"os"
	"regexp"
)

type Lexer struct {
	input         []string
	pCurrectValue int
	pNextValue    int
}

func (l *Lexer) GetP() (int, int) {
	return l.pCurrectValue, l.pNextValue
}
func (l *Lexer) IncrP() error {
	if l.pCurrectValue == (len(l.input))-1 {
		return errors.New("no more token")
	}
	l.pCurrectValue = l.pNextValue
	l.pNextValue++
	return nil

}
func OpenFile(path string) ([]byte, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return nil, err
	}
	return file, nil
}
func New(text []byte) (Lexer, error) {
	regex := GetRegex()
	r := regexp.MustCompile(regex)
	matchs := r.FindAllString(string(text), -1)
	return Lexer{matchs, -1, 0}, nil

}
func (l *Lexer) NextToken(n int) (Token.Token, error) {
	e := l.IncrP()
	if e != nil {
		return Token.Token{}, e
	}
	ttype := Token.GetTokenType(l.input[l.pCurrectValue])
	return Token.New(l.input[l.pCurrectValue], ttype, n), nil

}
