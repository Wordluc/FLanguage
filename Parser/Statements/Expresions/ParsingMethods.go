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
	return token.Type == Token.WORD || token.Type == Token.OPEN_CIRCLE_BRACKET || token.Type == Token.NUMBER || token.Type == Token.STRING
}

func IsAValidOperator(token Token.Token) bool {
	return !IsAValidBrach(token)
}

func GetParse(than Token.TokenType) (fParse, error) {
	switch than {
	case Token.DIV:
		return parseTree, nil
	case Token.MULT, Token.MINUS:
		return parseTree, nil
	case Token.EQUAL, Token.NOT_EQUAL, Token.LESS, Token.GREATER, Token.LESS_EQUAL, Token.GREATER_EQUAL:
		return parseTree, nil
	case Token.PLUS:
		return parseTree, nil
	case Token.WORD, Token.NUMBER, Token.STRING, Token.BOOLEAN:
		return parseLeaf, nil
	case Token.OPEN_CIRCLE_BRACKET:
		return parseExpresionBlock, nil
	case Token.OPEN_SQUARE_BRACKET:
		return parseDeclareArray, nil
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
func parseGetValueArray(l *Lexer.Lexer) (IExpresion, error) {
	array := &ExpresionGetValueArray{}
	array.Name = l.LookCurrent().Value
	l.IncrP()
	l.IncrP()
	value, e := ParseExpresion(l, Token.CLOSE_SQUARE_BRACKET)
	if e != nil {
		return nil, e
	}
	array.ValueId = value
	l.IncrP()
	return array, nil

}
func parseDeclareArray(l *Lexer.Lexer, _ IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {
	array := &ExpresionDeclareArray{}
	l.IncrP()
	values, e := parseGroup(l, nil, Token.CLOSE_SQUARE_BRACKET, Token.COMMA)
	if e != nil {
		return nil, e
	}
	l.IncrP()
	array.Values = values
	return array, nil

}
func parseCallFunc(l *Lexer.Lexer, _ IExpresion, _ ...Token.TokenType) (IExpresion, error) {

	callFunc := &ExpresionCallFunc{NameFunc: l.LookCurrent().Value}
	l.IncrP()
	l.IncrP()
	parms, e := parseGroup(l, nil, Token.CLOSE_CIRCLE_BRACKET, Token.COMMA)
	callFunc.Values = parms
	if e != nil {
		return nil, e
	}
	l.IncrP()
	return callFunc, nil
}
func parseGroup(l *Lexer.Lexer, _ IExpresion, exist Token.TokenType, delimiter Token.TokenType) ([]IExpresion, error) {
	var values []IExpresion

	for {

		if exist == l.LookCurrent().Type {
			break
		}

		parm, e := ParseExpresion(l, delimiter, exist)
		if e != nil {
			return nil, e
		}
		if parm == nil {
			break
		}
		values = append(values, parm)
		if exist == l.LookCurrent().Type {
			break
		}
		l.IncrP()
	}
	return values, nil
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
	if l.LookCurrent().Type == Token.WORD {

		if nextT.Type == Token.OPEN_CIRCLE_BRACKET {
			return parseCallFunc(l, nil)
		}
		if nextT.Type == Token.OPEN_SQUARE_BRACKET {
			return parseGetValueArray(l)
		}
	}
	curToken := l.LookCurrent()
	leaf := &ExpresionLeaf{}
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
