package Lexer

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func ReplLexer() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-")
	v, _ := reader.ReadBytes('\n')
	if string(v) == "{{\r\n" {
		v = readBlockLines(reader)
	}
	l, _ := New(v)
	fmt.Println("--------------")
	for {
		t, e := l.NextToken()
		if e != nil {
			break
		}
		fmt.Println(t)
	}
	ReplLexer()
}
func readBlockLines(reader *bufio.Reader) []byte {
	var text []byte
	for {
		v, _ := reader.ReadBytes('\n')
		if string(v) == "}}\r\n" {
			return text
		}
		text = slices.Concat(text, v)
	}
}
