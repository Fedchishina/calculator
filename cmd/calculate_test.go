package main

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	type testPair struct {
		value  PostfixExpr
		result string
	}

	var tests = []testPair{
		{PostfixExpr("2 3 + 4 * "), "20.000000"},
		{PostfixExpr("2 3 4 * + "), "14.000000"},
	}

	for _, pair := range tests {
		v := Calculate(pair.value)
		if v != pair.result {
			t.Error(
				"For", pair.value,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}