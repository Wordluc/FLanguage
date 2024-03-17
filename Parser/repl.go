package Parser

import (
	"FLanguage/Lexer"
	"bufio"
	"fmt"
	"os"
	"slices"
)

func ReplParser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-")
	v, _ := reader.ReadBytes('\n')
	if string(v) == "{{\r\n" {
		v = readBlockLines(reader)
	}
	l, _ := Lexer.New(v)
	fmt.Println("--------------")
	for {
		statement, e := parseStatement(&l)

		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(statement.getStatement())
	}
	ReplParser()
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
