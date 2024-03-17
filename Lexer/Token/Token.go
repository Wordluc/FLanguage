package Token

type Token struct {
	Type  TokenType
	Value string
}

func New(value string, tokenType TokenType) Token {
	return Token{tokenType, value}
}
