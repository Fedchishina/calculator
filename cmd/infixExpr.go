package main

import (
	"errors"
	"fmt"
	"github.com/Fedchishina/calculator/stack"
	"strings"
)

type InfixExpr string

func (i InfixExpr) Validate() error {
	value := string(i)

	if value == "" {
		return errors.New("the input data is empty. Enter data please")
	}

	if err := symbolsIsValid(value); err != nil {
		return err
	}

	if !checkBalance(value) {
		return errors.New("the input data is invalid: problem with parentheses")
	}

	return nil
}

func symbolsIsValid(input string) error {
	validSymbols := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
		"+", "-", "/", "*",
		"(", ")",
		"^",
	}

	for _, s := range strings.Split(input, "") {
		if !contains(validSymbols, s) {
			return fmt.Errorf("wrong symbol in input data: %q ", s)
		}
	}

	return nil
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
