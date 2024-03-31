package Expresions

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
	"FLanguage/Parser/Attraction.go"
	"errors"
)

type fParse func(l *Lexer.Lexer, expresion IExpresion) (IExpresion, error)

func And(e error, s string) error {
	v := e.Error()
	return errors.New(v + " " + s)
}
func IsWord(token Token.Token) bool {
	return token.Type == Token.WORD
}
func GetParse(than Token.TokenType) (fParse, error) {
	switch than {
	case Token.DIV:
		return ParseTree, nil
	case Token.MULT:
		return ParseTree, nil
	case Token.MINUS:
		return ParseTree, nil
	case Token.PLUS:
		return ParseTree, nil
	case Token.WORD:
		return ParseLeaf, nil
	}
	return nil, errors.New("GetParse: Operator:" + string(than) + "not implemented")
}
func ParseExpresion(l *Lexer.Lexer) (IExpresion, error) {
	var root IExpresion
	for {
		lookCurrVar := l.LookCurrent()
		if lookCurrVar.Type == Token.DOT_COMMA {
			break
		}
		fVar, e := GetParse(lookCurrVar.Type)
		if e != nil {
			return nil, e
		}
		node, e := fVar(l, root)
		if e != nil {
			return nil, e
		}
		switch node.(type) {
		case ExpresionNode:
			root = node
		case ExpresionLeaf:
			root = node
		}

		lookCurrVar = l.LookCurrent()
		if lookCurrVar.Type == Token.DOT_COMMA {
			break
		}
	}
	return root, nil
}
func ParseLeaf(l *Lexer.Lexer, _ IExpresion) (IExpresion, error) {
	leaf := &ExpresionLeaf{}
	curToken := l.LookCurrent()
	if curToken.Type != Token.WORD {
		return nil, errors.New("ParseLeaf: not implemented,expected a word,got:" + curToken.Value)
	}
	l.IncrP()
	return leaf.New(curToken), nil
}
func ParseTree(l *Lexer.Lexer, left IExpresion) (IExpresion, error) {
	tree := ExpresionNode{LeftExpresion: left}
	curOpToken := l.LookCurrent()
	powerCur, e := Attraction.GetForce(curOpToken.Type)
	if e != nil {
		return nil, e
	}
	if IsWord(curOpToken) {
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
	if !IsWord(lookNextVar) {
		return nil, errors.New("ParseTree: got a operator instead of an word")
	}
	fVar, e := GetParse(lookNextVar.Type)
	if e != nil {
		return nil, e
	}
	node, e := fVar(l, ExpresionLeaf{})
	if e != nil {
		return nil, e
	}
	lookNextOp := l.LookCurrent()
	if lookNextOp.Type == Token.DOT_COMMA {
		tree.SetRight(node)
		return tree, nil
	}
	if IsWord(lookNextOp) {
		return nil, errors.New("ParseTree: got a word instead of an operator")
	}
	powerNext, e := Attraction.GetForce(lookNextOp.Type)
	if e != nil {
		return nil, e
	}
	if powerCur < powerNext {
		fop, e := GetParse(lookNextOp.Type)
		if e != nil {
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
