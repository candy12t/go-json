package json

import "github.com/candy12t/go-json/parser"

func NewParser(input string) *parser.Parser {
	return parser.New(input)
}
