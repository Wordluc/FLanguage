package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Attraction.go"
	"errors"
	"slices"
)

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
	if powerCur <= powerNext {
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
