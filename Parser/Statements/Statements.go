package Statements

import (
	"FLanguage/Lexer/Token"
)

type IStatement interface {
	ToString() string
	TokenType() Token.TokenType
}
