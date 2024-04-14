package Lexer

import "strings"

func GetRegex() string {
	r := strings.Builder{}
	//	r.WriteString(`\w+\(|`)                          //Parole
	r.WriteString(`\w+|`)                            //Parole
	r.WriteString(`\/\*|\*\/|`)                      //Commenti multi linea
	r.WriteString(`\/{1,2}|`)                        //Commenti mono linea
	r.WriteString(`\+{1,2}|\{1,2}|\-{1,2}|\={1,2}|`) //Operazioni aritmentiche
	r.WriteString(`!=|\>\=?|\<\=?|`)                 //Operazioni boleani
	r.WriteString(`[(){}]|`)
	r.WriteString(`["][^"]+["]|`)
	r.WriteString(`['][^']+[']|`)
	r.WriteString(`[,.:;!?\"*]`) //Simboli
	return r.String()
}
