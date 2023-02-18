package main

import "strings"

func validSymbols() []string {
	return []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
		"+", "-", "/", "*",
		"(", ")",
		"^",
	}
}

func symbolsIsValid(v *Validator, input string) bool {
	validSymbols := validSymbols()
	for _, s := range strings.Split(input, "") {
		if !contains(validSymbols, s) {
			v.AddError("Wrong symbol: " + s)
		}
	}

	return v.Valid()
}
