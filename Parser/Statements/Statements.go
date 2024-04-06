package Statements

import "FLanguage/Parser/Expresions"

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
