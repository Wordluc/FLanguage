package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
)

func parseLeaf(l *Lexer.Lexer, _ IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {
	curToken := l.LookCurrent()
	leaf := ExpresionLeaf{}.New(curToken)

	if nextToken, _ := l.LookNext(); nextToken.Type == Token.OPEN_CIRCLE_BRACKET {
		l.IncrP()
		return parseCircleBracket(l, leaf, exitTokens...)
	}
	if nextToken, _ := l.LookNext(); nextToken.Type == Token.OPEN_GRAP_BRACKET {
		l.IncrP()
		return parseHash(l, leaf, exitTokens...)
	}
	l.IncrP()
	return leaf, nil
}
