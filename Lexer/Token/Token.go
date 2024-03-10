package Token

type Token struct {
	Type  TokenType
	Value string
	NLine int
}

func New(value string, tokenType TokenType, nLine int) Token {
	return Token{tokenType, value, nLine}
}
