package service

import (
	"calculator/application/entity"
	"testing"
)

func TestToPostfix(t *testing.T) {
	type testPair struct {
		value  entity.InfixExpr
		result entity.PostfixExpr
	}

	var tests = []testPair{
		{entity.InfixExpr("(2+3)*4"), entity.PostfixExpr("2 3 + 4 * ")},
		{entity.InfixExpr("2+(3*4)"), entity.PostfixExpr("2 3 4 * + ")},
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
