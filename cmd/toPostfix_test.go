package main

import (
	"testing"
)

func TestToPostfix(t *testing.T) {
	type testPair struct {
		value  InfixExpr
		result PostfixExpr
	}

	var tests = []testPair{
		{InfixExpr("(2+3)*4"), PostfixExpr("2 3 + 4 * ")},
		{InfixExpr("2+(3*4)"), PostfixExpr("2 3 4 * + ")},
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
