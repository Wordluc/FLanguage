package Token

import (
	"strings"
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
	PLUS
	MINUS
	DIV
	MULT
	EQUAL
	OPEN_CIRCLE_BRACKET
	CLOSE_CIRCLE_BRACKET
	OPEN_GRAP_BRACKET
	CLOSE_GRAP_BRACKET
	CALL_FUNC
	IF
	OPEN_COMM
	CLOSE_COMM
	DOT
	COMMA
	DOT_COMMA
	DOUBLE_QUOTE
	SINGLE_QUOTE
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
		return EQUAL
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
	default:
		if isACallFunc(typeF) {
			return CALL_FUNC
		}
		if isValidWord(typeF) {
			return WORD
		}
		if isValidNumber(typeF) {
			return NUMBER
		}
		if isAString(typeF) {
			return STRING
		}
		return ERROR_L
	}
}
func isValidWord(value string) bool {
	for _, cr := range value {
		c := string(cr)
		if !((c >= "a" && c <= "z") || (c >= "A" && c <= "Z")) {
			return false
		}
	}
	return true
}
func isValidNumber(value string) bool {
	for _, cr := range value {
		c := string(cr)
		if !(c >= "0" && c <= "9") {
			return false
		}
	}
	return true
}
func isACallFunc(value string) bool {
	parts := strings.Split(value, "(")
	if len(parts) != 2 {
		return false
	}
	if !isValidWord(parts[0]) {
		return false
	}
	return true
}
func isAString(value string) bool {
	parts := strings.Split(value, "\"")
	if len(parts) != 3 {
		return false
	}
	return true
}
