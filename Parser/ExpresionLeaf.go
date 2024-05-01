package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
)

func parseLeaf(l *Lexer.Lexer, _ IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {
	curToken := l.LookCurrent()
	leaf := ExpresionLeaf{}.New(curToken)

	if nextToken, _ := l.NextToken(); nextToken.Type == Token.OPEN_CIRCLE_BRACKET {
		return parseCircleBracket(l, leaf, exitTokens...)
	}
	return leaf, nil
}
