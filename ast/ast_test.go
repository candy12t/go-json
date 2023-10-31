package ast

import (
	"testing"
)

func TestString(t *testing.T) {
	want := `["", "abc", 0, -123.4, true, false, null, {}, {"key": "value"}, [], [1, 2]]`
	json := &ArrayLiteral{
		Values: []Value{
			&StringLiteral{Value: ""},
			&StringLiteral{Value: "abc"},
			&NumberLiteral{Value: 0},
			&NumberLiteral{Value: -123.4},
			&BoolLiteral{Value: true},
			&BoolLiteral{Value: false},
			&NullLiteral{},
			&ObjectLiteral{Pairs: nil},
			&ObjectLiteral{Pairs: map[StringLiteral]Value{
				{Value: "key"}: &StringLiteral{Value: "value"},
			}},
			&ArrayLiteral{
				Values: nil,
			},
			&ArrayLiteral{
				Values: []Value{
					&NumberLiteral{Value: 1},
					&NumberLiteral{Value: 2},
				},
			},
		},
	}

	if json.String() != want {
		t.Errorf("json.String() return %q", json.String())
	}
}
