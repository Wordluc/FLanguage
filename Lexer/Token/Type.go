package Token

type TokenType uint8

const (
	ERROR_L TokenType = 255
	END     TokenType = iota
	LET
	VALUE
	FUNC
	RETURN
	PLUS
	MINUS
	DIV
	MULT
	EQUAL
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
	case "\"":
		return DOUBLE_QUOTE
	case "'":
		return SINGLE_QUOTE
	default:
		if isValidValua(typeF) {
			return VALUE
		}
		return ERROR_L
	}
}
func isValidValua(value string) bool {
	for _, cr := range value {
		c := string(cr)
		if !((c >= "a" && c <= "z") || (c >= "A" && c <= "Z") || (c >= "0" && c <= "9")) {
			return false
		}
	}
	return true
}
