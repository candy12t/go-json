package lexer

import "github.com/candy12t/go-json/token"

type Lexer struct {
	input    string
	char     byte
	position int
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.position >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.position]
	}
	l.position++
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	l.skipWhitespace()

	switch l.char {
	case '{':
		tk = token.New(token.LBRACE, l.char)
	case '}':
		tk = token.New(token.RBRACE, l.char)
	case '[':
		tk = token.New(token.LBRACKET, l.char)
	case ']':
		tk = token.New(token.RBRACKET, l.char)
	case ':':
		tk = token.New(token.COLON, l.char)
	case ',':
		tk = token.New(token.COMMA, l.char)
	case '"':
		tk.Type = token.STRING
		tk.Literal = l.readString()
	case 't', 'f', 'n':
		tk.Literal = l.readIndentifier()
		tk.Type = token.LookupIdentifier(tk.Literal)
		return tk
	case 0:
		tk.Type = token.EOF
		tk.Literal = ""
	default:
		if isNumber(l.char) {
			tk.Type = token.NUMBER
			tk.Literal = l.readNumber()
			return tk
		} else {
			tk = token.New(token.ILLEGAL, l.char)
		}
	}

	l.readChar()
	return tk
}

func (l *Lexer) readString() string {
	position := l.position
	for {
		l.readChar()
		if l.char == '"' || l.char == 0 {
			break
		}
	}
	return l.input[position : l.position-1]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isNumber(l.char) {
		l.readChar()
	}
	return l.input[position-1 : l.position-1]
}

func (l *Lexer) readIndentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position-1 : l.position-1]
}

func isNumber(char byte) bool {
	return ('0' <= char && char <= '9') || char == '-' || char == '.' || char == 'e'
}

func isLetter(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z')
}
