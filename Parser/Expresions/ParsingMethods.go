package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Attraction.go"
	"errors"
	"slices"
)

type fParse func(l *Lexer.Lexer, expresion IExpresion) (IExpresion, error)

func And(e error, s string) error {
	v := e.Error()
	return errors.New(v + " " + s)
}
func IsABrach(token Token.Token) bool {
	return token.Type == Token.WORD || token.Type == Token.OPEN_CIRCLE_BRACKET
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
	case Token.WORD:
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
		root, e = fVar(l, root)
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
func parseExpresionBlock(l *Lexer.Lexer, _ IExpresion) (IExpresion, error) {
	l.IncrP()
	block, e := ParseExpresion(l, Token.CLOSE_CIRCLE_BRACKET)
	l.IncrP()
	if e != nil {
		return nil, e
	}
	return block, nil
}
func parseLeaf(l *Lexer.Lexer, _ IExpresion) (IExpresion, error) {
	leaf := &ExpresionLeaf{}
	curToken := l.LookCurrent()
	if curToken.Type != Token.WORD {
		return nil, errors.New("ParseLeaf: not implemented,expected a word,got:" + curToken.Value)
	}
	l.IncrP()
	return leaf.New(curToken), nil
}
func parseTree(l *Lexer.Lexer, left IExpresion) (IExpresion, error) {
	tree := ExpresionNode{LeftExpresion: left}
	curOpToken := l.LookCurrent()
	powerCur, e := Attraction.GetForce(curOpToken.Type)
	if e != nil {
		return nil, e
	}
	if IsABrach(curOpToken) {
		return nil, errors.New("ParseTree: got a word instead of an operator")
	}
	if curOpToken.Type == Token.DOT_COMMA {
		return tree, nil
	}
	tree.SetOperator(curOpToken)
	lookNextVar, e := l.NextToken()
	if e != nil {
		return nil, e
	}
	fVar, e := GetParse(lookNextVar.Type)
	if e != nil {
		return nil, e
	}
	node, e := fVar(l, ExpresionLeaf{})
	if e != nil {
		return nil, e
	}
	if !IsABrach(lookNextVar) {
		return tree, errors.New("ParseTree: not implemented,expected a word,got:" + lookNextVar.Value)
	}
	lookNextOp := l.LookCurrent()
	if lookNextOp.Type == Token.DOT_COMMA || lookNextOp.Type == Token.CLOSE_CIRCLE_BRACKET {
		tree.SetRight(node)
		return tree, nil
	}
	if IsABrach(lookNextOp) {
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
		treeRigth, e := fop(l, node)
		if e != nil {
			return nil, e
		}
		tree.SetRight(treeRigth)
		return tree, nil
	}
	tree.SetRight(node)
	return tree, nil
}
