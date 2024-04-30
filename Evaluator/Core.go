package Evaluator

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser"
	"errors"
)

func Eval(program *Parser.StatementNode, env *Environment) (iObject, error) {
	r, e := evalStatement(program.Statement, env)
	if e != nil {
		return nil, e
	}
	_, isReturn := r.(returnObject)
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

func Run(path string, env *Environment) (iObject, error) {
	file, _ := Lexer.GetByteFromFile(path)
	l, e := Lexer.New(file)
	if e != nil {
		return nil, errors.New("Lexer:" + e.Error())
	}
	p, e := Parser.ParsingStatement(&l, Token.END)
	if e != nil {
		return nil, errors.New("Parser:" + e.Error())
	}
	r, e := Eval(p.(*Parser.StatementNode), env)
	if e != nil {
		return nil, errors.New("Eval:" + e.Error())
	}
	return r, nil
}
