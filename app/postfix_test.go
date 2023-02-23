package app

import (
	"testing"
)

func TestToPostfix(t *testing.T) {
	type testPair struct {
		value  Infix
		result Postfix
	}

	var tests = []testPair{
		{Infix("(2+3)*4"), Postfix("2 3 + 4 * ")},
		{Infix("2+(3*4)"), Postfix("2 3 4 * + ")},
	}

	for _, pair := range tests {
		v := ToPostfix(pair.value)
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
