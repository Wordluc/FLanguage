package Type

type TokenType int8

const (
	LET TokenType = iota
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
	OPEN_COM
	CLOSE_COM
	DOT
	COMMA
	DOT_COMMA
	END
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
		return OPEN_COM
	case "*/":
		return CLOSE_COM
	case ".":
		return DOT
	case ",":
		return COMMA
	case ";":
		return DOT_COMMA
	default:
		return VALUE
	}
}
