package service

import (
	"calculator/application/entity"
	"testing"
)

func TestCalc(t *testing.T) {
	type testPair struct {
		value  entity.PostfixExpr
		result string
	}

	var tests = []testPair{
		{entity.PostfixExpr("2 3 + 4 * "), "20.000000"},
		{entity.PostfixExpr("2 3 4 * + "), "14.000000"},
	}

	for _, pair := range tests {
		v := Calc(pair.value)
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
