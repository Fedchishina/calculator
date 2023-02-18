package main

import (
	"github.com/Fedchishina/calculator/stack"
)

type InfixExpr string

func (i *InfixExpr) Validate(v *Validator, value string) {
	v.Check(value != "", "the input data is empty. Enter data please")
	v.Check(symbolsIsValid(v, value), "the input data is invalid")
	v.Check(checkBalance(value), "the input data is invalid: problem with parentheses")
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
