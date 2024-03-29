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
func (l *Lexer) GetAll() []string {
	return l.input
}
func GetByteFromFile(path string) ([]byte, error) {
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
	return Lexer{matchs, 0, 1}, nil

}
func (l *Lexer) NextToken() (Token.Token, error) {
	e := l.IncrP()
	if e != nil {
		return Token.Token{}, e
	}
	ttype := Token.GetTokenType(l.input[l.pCurrectValue])
	return Token.New(l.input[l.pCurrectValue], ttype), nil

}
func (l *Lexer) LookNext() (Token.Token, error) {
	if l.pCurrectValue == (len(l.input))-1 {
		return Token.Token{}, errors.New("no more token")
	}
	ttype := Token.GetTokenType(l.input[l.pNextValue])
	return Token.New(l.input[l.pNextValue], ttype), nil

}
func (l *Lexer) LookBack() (Token.Token, error) {
	if l.pCurrectValue-1 < 0 {
		return Token.Token{}, errors.New("no back token")
	}
	ttype := Token.GetTokenType(l.input[l.pCurrectValue-1])
	return Token.New(l.input[l.pCurrectValue-1], ttype), nil

}
func (l *Lexer) LookCurrent() Token.Token {
	ttype := Token.GetTokenType(l.input[l.pCurrectValue])
	return Token.New(l.input[l.pCurrectValue], ttype)

}
