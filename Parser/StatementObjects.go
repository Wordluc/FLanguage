package Parser

type IStatement interface {
	ToString() string
}
type StatementNode struct {
	Statement IStatement
	Next      *StatementNode
}

func (s *StatementNode) addNext(next *StatementNode) {
	s.Next = next
}

func (s *StatementNode) addStatement(statement IStatement) {
	s.Statement = statement
}
func (s StatementNode) ToString() string {
	r := "\t"
	if s.Statement != nil {
		r += s.Statement.ToString()
	}
	if s.Next != nil {
		r += "\n" + s.Next.ToString()
	}
	return r
}

type LetStatement struct {
	Identifier string
	Expresion  IExpresion
}

func (s LetStatement) ToString() string {
	return "LET " + s.Identifier + " = " + s.Expresion.ToString()
}

type IfStatement struct {
	Expresion IExpresion
	Body      IStatement
	Else      IStatement
}

func (s IfStatement) ToString() string {
	r := "IF ( "
	r += s.Expresion.ToString() + " ) "
	r += "{\n"
	if s.Body != nil {
		r += s.Body.ToString()
	}
	r += "\n} "
	if s.Else != nil {
		r += "ELSE {\n"
		r += s.Else.ToString()
		r += "\n}"
	}
	return r
}

type CallFuncStatement struct {
	Expresion IExpresion
}

func (s CallFuncStatement) ToString() string {
	return s.Expresion.ToString()
}

type AssignExpresionStatement struct {
	Identifier string
	Expresion  IExpresion
}

func (s AssignExpresionStatement) ToString() string {
	return s.Identifier + " = " + s.Expresion.ToString()
}

type FuncDeclarationStatement struct {
	Identifier string
	Body       IStatement
	Params     []string
}

func (s FuncDeclarationStatement) ToString() string {
	r := "Ff " + s.Identifier + " ( "
	for i := 0; i < len(s.Params); i++ {
		r += s.Params[i]
		if i < len(s.Params)-1 {
			r += ", "
		}
	}
	r += " ) {\n"
	if s.Body != nil {
		r += s.Body.ToString()
	}
	r += "\n}"
	return r
}

type ReturnStatement struct {
	Expresion IExpresion
}

func (s ReturnStatement) ToString() string {
	return "RETURN " + s.Expresion.ToString()
}

type SetArrayValueStatement struct {
	Identifier string
	Indexs     []IExpresion
	Value      IExpresion
}

func (s SetArrayValueStatement) ToString() string {
	r := s.Identifier + "["
	for i := 0; i < len(s.Indexs); i++ {
		r += s.Indexs[i].ToString()
		if i < len(s.Indexs)-1 {
			r += ", "
		}
	}
	r += "] = " + s.Value.ToString()
	return r
}

type SetHashValueStatement struct {
	Identifier string
	Index      IExpresion
	Value      IExpresion
}

func (s SetHashValueStatement) ToString() string {
	r := s.Identifier + "{"
	r += s.Index.ToString()
	r += "} = " + s.Value.ToString()
	return r
}

type WhileStatement struct {
	Cond IExpresion
	Body IStatement
}

func (s WhileStatement) ToString() string {
	r := "WHILE ( "
	r += s.Cond.ToString() + " ) {\n"
	if s.Body != nil {
		r += s.Body.ToString()
	}
	r += "\n}"
	return r
}
