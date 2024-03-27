package Parser

//func checkTypeStatement(t *testing.T, statement ParsingStatements.IStatement, expected []Token.TokenType) {
//	t.Log("pro", statement.GetStatement())
//	if statement.GetTokenType() != expected[0] {
//		t.Errorf("errore parsing: got %v instead %v", statement.GetTokenType(), expected[0])
//	}
//	head := statement.GetExpresion()
//	for _, i := range expected {
//		got := head.TypeToken
//		t.Log(got)
//		if got != i {
//			t.Errorf("errore parsing: got %v instead %v", got, i)
//		}
//		head = *head.NextExpresion
//	}
//}
//func TestParser_Sum(t *testing.T) {
//	ist := "let a=3+4;END"
//	lexer, e := Lexer.New([]byte(ist))
//	if e != nil {
//		//t.Error(e)
//	}
//	program, e := ParseProgram(&lexer)
//	if e != nil {
//		t.Error(e)
//	}
//	if program == nil {
//		t.Error("program is nil")
//	}
//	t.Log(program.String())
//	expected := [8]Token.TokenType{
//		Token.LET, Token.WORD, Token.EQUAL, Token.WORD, Token.PLUS, Token.WORD, Token.DOT_COMMA, Token.END}
//	checkTypeStatement(t, program.Statement, expected[:])
//}
