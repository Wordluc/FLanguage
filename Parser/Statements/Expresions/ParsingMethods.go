package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Attraction.go"
	"errors"
	"slices"
)

type fParse func(l *Lexer.Lexer, expresion IExpresion, exitTokens ...Token.TokenType) (IExpresion, error)

func And(e error, s string) error {
	v := e.Error()
	return errors.New(v + " " + s)
}

func IsAValidBrach(token Token.Token) bool {
	return token.Type == Token.WORD || token.Type == Token.OPEN_CIRCLE_BRACKET || token.Type == Token.NUMBER
}

func IsAValidOperator(token Token.Token) bool {
	return !IsAValidBrach(token)
}

func GetParse(than Token.TokenType) (fParse, error) {
	switch than {
	case Token.DIV:
		return parseTree, nil
	case Token.MULT:
		return parseTree, nil
	case Token.MINUS:
		return parseTree, nil
	case Token.PLUS:
		return parseTree, nil
	case Token.WORD, Token.NUMBER, Token.STRING:
		return parseLeaf, nil
	case Token.OPEN_CIRCLE_BRACKET:
		return parseExpresionBlock, nil
	}
	return nil, errors.New("GetParse: Operator:" + string(than) + "not implemented")
}

func ParseExpresion(l *Lexer.Lexer, exitTokens ...Token.TokenType) (IExpresion, error) {
	var root IExpresion
	if exitTokens == nil {
		return nil, errors.New("ParseExpresion: no exitTokens defined")
	}
	for {
		lookCurrVar := l.LookCurrent()
		if slices.Contains(exitTokens, lookCurrVar.Type) {
			break
		}
		fVar, e := GetParse(lookCurrVar.Type)
		if e != nil {
			return nil, e
		}
		root, e = fVar(l, root, exitTokens...)
		if e != nil {
			return nil, e
		}
		lookCurrVar = l.LookCurrent()
		if slices.Contains(exitTokens, lookCurrVar.Type) {
			break
		}

	}
	return root, nil
}

func parseCallFunc(l *Lexer.Lexer, _ IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {

	callFunc := &ExpresionCallFunc{NameFunc: l.LookCurrent().Value}
	l.IncrP()
	l.IncrP()
	for {
		if l.LookCurrent().Type == Token.CLOSE_CIRCLE_BRACKET {
			break
		}

		parm, e := ParseExpresion(l, Token.COMMA, Token.CLOSE_CIRCLE_BRACKET)
		if e != nil {
			return nil, e
		}
		if parm == nil {
			break
		}
		callFunc.AddParm(parm)
		if l.LookCurrent().Type == Token.CLOSE_CIRCLE_BRACKET {
			break
		}
		l.IncrP()
	}
	l.IncrP()
	return callFunc, nil
}

func parseExpresionBlock(l *Lexer.Lexer, _ IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {
	l.IncrP()
	block, e := ParseExpresion(l, Token.CLOSE_CIRCLE_BRACKET)
	if e != nil {
		return nil, e
	}
	l.IncrP()
	return block, nil
}

func parseLeaf(l *Lexer.Lexer, _ IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {
	nextT, e := l.LookNext()
	if e != nil {
		return nil, e
	}
	if nextT.Type == Token.OPEN_CIRCLE_BRACKET {
		return parseCallFunc(l, nil)
	}
	leaf := &ExpresionLeaf{}
	curToken := l.LookCurrent()
	l.IncrP()
	return leaf.New(curToken), nil
}

func parseTree(l *Lexer.Lexer, left IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {
	tree := ExpresionNode{LeftExpresion: left}
	curOpToken := l.LookCurrent()
	powerCur, e := Attraction.GetForce(curOpToken.Type)
	if e != nil {
		return nil, e
	}
	if !IsAValidOperator(curOpToken) {
		return nil, errors.New("ParseTree: got a word instead of an operator")
	}
	if slices.Contains(exitTokens, curOpToken.Type) {
		return tree, nil
	}
	tree.SetOperator(curOpToken)
	lookNextBranch, e := l.NextToken()
	if e != nil {
		return nil, e
	}
	if !IsAValidBrach(lookNextBranch) {
		return nil, errors.New("ParseTree: not implemented,expected a brach,got:" + lookNextBranch.Value)
	}
	fBranch, e := GetParse(lookNextBranch.Type)
	if e != nil {
		return nil, e
	}
	branch, e := fBranch(l, ExpresionLeaf{})
	if e != nil {
		return nil, e
	}

	lookNextOp := l.LookCurrent()
	if slices.Contains(exitTokens, lookNextOp.Type) {
		tree.SetRight(branch)
		return tree, nil
	}
	if !IsAValidOperator(lookNextOp) {
		return nil, errors.New("ParseTree: got a word instead of an operator")
	}

	powerNext, e := Attraction.GetForce(lookNextOp.Type)
	if e != nil {
		return nil, e
	}
	if powerCur < powerNext {
		fop, e := GetParse(lookNextOp.Type)
		if e != nil || fop == nil {
			return nil, e
		}
		treeRigth, e := fop(l, branch, exitTokens...)
		if e != nil {
			return nil, e
		}
		tree.SetRight(treeRigth)
		return tree, nil
	}
	tree.SetRight(branch)
	return tree, nil
}
