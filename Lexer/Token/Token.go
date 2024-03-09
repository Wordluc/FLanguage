package Lexer

import TToken "FLanguage/Lexer/Token/Type"

type Token struct {
	Type  TToken.TokenType
	Value string
	NLine int
}

func New(value string, tokenType TToken.TokenType, nLine int) Token {
	return Token{tokenType, value, nLine}
}
