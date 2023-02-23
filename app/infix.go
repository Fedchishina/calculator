package app

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Fedchishina/calculator/app/stack"
)

type Infix string

func (i Infix) Validate() error {
	switch {
	case i == "":
		return errors.New("the input data is empty. Enter data please")
	case symbolsIsValid(i) != nil:
		return symbolsIsValid(i)
	case !checkBalance(i):
		return errors.New("the input data is invalid: problem with parentheses")
	}

	return nil
}

func symbolsIsValid(input Infix) error {
	validSymbols := "1234567890+-/*()^"

	for _, i := range input {
		if !strings.ContainsRune(validSymbols, i) {
			return fmt.Errorf("wrong symbol in input data: %q ", i)
		}
	}

	return nil
}

func checkBalance(input Infix) bool {
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
