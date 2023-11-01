package parser

import (
	"fmt"
	"strconv"

	"github.com/candy12t/go-json/ast"
	"github.com/candy12t/go-json/lexer"
	"github.com/candy12t/go-json/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(input string) *Parser {
	p := &Parser{
		l: lexer.New(input),
	}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Parse() (ast.Value, error) {
	switch p.curToken.Type {
	case token.LBRACE:
		return p.parseObect()
	case token.LBRACKET:
		return p.parseArray()
	case token.NUMBER:
		f, err := strconv.ParseFloat(p.curToken.Literal, 64)
		if err != nil {
			return nil, fmt.Errorf("parse error: invalid number: %w", err)
		}
		return &ast.NumberLiteral{Value: f}, nil
	case token.STRING:
		return &ast.StringLiteral{Value: p.curToken.Literal}, nil
	case token.TRUE:
		return &ast.BoolLiteral{Value: true}, nil
	case token.FALSE:
		return &ast.BoolLiteral{Value: false}, nil
	case token.NULL:
		return &ast.NullLiteral{}, nil
	default:
		return nil, fmt.Errorf("parse error: invalid Token: %v", p.curToken.Literal)
	}
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(tk token.Type) bool {
	return p.curToken.Type == tk
}

func (p *Parser) parseObect() (*ast.ObjectLiteral, error) {
	obj := new(ast.ObjectLiteral)
	obj.Pairs = make(map[ast.StringLiteral]ast.Value)

	for p.peekToken.Type != token.RBRACE {
		p.nextToken()
		if p.curToken.Type != token.STRING {
			return nil, fmt.Errorf("parse error: object key is not string: %v", p.curToken.Literal)
		}
		key := ast.StringLiteral{Value: p.curToken.Literal}

		p.nextToken()
		if p.curToken.Type != token.COLON {
			return nil, fmt.Errorf("parse error: object separator is not colon: %v", p.curToken.Literal)
		}

		p.nextToken()
		value, err := p.Parse()
		if err != nil {
			return nil, err
		}
		obj.Pairs[key] = value

		if p.peekToken.Type == token.COMMA {
			p.nextToken()
		}
	}

	p.nextToken()
	if p.curToken.Type != token.RBRACE {
		return nil, fmt.Errorf("parser error: object end is not right brace: %v", p.curToken.Literal)
	}

	return obj, nil
}

func (p *Parser) parseArray() (*ast.ArrayLiteral, error) {
	ary := new(ast.ArrayLiteral)
	ary.Values = []ast.Value{}

	for p.peekToken.Type != token.RBRACKET {
		p.nextToken()
		value, err := p.Parse()
		if err != nil {
			return nil, err
		}
		ary.Values = append(ary.Values, value)

		if p.peekToken.Type == token.COMMA {
			p.nextToken()
		}
	}

	p.nextToken()
	if p.curToken.Type != token.RBRACKET {
		return nil, fmt.Errorf("parser error: array end is not right brace: %v", p.curToken.Literal)
	}

	return ary, nil
}
