package ParsingStatements

import "FLanguage/Lexer/Token"

type IStatement interface {
	GetStatement() string
	GetTokenType() Token.TokenType
}
