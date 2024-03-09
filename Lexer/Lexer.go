package lexer

import (
	Token "FLanguage/Lexer/Token"
	TokenT "FLanguage/Lexer/Token/Type"
	"bufio"
	"log"
	"os"
	"regexp"
)

type ParseResult struct {
	tokens []Token.Token
}

func Parse(path string) (ParseResult, error) {
	f, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return ParseResult{}, err
	}
	defer f.Close()
	scan := bufio.NewScanner(f)
	result := ParseResult{}
	regex := GetRegex()
	r := regexp.MustCompile(regex.Get())
	n := 0
	for scan.Scan() {
		var line = scan.Text()
		result.parseLine(line, r, n)
		n++
	}
	return result, nil
}
func (r *ParseResult) parseLine(line string, source *regexp.Regexp, n int) {
	matchs := source.FindAllString(line, -1)
	for _, match := range matchs {

		ttype := TokenT.GetTokenType(match)
		r.tokens = append(r.tokens, Token.New(match, ttype, n))
	}
}
