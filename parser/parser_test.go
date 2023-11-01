package parser

import (
	"testing"

	"github.com/candy12t/go-json/ast"
	"github.com/google/go-cmp/cmp"
)

func TestParseBool(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"true", true},
		{"false", false},
	}

	for _, tt := range tests {
		p := New(tt.input)
		obj, _ := p.Parse()

		got, ok := obj.(*ast.BoolLiteral)
		if !ok {
			t.Fatalf("not *ast.BoolLiteral, got is %T", obj)
		}
		if got.Value != tt.want {
			t.Errorf("got ast.BoolLiteral is %v, want ast.BoolLiteral is %v", got.Value, tt.want)
		}
	}
}

func TestParseString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{`"abc"`, "abc"},
		{`""`, ""},
	}

	for _, tt := range tests {
		p := New(tt.input)
		obj, _ := p.Parse()

		got, ok := obj.(*ast.StringLiteral)
		if !ok {
			t.Fatalf("not *ast.StringLiteral, got is %T", obj)
		}
		if got.Value != tt.want {
			t.Errorf("got StringLiteral is %v, want StringLiteral is %v", got.Value, tt.want)
		}
	}
}

func TestParseNumber(t *testing.T) {
	tests := []struct {
		input string
		want  float64
	}{
		{"0", 0},
		{"123", 123},
		{"123.4", 123.4},
		{"-123.4", -123.4},
		{"1e5", 100000},
	}

	for _, tt := range tests {
		p := New(tt.input)
		obj, _ := p.Parse()

		got, ok := obj.(*ast.NumberLiteral)
		if !ok {
			t.Fatalf("not *ast.NumberLiteral, got is %T", obj)
		}
		if got.Value != tt.want {
			t.Errorf("got NumberLiteral is %v, want NumberLiteral is %v", got.Value, tt.want)
		}
	}
}

func TestParseObject(t *testing.T) {
	tests := []struct {
		input string
		want  ast.ObjectLiteral
	}{
		{`{}`, ast.ObjectLiteral{Pairs: map[ast.StringLiteral]ast.Value{}}},
		{`{"key": "value"}`, ast.ObjectLiteral{
			Pairs: map[ast.StringLiteral]ast.Value{
				{Value: "key"}: &ast.StringLiteral{Value: "value"},
			},
		}},
		{`{"key1": "abc", "key2": -123.4, "key3": true, "key4": null, "key5": {"key5-1": "abc", "key5-2": -123.4}}`, ast.ObjectLiteral{
			Pairs: map[ast.StringLiteral]ast.Value{
				{Value: "key1"}: &ast.StringLiteral{Value: "abc"},
				{Value: "key2"}: &ast.NumberLiteral{Value: -123.4},
				{Value: "key3"}: &ast.BoolLiteral{Value: true},
				{Value: "key4"}: &ast.NullLiteral{},
				{Value: "key5"}: &ast.ObjectLiteral{
					Pairs: map[ast.StringLiteral]ast.Value{
						{Value: "key5-1"}: &ast.StringLiteral{Value: "abc"},
						{Value: "key5-2"}: &ast.NumberLiteral{Value: -123.4},
					},
				},
			},
		}},
	}

	for _, tt := range tests {
		p := New(tt.input)
		obj, _ := p.Parse()

		got, ok := obj.(*ast.ObjectLiteral)
		if !ok {
			t.Fatalf("not *ast.ObjectLiteral, got is %T", obj)
		}
		if !cmp.Equal(*got, tt.want) {
			t.Errorf("got ObjectLiteral is %v, want ObjectLiteral is %v", got, &tt.want)
		}
	}
}

func TestParseArray(t *testing.T) {
	tests := []struct {
		input string
		want  ast.ArrayLiteral
	}{
		{`[]`, ast.ArrayLiteral{Values: []ast.Value{}}},
		{`["1", -123.4, true, null, [1, 2]]`, ast.ArrayLiteral{
			Values: []ast.Value{
				&ast.StringLiteral{Value: "1"},
				&ast.NumberLiteral{Value: -123.4},
				&ast.BoolLiteral{Value: true},
				&ast.NullLiteral{},
				&ast.ArrayLiteral{
					Values: []ast.Value{
						&ast.NumberLiteral{Value: 1},
						&ast.NumberLiteral{Value: 2},
					},
				},
			},
		}},
	}

	for _, tt := range tests {
		p := New(tt.input)
		obj, _ := p.Parse()

		got, ok := obj.(*ast.ArrayLiteral)
		if !ok {
			t.Fatalf("not *ast.ArrayLiteral, got is %T", obj)
		}
		if !cmp.Equal(*got, tt.want) {
			t.Errorf("got ArrayLiteral is %v, want ArrayLiteral is %v", got, &tt.want)
		}
	}
}
