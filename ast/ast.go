package ast

import (
	"bytes"
	"fmt"
	"strings"
)

type Value interface {
	String() string
}

type ObjectLiteral struct {
	Pairs map[StringLiteral]Value
}

func (o *ObjectLiteral) String() string {
	var out bytes.Buffer

	pairs := make([]string, 0, len(o.Pairs))
	for k, v := range o.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", k.String(), v.String()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

type ArrayLiteral struct {
	Values []Value
}

func (a *ArrayLiteral) String() string {
	var out bytes.Buffer

	values := make([]string, 0, len(a.Values))
	for _, v := range a.Values {
		values = append(values, v.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(values, ", "))
	out.WriteString("]")

	return out.String()
}

type NumberLiteral struct {
	Value float64
}

func (n *NumberLiteral) String() string {
	return fmt.Sprintf("%v", n.Value)
}

type StringLiteral struct {
	Value string
}

func (s *StringLiteral) String() string {
	return fmt.Sprintf(`"%s"`, s.Value)
}

type BoolLiteral struct {
	Value bool
}

func (b *BoolLiteral) String() string {
	return fmt.Sprintf("%v", b.Value)
}

type NullLiteral struct {
}

func (n *NullLiteral) String() string {
	return "null"
}
