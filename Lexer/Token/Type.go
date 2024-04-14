package Token

import (
	"regexp"
)

type TokenType uint8

const (
	END TokenType = iota
	LET
	WORD
	NUMBER
	FUNC
	NONE
	RETURN
	STRING
	ASSIGN
	PLUS
	MINUS
	DIV
	MULT
	OPEN_CIRCLE_BRACKET
	CLOSE_CIRCLE_BRACKET
	OPEN_GRAP_BRACKET
	CLOSE_GRAP_BRACKET
	IF
	OPEN_COMM
	CLOSE_COMM
	DOT
	COMMA
	DOT_COMMA
	DOUBLE_QUOTE
	SINGLE_QUOTE
	GREATER
	LESS
	GREATER_EQUAL
	LESS_EQUAL
	NOT_EQUAL
	EQUAL
	ELSE
	BOOLEAN
	ERROR_L TokenType = 255
)

func GetTokenType(typeF string) TokenType {
	switch typeF {
	case "let":
		return LET
	case "Ff":
		return FUNC
	case "ret":
		return RETURN
	case "+":
		return PLUS
	case "*":
		return MULT
	case "-":
		return MINUS
	case "/":
		return DIV
	case "=":
		return ASSIGN
	case "(":
		return OPEN_CIRCLE_BRACKET
	case ")":
		return CLOSE_CIRCLE_BRACKET
	case "{":
		return OPEN_GRAP_BRACKET
	case "}":
		return CLOSE_GRAP_BRACKET
	case "if":
		return IF
	case "else":
		return ELSE
	case "END":
		return END
	case "/*":
		return OPEN_COMM
	case "*/":
		return CLOSE_COMM
	case ".":
		return DOT
	case ",":
		return COMMA
	case ";":
		return DOT_COMMA
	case ">":
		return GREATER
	case "<":
		return LESS
	case ">=":
		return GREATER_EQUAL
	case "<=":
		return LESS_EQUAL
	case "!=":
		return NOT_EQUAL
	case "==":
		return EQUAL
	case "true", "false":
		return BOOLEAN

	default:
		if isValidString(typeF) {
			return STRING
		}
		if isValidWord(typeF) {
			return WORD
		}
		if isValidNumber(typeF) {
			return NUMBER
		}

		return ERROR_L
	}
}

var isAWord = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`)

var isANumber = regexp.MustCompile(`^[0-9_]+$`)

func isValidWord(value string) bool {
	return isAWord.MatchString(string(value))
}

func isValidNumber(value string) bool {
	return isANumber.MatchString(string(value))
}

func isValidString(value string) bool {
	if value[0] == '"' && value[len(value)-1] == '"' {
		return true
	}
	if value[0] == '\'' && value[len(value)-1] == '\'' {
		return true
	}
	return false
}
