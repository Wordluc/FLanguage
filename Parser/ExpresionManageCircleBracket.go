package Parser

import (
	"FLanguage/Lexer"
	"FLanguage/Lexer/Token"
)

func parseCircleBracket(l *Lexer.Lexer, expresion IExpresion, exitTokens ...Token.TokenType) (IExpresion, error) {

	preToken, e := l.LookBack()
	if e != nil {
		return parseExpresionBlock(l, expresion, exitTokens...)
	}
	if isACallFunc(preToken) {
		return expresionCallFunc(l, expresion, exitTokens...)
	}
	return parseExpresionBlock(l, expresion, exitTokens...)

}

func isACallFunc(token Token.Token) bool {
	return token.Type == Token.WORD || token.Type == Token.CLOSE_CIRCLE_BRACKET || token.Type == Token.CLOSE_GRAP_BRACKET || token.Type == Token.CLOSE_SQUARE_BRACKET
}
