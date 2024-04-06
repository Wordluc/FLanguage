package Statements

import (
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Expresions"
)

type IStatement interface {
	ToString() string
}
type StatementNode struct {
	Statement IStatement
	Next      *StatementNode
}

func (s *StatementNode) AddNext(next *StatementNode) {
	s.Next = next
}
func (s *StatementNode) AddStatement(statement IStatement) {
	s.Statement = statement
}
func (s *StatementNode) ToString() string {
	r := "\t"
	if s.Statement != nil {
		r += s.Statement.ToString()
	}
	if s.Next != nil {
		r += "\n" + s.Next.ToString()
	}
	return r
}

type LetStatement struct {
	Identifier string
	Expresion  Expresions.IExpresion
}

func (s LetStatement) ToString() string {
	return "LET " + s.Identifier + " = " + s.Expresion.ToString()
}

type IfStatement struct {
	FirstExpresion Expresions.IExpresion
	ConditionType  Token.TokenType
	ConditionValue string
	LastExpresion  Expresions.IExpresion
	Body           IStatement
	Else           IStatement
}

func (s IfStatement) ToString() string {
	r := "IF ( "
	r += s.FirstExpresion.ToString() + " "
	r += s.ConditionValue + " "
	r += s.LastExpresion.ToString() + " ) "
	r += "{\n"
	if s.Body != nil {
		r += s.Body.ToString()
	}
	r += "\n} "
	if s.Else != nil {
		r += "ELSE {\n"
		r += s.Else.ToString()
		r += "\n}"
	}
	return r
}

type CallFuncStatement struct {
	Expresion Expresions.IExpresion
}

func (s CallFuncStatement) ToString() string {
	return s.Expresion.ToString()
}

type AssignExpresionStatement struct {
	Identifier string
	Expresion  Expresions.IExpresion
}

func (s AssignExpresionStatement) ToString() string {
	return s.Identifier + " = " + s.Expresion.ToString()
}
