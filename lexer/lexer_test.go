package lexer

import (
	"testing"

	"github.com/candy12t/go-json/token"
)

func TestLexer(t *testing.T) {
	input := `
{
  "Image": {
    "Width": 800,
    "Height": 600,
    "Title": "View from 15th Floor",
    "Thumbnail": {
      "Url": "http://www.example.com/image/481989943",
      "Height": 125,
      "Width": 100
    },
    "Animated": false,
    "IDs": [
      116,
      943,
      234,
      38793
    ]
  }
}
`

	tests := []struct {
		wantTokenType    token.Type
		wantTokenLiteral string
	}{
		{token.LBRACE, "{"},
		{token.STRING, "Image"},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.STRING, "Width"},
		{token.COLON, ":"},
		{token.NUMBER, "800"},
		{token.COMMA, ","},
		{token.STRING, "Height"},
		{token.COLON, ":"},
		{token.NUMBER, "600"},
		{token.COMMA, ","},
		{token.STRING, "Title"},
		{token.COLON, ":"},
		{token.STRING, "View from 15th Floor"},
		{token.COMMA, ","},
		{token.STRING, "Thumbnail"},
		{token.COLON, ":"},
		{token.LBRACE, "{"},
		{token.STRING, "Url"},
		{token.COLON, ":"},
		{token.STRING, "http://www.example.com/image/481989943"},
		{token.COMMA, ","},
		{token.STRING, "Height"},
		{token.COLON, ":"},
		{token.NUMBER, "125"},
		{token.COMMA, ","},
		{token.STRING, "Width"},
		{token.COLON, ":"},
		{token.NUMBER, "100"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.STRING, "Animated"},
		{token.COLON, ":"},
		{token.FALSE, "false"},
		{token.COMMA, ","},
		{token.STRING, "IDs"},
		{token.COLON, ":"},
		{token.LBRACKET, "["},
		{token.NUMBER, "116"},
		{token.COMMA, ","},
		{token.NUMBER, "943"},
		{token.COMMA, ","},
		{token.NUMBER, "234"},
		{token.COMMA, ","},
		{token.NUMBER, "38793"},
		{token.RBRACKET, "]"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	l := New(input)
	for i, tt := range tests {
		got := l.NextToken()

		if got.Type != tt.wantTokenType {
			t.Errorf("%d: got TokenType is %v, want TokenType is %v\n", i+1, got.Type, tt.wantTokenType)
		}

		if got.Literal != tt.wantTokenLiteral {
			t.Errorf("%d: got Literal is %q, want Literal is %q\n", i+1, got.Literal, tt.wantTokenLiteral)
		}

	}
}
