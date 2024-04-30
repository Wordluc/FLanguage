package Lexer

import (
	"strings"
)

func GetRegex() string {
	r := strings.Builder{}
	r.WriteString(`[0-9]*\.[0-9]+|`)                 //Numeri con virgola
	r.WriteString(`\w+|`)                            //Parole e numeri interi
	r.WriteString(`\/\*|\*\/|`)                      //Commenti multi linea
	r.WriteString(`\/\/[^\n]*|`)                     //Commenti mono linea
	r.WriteString(`\+{1,2}|\{1,2}|\-{1,2}|\={1,2}|`) //Operazioni aritmentiche
	r.WriteString(`!=|\>\=?|\<\=?|`)                 //Operazioni boleani
	r.WriteString(`[(){}\[\]]|`)
	r.WriteString(`["][^"]*["]|`)
	r.WriteString(`['][^']+[']|`)
	r.WriteString(`[,.:;!?\"*/@]`) //Simboli
	return r.String()
}
