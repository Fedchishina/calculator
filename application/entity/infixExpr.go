package entity

import (
	"calculator/application/stack"
	"calculator/application/validator"
	"strings"
)

type InfixExpr string

func (i *InfixExpr) Validate(v *validator.Validator, value string) {
	v.Check(value != "", "the input data is empty. Enter data please")
	v.Check(symbolsValid(v, value), "the input data is invalid")
	v.Check(checkBalance(value), "the input data is invalid: problem with parentheses")
}

func symbolsValid(v *validator.Validator, input string) bool {
	validSymbols := getValidSymbols()
	for _, s := range strings.Split(input, "") {
		if !contains(validSymbols, s) {
			v.AddError("Wrong symbol: " + s)
		}
	}

	return v.Valid()
}

func contains(symbols []string, s string) bool {
	for _, n := range symbols {
		if s == n {
			return true
		}
	}
	return false
}

func getValidSymbols() []string {
	return []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
		"+", "-", "/", "*",
		"(", ")",
		"^",
	}
}

func checkBalance(input string) bool {
	st := stack.Stack{}

	for _, char := range input {
		switch char {
		case '(':
			st.Push(string(char))
		case ')':
			top := st.Top()
			if top == "(" {
				st.Pop()
			} else {
				return false
			}
		}
	}

	if st.IsEmpty() {
		return true
	}
	return false
}
