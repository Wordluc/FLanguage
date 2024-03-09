package lexer

import "strings"

type RegexToken struct {
	s strings.Builder
}

func (r *RegexToken) AddRegexPart(p string) {
	r.s.WriteString(p)
}
func (r RegexToken) Get() string {
	return r.s.String()
}

func GetRegex() RegexToken {
	r := RegexToken{}
	r.AddRegexPart(`\w+|`)                            //Parole
	r.AddRegexPart(`\/\*|\*\/|`)                      //Commenti multi linea
	r.AddRegexPart(`\/{1,2}|`)                        //Commenti mono linea
	r.AddRegexPart(`\+{1,2}|\{1,2}|\-{1,2}|\={1,2}|`) //Operazioni aritmentiche
	r.AddRegexPart(`!=|\>\=?|\<\=?|`)                 //Operazioni boleani
	r.AddRegexPart(`[(){}]|`)                         //Parentesi
	r.AddRegexPart(`[,.:;!?\"*]`)                     //Simboli
	return r
}
