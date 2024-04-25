package Evaluator

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Statements"
	"errors"
)

func Eval(program *Statements.StatementNode, env *Environment) (IObject, error) {
	r, e := evalStatement(program.Statement, env)
	if e != nil {
		return nil, e
	}
	_, isReturn := r.(ReturnObject)
	if isReturn {
		return r, nil
	}

	if program.Next == nil {
		return r, nil
	}
	if program.Next.Statement == nil {

		return r, nil
	}
	return Eval(program.Next, env)
}

func Run(path string, env *Environment) (IObject, error) {
	file, _ := Lexer.GetByteFromFile(path)
	l, e := Lexer.New(file)
	if e != nil {
		return nil, errors.New("Lexer:" + e.Error())
	}
	p, e := Statements.ParsingStatement(&l, Token.END)
	if e != nil {
		return nil, errors.New("Parser:" + e.Error())
	}
	r, e := Eval(p.(*Statements.StatementNode), env)
	if e != nil {
		return nil, errors.New("Eval:" + e.Error())
	}
	return r, nil
}
